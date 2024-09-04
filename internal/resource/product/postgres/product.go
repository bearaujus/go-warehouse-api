package postgres

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *productResourcePostgresImpl) GetProductsWithStockByUser(ctx context.Context, userId uint64) ([]*model.ProductWithStock, error) {
	var products []*model.Product
	var productsWithStock []*model.ProductWithStock
	err := r.db.WithContext(ctx).Where("products.user_id = ?", userId).Preload("ProductStock.Warehouse").Find(&products).Error
	if err != nil {
		return nil, model.ErrRProductPostgresGetProductsWithStockByUser.New(err)
	}

	for _, product := range products {
		productWithStock := &model.ProductWithStock{
			Product: *product,
		}

		var totalStock, totalInactiveStock int
		var activeProductStock []*model.ProductStock
		var inactiveProductStock []*model.ProductStock

		for _, stock := range product.ProductStock {
			if stock.Warehouse.Status == model.WarehouseStatusActive {
				activeProductStock = append(activeProductStock, stock)
				totalStock += stock.Quantity
			} else {
				inactiveProductStock = append(inactiveProductStock, stock)
				totalInactiveStock += stock.Quantity
			}
		}

		productWithStock.ProductStock = activeProductStock
		productWithStock.InactiveProductStock = inactiveProductStock
		productWithStock.TotalStock = totalStock
		productWithStock.TotalInactiveStock = totalInactiveStock

		productsWithStock = append(productsWithStock, productWithStock)
	}

	return productsWithStock, nil
}

func (r *productResourcePostgresImpl) CreateProduct(ctx context.Context, product *model.Product) (uint64, error) {
	err := r.db.WithContext(ctx).Create(product).Error
	if err != nil {
		return 0, model.ErrRProductPostgresCreateProduct.New(err)
	}
	return product.Id, nil
}
