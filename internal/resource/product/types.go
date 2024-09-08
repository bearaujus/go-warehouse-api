package product

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

type ProductResource interface {
	GetProductsByShopUserIdAndWarehouseStatus(ctx context.Context, shopUserId uint64, warehouseStatus model.WarehouseStatus) ([]*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) (uint64, error)
}
