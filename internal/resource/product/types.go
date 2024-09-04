package product

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

type ProductResource interface {
	GetProductsWithStockByUser(ctx context.Context, userId uint64) ([]*model.ProductWithStock, error)
	CreateProduct(ctx context.Context, product *model.Product) (uint64, error)
}
