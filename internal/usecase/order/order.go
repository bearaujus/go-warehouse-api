package order

import (
	"context"
	"errors"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/errwrap"
	"log"
	"time"
)

func (u *orderUsecaseImpl) GetOrdersByUserIdAndStatus(ctx context.Context, userId uint64, status model.OrderStatus) ([]*model.Order, error) {
	if status != "" && status != model.OrderStatusPending && status != model.OrderStatusComplete && status != model.OrderStatusExpired {
		return nil, model.ErrUOrderGetOrdersByUserIdAndStatus.New("invalid status")
	}
	return u.rOrder.GetOrdersByUserIdAndStatus(ctx, userId, status)
}

func (u *orderUsecaseImpl) CreateOrder(ctx context.Context, userId uint64, orderItems []*model.OrderItem) (*model.Order, error) {
	if len(orderItems) == 0 {
		return nil, model.ErrUOrderCreateOrder.New("empty order items")
	}

	// merge duplicate order items
	tmp := make(map[uint64]int) // map[product_id]quantity
	for _, orderItem := range orderItems {
		_, ok := tmp[orderItem.ProductId]
		if !ok {
			tmp[orderItem.ProductId] = 0
		}
		tmp[orderItem.ProductId] += orderItem.Quantity
	}

	// create new order items
	newOrderItems := make([]*model.OrderItem, len(tmp))
	i := 0
	for productId, quantity := range tmp {
		newOrderItems[i] = &model.OrderItem{
			ProductId: productId,
			Quantity:  quantity,
		}
		i++
	}

	// validate
	for _, orderItem := range newOrderItems {
		err := orderItem.Validate()
		if err != nil {
			return nil, model.ErrUOrderCreateOrder.New(err)
		}
	}

	return u.rOrder.CreateOrder(ctx, userId, newOrderItems)
}

func (u *orderUsecaseImpl) CompleteOrder(ctx context.Context, userId uint64, id uint64) (*model.Order, error) {
	return u.rOrder.CompleteOrder(ctx, userId, id)
}

func (u *orderUsecaseImpl) CronJobExpiredOrder(ctx context.Context) func() {
	return func() {
		log.Println("[CRON] Starting cron job to process expired orders...")
		startTime := time.Now()
		affectedOrderIds, err := u.rOrder.ProcessExpiredOrders(ctx)
		processTime := time.Since(startTime).Seconds()
		if err != nil {
			var errWrap errwrap.ErrWrap
			if errors.As(err, &errWrap) {
				log.Printf("[CRON] Cron job for processing expired orders failed: %v at %v", err.Error(), errWrap.StackTrace())
			} else {
				log.Printf("[CRON] Cron job for processing expired orders failed: %v", err)
			}
			log.Printf("[CRON] Processing Time: %.9f seconds", processTime)
			return
		}
		log.Printf("[CRON] Cron job for processing expired orders completed successfully.")
		log.Printf("[CRON] Affected Orders: %v", affectedOrderIds)
		log.Printf("[CRON] Processing Time: %.9f seconds", processTime)
	}
}
