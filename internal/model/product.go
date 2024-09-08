package model

import (
	"fmt"
	"gorm.io/plugin/optimisticlock"
	"time"
)

type Product struct {
	Id          uint64                 `json:"id,omitempty"`
	ShopUserId  uint64                 `json:"shop_user_id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Price       float64                `json:"price,omitempty"`
	CreatedAt   *time.Time             `json:"created_at,omitempty"`
	Version     optimisticlock.Version `json:"-"`

	// fk
	WarehouseProductStocks []*WarehouseProductStock `json:"warehouse_product_stocks,omitempty"`
}

func (m *Product) Validate() error {
	// validate name
	if m.Name == "" {
		return fmt.Errorf("product name is required")
	}

	// validate description
	if m.Description == "" {
		return fmt.Errorf("product description is required")
	}

	if len(m.Description) < 20 {
		return fmt.Errorf("product description must be at least 20 characters")
	}

	// validate price
	if m.Price == 0 {
		return fmt.Errorf("product price is required")
	}

	if m.Price < 0 {
		return fmt.Errorf("product price is invalid")
	}

	return nil
}
