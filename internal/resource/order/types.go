package order

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

type OrderResource interface {
	GetOrdersByUserIdAndStatus(ctx context.Context, userId uint64, status model.OrderStatus) ([]*model.Order, error)
	CreateOrder(ctx context.Context, userId uint64, orderItems []*model.OrderItem) (*model.Order, error)
	CompleteOrder(ctx context.Context, userId uint64, id uint64) (*model.Order, error)
	ProcessExpiredOrders(ctx context.Context) ([]uint64, error)
}
