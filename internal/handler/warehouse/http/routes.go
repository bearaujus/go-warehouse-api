package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func (h *warehouseHandlerHTTPImpl) RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc) {
	s.Use(ms...)

	s.GET("/seller/warehouses", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.GetWarehousesByShopUserId)
	s.POST("/seller/warehouses", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.CreateWarehouse)
	s.PUT("/seller/warehouses/:id", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.UpdateWarehouse)
	s.POST("/seller/warehouses/:id/product-stocks/:product_id/add-stock", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.AddWarehouseProductStock)
	s.POST("/seller/warehouses/:id/product-stocks/:product_id/transfer", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.TransferWarehouseProductStock)
}
