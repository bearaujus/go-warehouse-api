package order

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

type OrderResource interface {
	GetOrdersByUser(ctx context.Context, userId uint64) ([]*model.Order, error)
	CreateOrder(ctx context.Context, order *model.Order) (uint64, error)
}
