package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func (h *warehouseHandlerHTTPImpl) RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc) {
	s.Use(ms...)

	s.GET("/warehouses", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.GetWarehouses)
	s.POST("/warehouses", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.CreateWarehouse)
	s.POST("/warehouses/:id/inbound", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.CreateWarehouseInboundTransaction)
}
