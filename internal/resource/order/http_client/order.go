package http_client

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *orderResourceHTTPClientImpl) GetOrdersByUser(ctx context.Context, userId uint64) ([]*model.Order, error) {
	return nil, model.ErrROrderHTTPClientGetOrdersByUser.New(model.ErrCommonNotImplemented)
}

func (r *orderResourceHTTPClientImpl) CreateOrder(ctx context.Context, order *model.Order) (uint64, error) {
	return 0, model.ErrROrderHTTPClientCreateOrder.New(model.ErrCommonNotImplemented)
}
