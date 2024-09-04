package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type OrderHandlerHTTP interface {
	RegisterRoutes(s *server.Hertz, mAuth *auth.MiddlewareAuth, ms ...app.HandlerFunc)
}

type orderHandlerHTTPImpl struct {
	uOrder order.OrderUsecase
}

func NewOrderHandlerHTTP(uOrder order.OrderUsecase) OrderHandlerHTTP {
	return &orderHandlerHTTPImpl{uOrder: uOrder}
}
