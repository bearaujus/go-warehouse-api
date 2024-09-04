package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func (h *productHandlerHTTPImpl) RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc) {
	s.Use(ms...)

	s.GET("/products", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.GetProductsWithStock)
	s.POST("/products", mAuth.AuthenticateUser(), mAuth.AuthorizeSeller(), h.CreateProduct)
}
