package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func (h *shopHandlerHTTPImpl) RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc) {
	s.Use(ms...)

	s.POST("/internal/shops", mAuth.AuthenticateService(), h.CreateShop)
}
