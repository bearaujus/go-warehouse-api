package model

import (
	"gorm.io/plugin/optimisticlock"
	"time"
)

type Warehouse struct {
	Id        uint64                 `gorm:"primaryKey" json:"id,omitempty"`
	ShopId    uint64                 `gorm:"index" json:"shop_id,omitempty"`
	Name      string                 `gorm:"size:255;not null" json:"name,omitempty"`
	Location  string                 `json:"location,omitempty"`
	Status    WarehouseStatus        `json:"status,omitempty"`
	CreatedAt *time.Time             `gorm:"autoCreateTime" json:"created_at,omitempty"`
	Version   optimisticlock.Version `json:"-"`
}

type WarehouseTransfer struct {
	Id              uint64                 `gorm:"primaryKey" json:"id,omitempty"`
	ProductId       uint64                 `gorm:"index" json:"product_id,omitempty"`
	FromWarehouseId uint64                 `gorm:"index" json:"from_warehouse_id,omitempty"`
	ToWarehouseId   uint64                 `gorm:"index" json:"to_warehouse_id,omitempty"`
	Quantity        int                    `gorm:"not null" json:"quantity,omitempty"`
	TransferredAt   *time.Time             `gorm:"autoCreateTime" json:"transferred_at,omitempty"`
	Product         *Product               `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE;" json:"product,omitempty"`
	FromWarehouse   *Warehouse             `gorm:"foreignKey:FromWarehouseId;constraint:OnDelete:SET NULL;" json:"from_warehouse,omitempty"`
	ToWarehouse     *Warehouse             `gorm:"foreignKey:ToWarehouseId;constraint:OnDelete:SET NULL;" json:"to_warehouse,omitempty"`
	Version         optimisticlock.Version `json:"-"`
}

type WarehouseStatus string

const (
	WarehouseStatusActive   WarehouseStatus = "active"
	WarehouseStatusInactive WarehouseStatus = "inactive"
)
