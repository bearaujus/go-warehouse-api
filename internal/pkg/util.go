package pkg

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"github.com/bearaujus/bworker/flex"
	"github.com/bearaujus/bworker/pool"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-gorm/caches/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func InitBaseApp() (context.Context, *config.Config, func()) {
	ctx, cancel := context.WithCancel(context.Background())

	logPath := flag.String("log", "", "path to the log file")
	flag.Parse()

	closer, err := initLogStd(*logPath)
	if err != nil {
		log.Printf("failed to initialize log: %v", err)
		os.Exit(1)
	}

	return ctx, config.Read(), func() {
		closer()
		cancel()
	}
}

func initLogStd(name string) (func(), error) {
	logFile, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_SYNC, os.ModePerm)
	if err != nil {
		return nil, err
	}

	log.Printf("New log file will set. Future output log will appear in: %v", name)

	// set stdout log file
	os.Stdout = logFile
	os.Stderr = logFile

	// set default log file
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(logFile)

	// set hertz router log file
	hlog.SetOutput(logFile)
	hlog.SetLevel(hlog.LevelTrace)

	return func() {
		_ = logFile.Close()
	}, nil
}

func InitRedis(numRetry int, initDelay time.Duration, address string, port string, password string) (*redis.Client, error) {
	var rdb *redis.Client
	var wErr error

	w := flex.NewBWorkerFlex(
		flex.WithRetry(numRetry),
		flex.WithError(&wErr),
	)

	w.Do(func() error {
		time.Sleep(initDelay)

		rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", address, port),
			Password: password,
		})

		err := rdb.Ping(context.Background()).Err()
		if err != nil {
			log.Printf("fail to initialize redis: %v. retrying...", err)
			return err
		}

		return nil
	})

	w.Wait()
	if wErr != nil {
		return nil, wErr
	}

	return rdb, nil
}

func InitPostgres(numRetry int, initDelay time.Duration, dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var wErr error

	w := flex.NewBWorkerFlex(
		flex.WithRetry(numRetry),
		flex.WithError(&wErr),
	)

	w.Do(func() error {
		time.Sleep(initDelay)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.New(log.Default(), logger.Config{
				LogLevel: logger.Silent,
			}),
		})

		if err != nil {
			log.Printf("fail to initialize postgres: %v. retrying...", err)
			return err
		}

		return nil
	})

	w.Wait()
	if wErr != nil {
		return nil, wErr
	}

	return db, nil
}

func UsePostgresReplicas(numRetry int, initDelay time.Duration, masterDB *gorm.DB, slavesDSN []string) error {
	slavesDialector := make([]gorm.Dialector, len(slavesDSN))
	mu := sync.Mutex{}
	var wErr error

	w := pool.NewBWorkerPool(
		len(slavesDSN),
		pool.WithError(&wErr),
	)
	defer w.Shutdown()

	for i, v := range slavesDSN {
		icp := i
		w.Do(func() error {
			slaveDB, err := InitPostgres(numRetry, initDelay, v)
			if err != nil {
				return err
			}

			mu.Lock()
			slavesDialector[icp] = slaveDB.Dialector
			mu.Unlock()
			return nil
		})
	}

	w.Wait()
	if wErr != nil {
		return wErr
	}

	w2 := flex.NewBWorkerFlex(
		flex.WithRetry(numRetry),
		flex.WithError(&wErr),
	)

	w2.Do(func() error {
		time.Sleep(initDelay)

		err := masterDB.Use(dbresolver.Register(
			dbresolver.Config{
				Sources:           []gorm.Dialector{masterDB.Dialector},
				Replicas:          slavesDialector,
				Policy:            dbresolver.RoundRobinPolicy(),
				TraceResolverMode: false,
			},
		))

		if err != nil {
			log.Printf("fail to initialize postgres replicas: %v. retrying...", err)
			return err
		}

		return nil
	})

	w2.Wait()
	if wErr != nil {
		return wErr
	}

	log.Printf("%v Postgres replicas attached", len(slavesDialector))
	return nil
}

func UsePostgresCache(masterDB *gorm.DB, cacher caches.Cacher) error {
	err := masterDB.Use(&caches.Caches{Conf: &caches.Config{Easer: true, Cacher: cacher}})
	if err != nil {
		return err
	}
	log.Print("Postgres cache attached")
	return nil
}

func WaitShutdownSigterm(ctx context.Context, shutdownFunc func(), attachedStatusChan chan struct{}) {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	log.Print("Shutdown sigterm hook attached")
	attachedStatusChan <- struct{}{}
	select {
	case <-ctx.Done():
		log.Print("Shutdown sigterm hook detached due to parent context cancellation. Shutting down server...")
	case <-stopChan:
		log.Print("Shutdown signal received. Shutting down server...")
	}

	signal.Stop(stopChan)
	close(stopChan)
	shutdownFunc()
}

func ToSHA256Hash(s string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func StringToUint64(s string) (uint64, error) {
	o, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return o, nil
}

func GenerateHostPort(isHTTPS bool, host string, port string) string {
	protocol := "http"
	if isHTTPS {
		protocol = "https"
	}
	return fmt.Sprintf("%v://%v:%v", protocol, host, port)
}
