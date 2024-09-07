package http_client

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *orderResourceHTTPClientImpl) GetOrdersByUserIdAndStatus(ctx context.Context, userId uint64) ([]*model.Order, error) {
	return nil, model.ErrROrderHTTPClientGetOrdersByUserIdAndStatus.New(model.ErrCommonNotImplemented)
}

func (r *orderResourceHTTPClientImpl) CreateOrder(ctx context.Context, userId uint64, orderItems []*model.OrderItem) (*model.Order, error) {
	return nil, model.ErrROrderHTTPClientCreateOrder.New(model.ErrCommonNotImplemented)
}

func (r *orderResourceHTTPClientImpl) CompleteOrder(ctx context.Context, userId uint64, id uint64) (*model.Order, error) {
	return nil, model.ErrROrderHTTPClientCompleteOrder.New(model.ErrCommonNotImplemented)
}

func (r *orderResourceHTTPClientImpl) ProcessExpiredOrders(ctx context.Context) ([]uint64, error) {
	return nil, model.ErrROrderHTTPClientProcessExpiredOrders.New(model.ErrCommonNotImplemented)
}
