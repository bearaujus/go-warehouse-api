package main

import (
	. "github.com/bearaujus/go-warehouse-api/internal/handler/order/http"
	. "github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/tracker"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/postgres_cacher"
	. "github.com/bearaujus/go-warehouse-api/internal/resource/order/postgres"
	. "github.com/bearaujus/go-warehouse-api/internal/resource/user/http_client"
	. "github.com/bearaujus/go-warehouse-api/internal/usecase/order"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
)

func main() {
	cfg, cancel := pkg.InitBaseApp()
	defer cancel()

	log.Printf("Starting %v service...", cfg.ServiceOrderContainerName)

	redis, err := pkg.InitRedis(cfg.RedisInitRetry, cfg.RedisInitDelay, cfg.RedisContainerName, cfg.RedisPort, cfg.RedisPass)
	if err != nil {
		log.Fatalln(err)
	}

	postgres, err := pkg.InitPostgres(cfg.PostgresInitRetry, cfg.PostgresInitDelay, cfg.GetPostgresMasterDSN())
	if err != nil {
		log.Fatalln(err)
	}

	err = pkg.UsePostgresReplicas(cfg.PostgresInitRetry, cfg.PostgresInitDelay, postgres, cfg.GetPostgresReplicasDSN())
	if err != nil {
		log.Fatalln(err)
	}

	err = pkg.UsePostgresCache(postgres, postgres_cacher.NewPostgresCacher(redis, cfg.RedisTTL))
	if err != nil {
		log.Fatalln(err)
	}

	rOrderPostgres := NewOrderResourcePostgres(postgres)
	rUserHTTPClient := NewUserResourceHTTPClient(
		pkg.GenerateHostPort(false, cfg.ServiceUserContainerName, cfg.ServiceUserPort),
		cfg.ServiceOrderHTTPCallTimeout,
		cfg.ServiceOrderContainerName,
		cfg.ServiceOrderHTTPCallAuthKey,
		cfg.ServiceOrderHTTPCallAuthTTL,
	)

	uOrder := NewOrderUsecase(rOrderPostgres)

	hOrderHTTP := NewOrderHandlerHTTP(uOrder)

	mAuth := NewAuthMiddleware(rUserHTTPClient, cfg.ServiceUserAuthSecretKey, nil)

	err = httputil.StartHTTPServer(cfg.ServiceOrderPort, func(s *server.Hertz) {
		hOrderHTTP.RegisterRoutes(s, mAuth, tracker.MiddlewareTracker())
	})
	if err != nil {
		log.Fatalln(err)
	}
}
