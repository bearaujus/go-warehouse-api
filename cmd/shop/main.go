package main

import (
	. "github.com/bearaujus/go-warehouse-api/internal/handler/shop/http"
	. "github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/tracker"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/postgres_cacher"
	. "github.com/bearaujus/go-warehouse-api/internal/resource/shop/postgres"
	. "github.com/bearaujus/go-warehouse-api/internal/resource/user/http_client"
	. "github.com/bearaujus/go-warehouse-api/internal/usecase/shop"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
)

func main() {
	ctx, cfg, cancel := pkg.InitBaseApp()
	defer cancel()

	log.Printf("Starting %v service...", cfg.ServiceShopContainerName)

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

	rShopPostgres := NewShopResourcePostgres(postgres)
	rUserHTTPClient := NewUserResourceHTTPClient(
		pkg.GenerateHostPort(false, cfg.ServiceUserContainerName, cfg.ServiceUserPort),
		cfg.ServiceShopHTTPCallTimeout,
		cfg.ServiceShopContainerName,
		cfg.ServiceShopHTTPCallAuthKey,
		cfg.ServiceShopHTTPCallAuthTTL,
	)

	uShop := NewShopUsecase(rShopPostgres)

	hShopHTTP := NewShopHandlerHTTP(uShop)

	mAuth := NewAuthMiddleware(rUserHTTPClient, cfg.ServiceUserAuthSecretKey, map[string]string{
		cfg.ServiceUserContainerName: cfg.ServiceUserHTTPCallAuthKey,
	})

	err = httputil.StartHTTPServer(ctx, cfg.ServiceShopPort, func(s *server.Hertz) {
		hShopHTTP.RegisterRoutes(s, mAuth, tracker.MiddlewareTracker())
	})
	if err != nil {
		log.Fatalln(err)
	}
}
