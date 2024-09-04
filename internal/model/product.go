package model

import (
	"gorm.io/plugin/optimisticlock"
	"time"
)

type Product struct {
	Id          uint64                 `gorm:"primaryKey" json:"id,omitempty"`
	UserId      uint64                 `gorm:"index" json:"user_id,omitempty"`
	Name        string                 `gorm:"size:255;not null" json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Price       float64                `gorm:"type:decimal(10,2);not null" json:"price,omitempty"`
	CreatedAt   *time.Time             `gorm:"autoCreateTime" json:"created_at,omitempty"`
	Version     optimisticlock.Version `json:"-"`

	// Correct relationship for ProductStock
	ProductStock []*ProductStock `gorm:"foreignKey:ProductId" json:"product_stock,omitempty"`
}

type ProductWithStock struct {
	Product
	ProductStock         []*ProductStock `json:"product_stock,omitempty"`
	InactiveProductStock []*ProductStock `json:"inactive_product_stock,omitempty"`
	TotalStock           int             `json:"total_stock"`
	TotalInactiveStock   int             `json:"total_inactive_stock"`
}

type ProductStock struct {
	ProductId   uint64                 `gorm:"primaryKey" json:"product_id,omitempty"`
	WarehouseId uint64                 `gorm:"primaryKey" json:"warehouse_id,omitempty"`
	Quantity    int                    `gorm:"not null" json:"quantity,omitempty"`
	Product     *Product               `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE;" json:"product,omitempty"`
	Warehouse   *Warehouse             `gorm:"foreignKey:WarehouseId;constraint:OnDelete:CASCADE;" json:"warehouse,omitempty"`
	Version     optimisticlock.Version `json:"-"`
}
