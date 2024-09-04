package order

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (u *orderUsecaseImpl) GetOrdersByUser(ctx context.Context, userId uint64) ([]*model.Order, error) {
	return u.rOrder.GetOrdersByUser(ctx, userId)
}

func (u *orderUsecaseImpl) CreateOrder(ctx context.Context, order *model.Order) (uint64, error) {
	return u.rOrder.CreateOrder(ctx, order)
}
