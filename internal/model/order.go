package model

import (
	"gorm.io/plugin/optimisticlock"
	"time"
)

type Order struct {
	Id           uint64                 `gorm:"primaryKey" json:"id,omitempty"`
	UserId       uint64                 `gorm:"index" json:"user_id,omitempty"`
	TotalPrice   float64                `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Status       string                 `gorm:"size:50;not null" json:"status,omitempty"`
	CreatedAt    *time.Time             `gorm:"autoCreateTime" json:"created_at,omitempty"`
	Items        []*OrderItem           `gorm:"foreignKey:OrderId" json:"items,omitempty"`
	Reservations []*StockReservation    `gorm:"foreignKey:OrderId" json:"reservations,omitempty"`
	Version      optimisticlock.Version `json:"-"`
}

type OrderItem struct {
	Id        uint64                 `gorm:"primaryKey" json:"id,omitempty"`
	OrderId   uint64                 `gorm:"index" json:"order_id,omitempty"`
	ProductId uint64                 `gorm:"index" json:"product_id,omitempty"`
	Quantity  int                    `gorm:"not null" json:"quantity,omitempty"`
	Price     float64                `gorm:"type:decimal(10,2);not null" json:"price,omitempty"`
	Product   *Product               `gorm:"foreignKey:ProductId;constraint:OnDelete:SET NULL;" json:"product,omitempty"`
	Version   optimisticlock.Version `json:"-"`
}

type StockReservation struct {
	Id               uint64                 `gorm:"primaryKey" json:"id,omitempty"`
	ProductId        uint64                 `gorm:"index" json:"product_id,omitempty"`
	WarehouseId      uint64                 `gorm:"index" json:"warehouse_id,omitempty"`
	OrderId          uint64                 `gorm:"index" json:"order_id,omitempty"`
	ReservedQuantity int                    `gorm:"not null" json:"reserved_quantity,omitempty"`
	ReservedAt       time.Time              `gorm:"autoCreateTime" json:"reserved_at,omitempty"`
	ExpiresAt        time.Time              `gorm:"not null" json:"expires_at,omitempty"`
	Product          *Product               `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE;" json:"product,omitempty"`
	Warehouse        *Warehouse             `gorm:"foreignKey:WarehouseId;constraint:OnDelete:CASCADE;" json:"warehouse,omitempty"`
	Order            *Order                 `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE;" json:"order,omitempty"`
	Version          optimisticlock.Version `json:"-"`
}
