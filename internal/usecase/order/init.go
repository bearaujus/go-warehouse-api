package order

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/resource/order"
)

type OrderUsecase interface {
	GetOrdersByUser(ctx context.Context, userId uint64) ([]*model.Order, error)
	CreateOrder(ctx context.Context, order *model.Order) (uint64, error)
}

type orderUsecaseImpl struct {
	rOrder order.OrderResource
}

func NewOrderUsecase(rOrderPostgres order.OrderResource) OrderUsecase {
	return &orderUsecaseImpl{rOrder: rOrderPostgres}
}
