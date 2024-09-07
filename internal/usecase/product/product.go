package product

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (u *productUsecaseImpl) GetProductsByShopUserIdAndWarehouseStatus(ctx context.Context, shopUserId uint64, status model.WarehouseStatus) ([]*model.Product, error) {
	if status != "" && status != model.WarehouseStatusActive && status != model.WarehouseStatusInactive {
		return nil, model.ErrUProductGetProductsByShopUserIdAndWarehouseStatus.New("invalid status")
	}
	products, err := u.rProduct.GetProductsByShopUserIdAndWarehouseStatus(ctx, shopUserId, status)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *productUsecaseImpl) CreateProduct(ctx context.Context, product *model.Product) (uint64, error) {
	err := product.Validate()
	if err != nil {
		return 0, model.ErrUProductCreateProduct.New(err)
	}

	product.Id = 0
	return u.rProduct.CreateProduct(ctx, product)
}
