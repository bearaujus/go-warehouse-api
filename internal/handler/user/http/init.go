package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/handler"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/user"
)

type userHandlerHTTPImpl struct {
	uUser user.UserUsecase
}

func NewUserHandlerHTTP(uUser user.UserUsecase) handler.HTTPHandler {
	return &userHandlerHTTPImpl{uUser: uUser}
}
