package http

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/shop"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type ShopHandlerHTTP interface {
	GetShops(ctx context.Context, rCtx *app.RequestContext)
	CreateShop(ctx context.Context, rCtx *app.RequestContext)

	RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc)
}

type shopHandlerHTTPImpl struct {
	uShop shop.ShopUsecase
}

func NewShopHandlerHTTP(uShop shop.ShopUsecase) ShopHandlerHTTP {
	return &shopHandlerHTTPImpl{uShop: uShop}
}
