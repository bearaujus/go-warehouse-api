package postgres

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *orderResourcePostgresImpl) GetOrdersByUser(ctx context.Context, userId uint64) ([]*model.Order, error) {
	var orders []*model.Order
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		return nil, model.ErrROrderPostgresGetOrdersByUser.New(err)
	}
	return orders, nil
}

func (r *orderResourcePostgresImpl) CreateOrder(ctx context.Context, order *model.Order) (uint64, error) {
	order.Id = 0
	err := r.db.WithContext(ctx).Create(order).Error
	if err != nil {
		return 0, model.ErrROrderPostgresCreateOrder.New(err)
	}
	return order.Id, nil
}
