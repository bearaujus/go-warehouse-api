// TODO:
// - make more efficient & modular
// - separate upstream functions to usecase
// - make a http call upstream service instead doing this implementation
// - think about failure and rollback interface

package postgres

import (
	"context"
	"fmt"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"time"
)

func (r *orderResourcePostgresImpl) GetOrdersByUserIdAndStatus(ctx context.Context, userId uint64, status model.OrderStatus) ([]*model.Order, error) {
	var orders []*model.Order
	q := r.db.WithContext(ctx).Where("user_id = ?", userId)

	if status != "" {
		q = q.Where("status = ?", status)
	}

	err := q.Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, model.ErrROrderPostgresGetOrdersByUserIdAndStatus.New(err)
	}

	return orders, nil
}

func (r *orderResourcePostgresImpl) CreateOrder(ctx context.Context, userId uint64, orderItems []*model.OrderItem) (*model.Order, error) {
	// CASE: CREATE ORDER
	// 1. create new order with status 'pending' and expiration relative to configurable ttl
	// 2. for loop from orderItems[i] and select warehouse product stocks by orderItems[i].ProductId
	//    where the sum of it is > orderItems[i].Quantity
	// 3. for loop from warehouse product stocks and reduce the stocks from it
	//   relative to remaining orderItems[i].Quantity. Example (A: 5, B: 3 | BuyReq: 7 | A: 0, B: 1)
	// 4. update the warehouse product stocks
	// 5. create order item reservations
	// 6. update order for assign the order total price
	// 7. create order items

	var err error
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. create new order with status 'pending' and expiration relative to configurable ttl
	createdAt := time.Now()
	expiresAt := createdAt.Add(r.orderExpirationTTL)
	order := model.Order{
		UserId:    userId,
		Status:    model.OrderStatusPending,
		CreatedAt: &createdAt,
		ExpiresAt: &expiresAt,
	}
	err = tx.Create(&order).Error
	if err != nil {
		return nil, model.ErrROrderPostgresCreateOrder.New(err)
	}

	// 2. for loop from orderItems[i] and select warehouse product stocks by orderItems[i].ProductId
	//    where the sum of it is > orderItems[i].Quantity
	for i := range orderItems {
		orderItems[i].OrderId = order.Id
		var wpss []*model.WarehouseProductStock
		err = tx.Joins("JOIN warehouses ON warehouses.id = warehouse_product_stocks.warehouse_id").
			Where("warehouse_product_stocks.product_id = ?", orderItems[i].ProductId).
			Where("warehouses.status = ?", model.WarehouseStatusActive).
			Where("(SELECT SUM(wpss.quantity) FROM warehouse_product_stocks wpss JOIN warehouses w ON wpss.warehouse_id = w.id WHERE wpss.product_id = ? AND w.status = ?) >= ?", orderItems[i].ProductId, model.WarehouseStatusActive, orderItems[i].Quantity).
			Preload("Product").
			Find(&wpss).Error
		if err != nil {
			return nil, model.ErrROrderPostgresCreateOrder.New(fmt.Printf("warehouse product stocks is not enough: %v", err))
		}

		// 3. for loop from warehouse product stocks and reduce the stocks from it
		//   relative to remaining orderItems[i].Quantity. Example (A: 5, B: 3 | BuyReq: 7 | A: 0, B: 1)
		buyQuantity := orderItems[i].Quantity
		var (
			oirt []*model.OrderItemReservation  // order item reservation update targets
			wpst []*model.WarehouseProductStock // warehouse product stock update targets
		)
		for j := range wpss {
			if orderItems[i].Price == 0 {
				orderItems[i].Price = wpss[j].Product.Price // set order item price
			}
			oir := &model.OrderItemReservation{OrderId: order.Id, WarehouseProductStockId: wpss[j].Id, Quantity: wpss[j].Quantity}
			if buyQuantity <= wpss[j].Quantity { // indicates buyers stock demand already fulfilled
				oir.Quantity = buyQuantity
				wpss[j].Quantity -= buyQuantity
				wpst = append(wpst, wpss[j])
				oirt = append(oirt, oir)
				break
			}
			buyQuantity -= wpss[j].Quantity
			wpss[j].Quantity = 0
			wpst = append(wpst, wpss[j])
			oirt = append(oirt, oir)
		}

		// add order item price to order total price
		order.TotalPrice += float64(orderItems[i].Quantity) * orderItems[i].Price

		// 4. update the warehouse product stocks
		err = tx.Save(&wpst).Error
		if err != nil {
			return nil, model.ErrROrderPostgresCreateOrder.New(err)
		}

		// 5. create order item reservations
		err = tx.Create(&oirt).Error
		if err != nil {
			return nil, model.ErrROrderPostgresCreateOrder.New(err)
		}
	}

	// 6. update order for assign the order total price
	err = tx.Save(&order).Error
	if err != nil {
		return nil, model.ErrROrderPostgresCreateOrder.New(err)
	}

	// 7. create order items
	err = tx.Create(&orderItems).Error
	if err != nil {
		return nil, model.ErrROrderPostgresCreateOrder.New(err)
	}

	// commit the db transaction
	err = tx.Commit().Error
	if err != nil {
		return nil, model.ErrROrderPostgresCreateOrder.New(err)
	}

	order.OrderItems = orderItems
	return &order, nil
}

func (r *orderResourcePostgresImpl) CompleteOrder(ctx context.Context, userId uint64, id uint64) (*model.Order, error) {
	// CASE: COMPLETE ORDER
	// 1. get order by id and user id
	// 2. delete order item reservations
	// 3. update order status to 'completed' so the cron will not process this order id at ProcessExpiredOrders
	// 4. update order expires_at to NULL
	var err error
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. get order by id and user id
	order := model.Order{}
	err = tx.
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Where("completed_at IS NULL").
		Where("status = ?", model.OrderStatusPending).
		Preload("OrderItems").
		First(&order).Error
	if err != nil {
		return nil, model.ErrROrderPostgresCompleteOrder.New("order is not exist or already completed")
	}

	// 2. delete order item reservations
	var oirt []*model.OrderItemReservation
	err = tx.Where("order_id = ?", order.Id).Find(&oirt).Error
	if err != nil {
		return nil, model.ErrROrderPostgresCompleteOrder.New(err)
	}

	err = tx.Delete(oirt).Error
	if err != nil {
		return nil, model.ErrROrderPostgresCompleteOrder.New(err)
	}

	// 3. update order status to 'completed' so the cron will not process this order id at ProcessExpiredOrders
	order.Status = model.OrderStatusComplete
	orderCompletedAt := time.Now()
	order.CompletedAt = &orderCompletedAt
	err = tx.Save(&order).Error
	if err != nil {
		return nil, model.ErrROrderPostgresCompleteOrder.New(err)
	}

	// 4. update order expires_at to NULL
	err = tx.Exec("UPDATE orders SET expires_at = NULL WHERE id = $1", order.Id).Error
	if err != nil {
		return nil, model.ErrROrderPostgresCompleteOrder.New(err)
	}
	order.ExpiresAt = nil

	// commit the db transaction
	err = tx.Commit().Error
	if err != nil {
		return nil, model.ErrROrderPostgresCompleteOrder.New(err)
	}

	return &order, nil
}

func (r *orderResourcePostgresImpl) ProcessExpiredOrders(ctx context.Context) ([]uint64, error) {
	// CASE: PROCESS EXPIRED ORDERS
	// 1. get expired orders
	// 2. get order item reservations by order id
	// 3. return the reserved stock by updating the warehouse product stocks
	// 4. delete order item reservations
	// 5. update order status to 'expired'

	var err error
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. get expired orders
	var orders []*model.Order
	err = tx.
		Where("status = ?", model.OrderStatusPending).
		Where("completed_at IS NULL").
		Where("expires_at <= ?", time.Now()).
		Find(&orders).Error
	if err != nil {
		return nil, model.ErrROrderPostgresProcessExpiredOrders.New(err)
	}

	affectedOrderIds := make([]uint64, len(orders))
	for i, order := range orders {
		// 2. get order item reservations by order id
		var oirs []*model.OrderItemReservation
		err = tx.Where("order_id = ?", order.Id).Find(&oirs).Error
		if err != nil {
			return nil, model.ErrROrderPostgresProcessExpiredOrders.New(err)
		}

		// 3. return the reserved stock by updating the warehouse product stocks
		var wpss []*model.WarehouseProductStock
		for _, oir := range oirs {
			wps := model.WarehouseProductStock{Id: oir.WarehouseProductStockId}
			err = tx.First(&wps).Error
			if err != nil {
				return nil, model.ErrROrderPostgresProcessExpiredOrders.New(err)
			}
			wps.Quantity += oir.Quantity
			wpss = append(wpss, &wps)
		}

		err = tx.Save(wpss).Error
		if err != nil {
			return nil, model.ErrROrderPostgresProcessExpiredOrders.New(err)
		}

		// 4. delete order item reservations
		err = tx.Delete(oirs).Error
		if err != nil {
			return nil, model.ErrROrderPostgresProcessExpiredOrders.New(err)
		}

		// 5. update order status to 'expired'
		order.Status = model.OrderStatusExpired
		err = tx.Save(order).Error
		if err != nil {
			return nil, model.ErrROrderPostgresProcessExpiredOrders.New(err)
		}

		affectedOrderIds[i] = order.Id
	}

	// commit the db transaction
	err = tx.Commit().Error
	if err != nil {
		return nil, model.ErrROrderPostgresProcessExpiredOrders.New(err)
	}

	return affectedOrderIds, nil
}
