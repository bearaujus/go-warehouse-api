package http

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/warehouse"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type WarehouseHandlerHTTP interface {
	GetWarehouses(ctx context.Context, rCtx *app.RequestContext)
	CreateWarehouse(ctx context.Context, rCtx *app.RequestContext)
	CreateWarehouseInboundTransaction(ctx context.Context, rCtx *app.RequestContext)

	RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc)
}

type warehouseHandlerHTTPImpl struct {
	uWarehouse warehouse.WarehouseUsecase
}

func NewWarehouseHandlerHTTP(uWarehouse warehouse.WarehouseUsecase) WarehouseHandlerHTTP {
	return &warehouseHandlerHTTPImpl{uWarehouse: uWarehouse}
}
