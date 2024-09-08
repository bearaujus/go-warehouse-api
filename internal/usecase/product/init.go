package product

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/resource/product"
)

type ProductUsecase interface {
	GetProductsByShopUserIdAndWarehouseStatus(ctx context.Context, shopUserId uint64, warehouseStatus model.WarehouseStatus) ([]*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) (uint64, error)
}

type productUsecaseImpl struct {
	rProduct product.ProductResource
}

func NewProductUsecase(rProduct product.ProductResource) ProductUsecase {
	return &productUsecaseImpl{rProduct: rProduct}
}
