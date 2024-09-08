package postgres

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/order"
	"gorm.io/gorm"
	"time"
)

type orderResourcePostgresImpl struct {
	db                 *gorm.DB
	orderExpirationTTL time.Duration
}

func NewOrderResourcePostgres(db *gorm.DB, orderExpirationTTL time.Duration) order.OrderResource {
	return &orderResourcePostgresImpl{db: db, orderExpirationTTL: orderExpirationTTL}
}
