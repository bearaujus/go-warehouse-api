package postgres

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *productResourcePostgresImpl) GetProductsByShopUserIdAndWarehouseStatus(ctx context.Context, shopUserId uint64, warehouseStatus model.WarehouseStatus) ([]*model.Product, error) {
	var products []*model.Product
	q := r.db.WithContext(ctx)
	if shopUserId != 0 {
		q = q.Where("shop_user_id = ?", shopUserId)
	}
	if warehouseStatus != "" {
		q = q.Preload("WarehouseProductStocks", "warehouse_id IN (SELECT id FROM warehouses WHERE status = ?)", warehouseStatus)
	} else {
		q = q.Preload("WarehouseProductStocks")
	}
	err := q.Find(&products).Error
	if err != nil {
		return nil, model.ErrRProductPostgresGetProductsByShopUserIdAndWarehouseStatus.New(err)
	}
	return products, nil
}

func (r *productResourcePostgresImpl) CreateProduct(ctx context.Context, product *model.Product) (uint64, error) {
	err := r.db.WithContext(ctx).Create(product).Error
	if err != nil {
		return 0, model.ErrRProductPostgresCreateProduct.New(err)
	}
	return product.Id, nil
}
