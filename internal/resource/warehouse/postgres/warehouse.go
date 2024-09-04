package postgres

import (
	"context"
	"errors"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"gorm.io/gorm"
)

func (r *warehouseResourcePostgresImpl) GetWarehousesByUser(ctx context.Context, userId uint64) ([]*model.Warehouse, error) {
	var warehouses []*model.Warehouse
	err := r.db.WithContext(ctx).Joins("JOIN shops ON shops.id = warehouses.shop_id").Where("shops.user_id = ?", userId).Find(&warehouses).Error
	if err != nil {
		return nil, model.ErrRWarehousePostgresGetWarehousesByUser.New(err)
	}
	return warehouses, nil
}

func (r *warehouseResourcePostgresImpl) GetWarehousesByUserAndShop(ctx context.Context, userId, shopId uint64) ([]*model.Warehouse, error) {
	var warehouses []*model.Warehouse
	err := r.db.WithContext(ctx).Joins("JOIN shops ON shops.id = warehouses.shop_id").Where("shops.user_id = ? AND shop_id = ?", userId, shopId).Find(&warehouses).Error
	if err != nil {
		return nil, model.ErrRWarehousePostgresGetWarehousesByUserAndShop.New(err)
	}
	return warehouses, nil
}

func (r *warehouseResourcePostgresImpl) CreateWarehouse(ctx context.Context, userId uint64, warehouse *model.Warehouse) (uint64, error) {
	var shop model.Shop
	err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", warehouse.ShopId, userId).First(&shop).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, model.ErrRWarehousePostgresCreateWarehouse.New("shop warehouse not found or does not belong to the user")
		}
		return 0, model.ErrRWarehousePostgresCreateWarehouse.New(err)
	}

	err = r.db.WithContext(ctx).Create(warehouse).Error
	if err != nil {
		return 0, model.ErrRWarehousePostgresCreateWarehouse.New(err)
	}

	return warehouse.Id, nil
}

func (r *warehouseResourcePostgresImpl) CreateWarehouseInboundTransaction(ctx context.Context, userId, id, productId uint64, quantity int) error {
	var warehouse model.Warehouse
	err := r.db.WithContext(ctx).Model(&warehouse).Where("id = ?", id).First(&warehouse).Error
	if err != nil {
		return model.ErrRWarehousePostgresCreateWarehouseInboundTransaction.New(err)
	}

	var shop model.Shop
	err = r.db.WithContext(ctx).Where("id = ? AND user_id = ?", warehouse.ShopId, userId).First(&shop).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrRWarehousePostgresCreateWarehouseInboundTransaction.New("shop warehouse not found or does not belong to the user")
		}
		return model.ErrRWarehousePostgresCreateWarehouseInboundTransaction.New(err)
	}

	var productStock model.ProductStock
	err = r.db.WithContext(ctx).
		Where("product_id = ? AND warehouse_id = ?", productId, warehouse.Id).
		First(&productStock).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrRWarehousePostgresCreateWarehouseInboundTransaction.New(err)
		}
		productStock = model.ProductStock{
			ProductId:   productId,
			WarehouseId: warehouse.Id,
			Quantity:    quantity,
		}
		if err = r.db.WithContext(ctx).Create(&productStock).Error; err != nil {
			return model.ErrRWarehousePostgresCreateWarehouseInboundTransaction.New(err)
		}
		return nil
	}

	productStock.Quantity += quantity
	if err = r.db.WithContext(ctx).Save(&productStock).Error; err != nil {
		return model.ErrRWarehousePostgresCreateWarehouseInboundTransaction.New(err)
	}

	return nil
}
