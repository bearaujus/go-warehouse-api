package postgres

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/order"
	"gorm.io/gorm"
)

type orderResourcePostgresImpl struct {
	db *gorm.DB
}

func NewOrderResourcePostgres(db *gorm.DB) order.OrderResource {
	return &orderResourcePostgresImpl{db: db}
}
