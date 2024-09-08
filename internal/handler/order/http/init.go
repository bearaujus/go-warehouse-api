package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/handler"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/order"
)

type orderHandlerHTTPImpl struct {
	uOrder order.OrderUsecase
}

func NewOrderHandlerHTTP(uOrder order.OrderUsecase) handler.HTTPHandler {
	return &orderHandlerHTTPImpl{uOrder: uOrder}
}
