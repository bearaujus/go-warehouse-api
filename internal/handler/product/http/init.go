package http

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type ProductHandlerHTTP interface {
	CreateProduct(ctx context.Context, rCtx *app.RequestContext)
	GetProductsWithStock(ctx context.Context, rCtx *app.RequestContext)

	RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc)
}

type productHandlerHTTPImpl struct {
	uProduct product.ProductUsecase
}

func NewProductHandlerHTTP(uProduct product.ProductUsecase) ProductHandlerHTTP {
	return &productHandlerHTTPImpl{uProduct: uProduct}
}
