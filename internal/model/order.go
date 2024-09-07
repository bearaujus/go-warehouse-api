package model

import (
	"fmt"
	"gorm.io/plugin/optimisticlock"
	"time"
)

type OrderStatus string

const (
	OrderStatusPending  OrderStatus = "pending"
	OrderStatusComplete OrderStatus = "completed"
	OrderStatusExpired  OrderStatus = "expired"
)

type Order struct {
	Id          uint64                 `json:"id,omitempty"`
	UserId      uint64                 `json:"user_id,omitempty"`
	TotalPrice  float64                `json:"total_price"`
	Status      OrderStatus            `json:"status,omitempty"`
	CreatedAt   *time.Time             `json:"created_at,omitempty"`
	ExpiresAt   *time.Time             `json:"expires_at,omitempty"`
	CompletedAt *time.Time             `json:"completed_at,omitempty"`
	Version     optimisticlock.Version `json:"-"`

	// fk
	OrderItems            []*OrderItem            `json:"order_item,omitempty"`
	OrderItemReservations []*OrderItemReservation `json:"order_item_reservation,omitempty"`
}

type OrderItem struct {
	Id        uint64  `json:"id,omitempty"`
	OrderId   uint64  `json:"order_id,omitempty"`
	ProductId uint64  `json:"product_id,omitempty"`
	Quantity  int     `json:"quantity,omitempty"`
	Price     float64 `json:"price,omitempty"`
}

func (m *OrderItem) Validate() error {
	// validate product id
	if m.ProductId == 0 {
		return fmt.Errorf("order item product id is required")
	}

	// validate quantity
	if m.Quantity == 0 {
		return fmt.Errorf("order item quantity is required")
	}

	if m.Quantity < 0 {
		return fmt.Errorf("order item quantity is invalid")
	}

	return nil
}

type OrderItemReservation struct {
	Id                      uint64 `json:"id,omitempty"`
	OrderId                 uint64 `json:"order_id,omitempty"`
	WarehouseProductStockId uint64 `json:"warehouse_product_stock_id,omitempty"`
	Quantity                int    `json:"quantity,omitempty"`
}
