package http

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type UserHandlerHTTP interface {
	GetUserById(ctx context.Context, rCtx *app.RequestContext)
	Register(ctx context.Context, rCtx *app.RequestContext)
	Login(ctx context.Context, rCtx *app.RequestContext)

	RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc)
}

type userHandlerHTTPImpl struct {
	uUser user.UserUsecase
}

func NewUserHandlerHTTP(uUser user.UserUsecase) UserHandlerHTTP {
	return &userHandlerHTTPImpl{uUser: uUser}
}
