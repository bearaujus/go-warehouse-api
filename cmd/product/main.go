package main

import (
	. "github.com/bearaujus/go-warehouse-api/internal/handler/product/http"
	. "github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/tracker"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/postgres_cacher"
	. "github.com/bearaujus/go-warehouse-api/internal/resource/product/postgres"
	. "github.com/bearaujus/go-warehouse-api/internal/resource/user/http_client"
	. "github.com/bearaujus/go-warehouse-api/internal/usecase/product"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
)

func main() {
	cfg, cancel := pkg.InitBaseApp()
	defer cancel()

	log.Printf("Starting %v service...", cfg.ServiceProductContainerName)

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

	rProductPostgres := NewProductResourcePostgres(postgres)
	rUserHTTPClient := NewUserResourceHTTPClient(
		pkg.GenerateHostPort(false, cfg.ServiceUserContainerName, cfg.ServiceUserPort),
		cfg.ServiceProductHTTPCallTimeout,
		cfg.ServiceProductContainerName,
		cfg.ServiceProductHTTPCallAuthKey,
		cfg.ServiceProductHTTPCallAuthTTL,
	)

	uProduct := NewProductUsecase(rProductPostgres)

	hProductHTTP := NewProductHandlerHTTP(uProduct)

	mAuth := NewAuthMiddleware(rUserHTTPClient, cfg.ServiceUserAuthSecretKey, nil)

	err = httputil.StartHTTPServer(cfg.ServiceProductPort, func(s *server.Hertz) {
		hProductHTTP.RegisterRoutes(s, mAuth, tracker.MiddlewareTracker())
	})
	if err != nil {
		log.Fatalln(err)
	}
}
