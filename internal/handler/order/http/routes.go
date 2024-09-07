package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func (h *orderHandlerHTTPImpl) RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc) {
	s.Use(ms...)

	s.GET("/buyer/orders", mAuth.AuthenticateUser(), mAuth.AuthorizeBuyer(), h.GetOrdersByUserIdAndStatus)
	s.POST("/buyer/orders", mAuth.AuthenticateUser(), mAuth.AuthorizeBuyer(), h.CreateOrder)
	s.POST("/buyer/orders/:id/complete", mAuth.AuthenticateUser(), mAuth.AuthorizeBuyer(), h.CompleteOrder)
}
