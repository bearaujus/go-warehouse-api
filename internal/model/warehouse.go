package model

import (
	"fmt"
	"gorm.io/plugin/optimisticlock"
	"time"
)

type WarehouseStatus string

const (
	WarehouseStatusActive   WarehouseStatus = "active"
	WarehouseStatusInactive WarehouseStatus = "inactive"
)

type Warehouse struct {
	Id         uint64                 `json:"id,omitempty"`
	ShopUserId uint64                 `json:"shop_user_id,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Location   string                 `json:"location,omitempty"`
	Status     WarehouseStatus        `json:"status,omitempty"`
	CreatedAt  *time.Time             `json:"created_at,omitempty"`
	Version    optimisticlock.Version `json:"-"`
}

func (m *Warehouse) Validate() error {
	// validate name
	if m.Name == "" {
		return fmt.Errorf("warehouse name is required")
	}

	// validate location
	if m.Location == "" {
		return fmt.Errorf("warehouse location is required")
	}

	// validate status
	if m.Status != "" && m.Status != WarehouseStatusActive && m.Status != WarehouseStatusInactive {
		return fmt.Errorf("invalid warehouse status. warehouse status must be one of [%v, %v]", WarehouseStatusActive, WarehouseStatusInactive)
	}

	return nil
}

type WarehouseProductStock struct {
	Id          uint64                 `json:"id,omitempty"`
	ProductId   uint64                 `json:"product_id,omitempty"`
	WarehouseId uint64                 `json:"warehouse_id,omitempty"`
	Quantity    int                    `json:"quantity,omitempty"`
	Version     optimisticlock.Version `json:"-"`

	Product *Product `json:"product,omitempty"`
}

type WarehouseProductTransfer struct {
	Id              uint64     `json:"id,omitempty"`
	ProductId       uint64     `json:"product_id,omitempty"`
	FromWarehouseId uint64     `json:"from_warehouse_id,omitempty"`
	ToWarehouseId   uint64     `json:"to_warehouse_id,omitempty"`
	Quantity        int        `json:"quantity,omitempty"`
	TransferredAt   *time.Time `json:"transferred_at,omitempty"`
}
