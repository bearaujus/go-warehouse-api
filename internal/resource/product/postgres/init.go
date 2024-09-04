package postgres

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/product"
	"gorm.io/gorm"
)

type productResourcePostgresImpl struct {
	db *gorm.DB
}

func NewProductResourcePostgres(db *gorm.DB) product.ProductResource {
	return &productResourcePostgresImpl{db: db}
}
