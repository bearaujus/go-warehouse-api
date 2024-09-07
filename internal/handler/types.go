package handler

import (
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type HTTPHandler interface {
	RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc)
}
