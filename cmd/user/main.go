package main

import (
	. "github.com/bearaujus/go-warehouse-api/internal/handler/user/http"
	. "github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/tracker"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/postgres_cacher"
	. "github.com/bearaujus/go-warehouse-api/internal/resource/shop/http_client"
	. "github.com/bearaujus/go-warehouse-api/internal/resource/user/postgres"
	. "github.com/bearaujus/go-warehouse-api/internal/usecase/user"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
)

func main() {
	ctx, cfg, cancel := pkg.InitBaseApp()
	defer cancel()

	log.Printf("Starting %v service...", cfg.ServiceUserContainerName)

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

	rUserPostgres := NewUserResourcePostgres(postgres)
	rShopHTTPClient := NewShopResourceHTTPClient(
		pkg.GenerateHostPort(false, cfg.ServiceShopContainerName, cfg.ServiceShopPort),
		cfg.ServiceUserHTTPCallTimeout,
		cfg.ServiceUserContainerName,
		cfg.ServiceUserHTTPCallAuthKey,
		cfg.ServiceUserHTTPCallAuthTTL,
	)

	uUser := NewUserUsecase(rUserPostgres, rShopHTTPClient, cfg.ServiceUserAuthSecretKey, cfg.ServiceUserAuthTTL)

	hUserHTTP := NewUserHandlerHTTP(uUser)

	mAuth := NewAuthMiddleware(rUserPostgres, cfg.ServiceUserAuthSecretKey, map[string]string{
		cfg.ServiceOrderContainerName:     cfg.ServiceOrderHTTPCallAuthKey,
		cfg.ServiceProductContainerName:   cfg.ServiceProductHTTPCallAuthKey,
		cfg.ServiceShopContainerName:      cfg.ServiceShopHTTPCallAuthKey,
		cfg.ServiceWarehouseContainerName: cfg.ServiceWarehouseHTTPCallAuthKey,
	})

	err = httputil.StartHTTPServer(ctx, cfg.ServiceUserPort, func(s *server.Hertz) {
		hUserHTTP.RegisterRoutes(s, mAuth, tracker.MiddlewareTracker())
	})
	if err != nil {
		log.Fatalln(err)
	}
}
