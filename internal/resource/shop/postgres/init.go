package postgres

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/shop"
	"gorm.io/gorm"
)

type shopResourcePostgresImpl struct {
	db *gorm.DB
}

func NewShopResourcePostgres(db *gorm.DB) shop.ShopResource {
	return &shopResourcePostgresImpl{db: db}
}
