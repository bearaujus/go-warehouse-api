package postgres

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/warehouse"
	"gorm.io/gorm"
)

type warehouseResourcePostgresImpl struct {
	db *gorm.DB
}

func NewWarehouseResourcePostgres(db *gorm.DB) warehouse.WarehouseResource {
	return &warehouseResourcePostgresImpl{db: db}
}
