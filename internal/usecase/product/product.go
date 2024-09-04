package product

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (u *productUsecaseImpl) GetProductsWithStockByUser(ctx context.Context, userId uint64) ([]*model.ProductWithStock, error) {
	return u.rProduct.GetProductsWithStockByUser(ctx, userId)
}

func (u *productUsecaseImpl) CreateProduct(ctx context.Context, product *model.Product) (uint64, error) {
	// validate name
	if product.Name == "" {
		return 0, model.ErrUProductCreateProduct.New("name is required")
	}

	// validate description
	if product.Description == "" {
		return 0, model.ErrUProductCreateProduct.New("description is required")
	}

	if len(product.Description) < 20 {
		return 0, model.ErrUProductCreateProduct.New("description must be at least 20 characters")
	}

	// validate price
	if product.Price == 0 {
		return 0, model.ErrUProductCreateProduct.New("price is required")
	}

	if product.Price < 0 {
		return 0, model.ErrUProductCreateProduct.New("price is invalid")
	}

	product.Id = 0
	return u.rProduct.CreateProduct(ctx, product)
}
