package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"gorm.io/gorm"
)

func (r *warehouseResourcePostgresImpl) GetWarehousesByShopUserId(ctx context.Context, shopUserId uint64) ([]*model.Warehouse, error) {
	var warehouses []*model.Warehouse
	err := r.db.WithContext(ctx).Where("shop_user_id = ?", shopUserId).Find(&warehouses).Error
	if err != nil {
		return nil, model.ErrRWarehousePostgresGetWarehousesByShopUserId.New(err)
	}
	return warehouses, nil
}

func (r *warehouseResourcePostgresImpl) GetActiveWarehouseProductStocksByProductId(ctx context.Context, productId uint64) ([]*model.WarehouseProductStock, error) {
	var warehouseProductStocks []*model.WarehouseProductStock
	err := r.db.WithContext(ctx).
		Joins("JOIN warehouses ON warehouses.id = warehouse_product_stocks.warehouse_id").
		Where("warehouse_product_stocks.product_id = ? AND warehouses.status = ?", productId, model.WarehouseStatusActive).
		Find(&warehouseProductStocks).Error
	if err != nil {
		return nil, model.ErrRWarehousePostgresGetActiveWarehouseProductStocksByProductId.New(err)
	}
	return warehouseProductStocks, nil
}

func (r *warehouseResourcePostgresImpl) GetWarehouseProductStocksByShopUserIdAndProductId(ctx context.Context, shopUserId, productId uint64) ([]*model.WarehouseProductStock, error) {
	var warehouseProductStocks []*model.WarehouseProductStock
	err := r.db.WithContext(ctx).
		Joins("JOIN warehouses ON warehouses.id = warehouse_product_stocks.warehouse_id").
		Where("warehouses.user_id = ? AND warehouse_product_stocks.product_id = ?", shopUserId, productId).
		Find(&warehouseProductStocks).Error
	if err != nil {
		return nil, model.ErrRWarehousePostgresGetWarehouseProductStocksByShopUserIdAndProductId.New(err)
	}
	return warehouseProductStocks, nil
}

func (r *warehouseResourcePostgresImpl) CreateWarehouse(ctx context.Context, warehouse *model.Warehouse) (uint64, error) {
	err := r.db.WithContext(ctx).Create(warehouse).Error
	if err != nil {
		return 0, model.ErrRWarehousePostgresCreateWarehouse.New(err)
	}

	return warehouse.Id, nil
}

func (r *warehouseResourcePostgresImpl) UpdateWarehouse(ctx context.Context, warehouse *model.Warehouse) error {
	err := r.db.WithContext(ctx).Where("id = ? AND shop_user_id = ?", warehouse.Id, warehouse.ShopUserId).Updates(warehouse).Error
	if err != nil {
		return model.ErrRWarehousePostgresUpdateWarehouse.New(err)
	}

	return nil
}

func (r *warehouseResourcePostgresImpl) AddWarehouseProductStock(ctx context.Context, shopUserId, id, productId uint64, quantity int) error {
	var warehouse model.Warehouse
	err := r.db.WithContext(ctx).Model(&warehouse).Where("shop_user_id = ? AND id = ?", shopUserId, id).First(&warehouse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrRWarehousePostgresAddWarehouseProductStock.New(fmt.Errorf("warehouse is not found or not belongs to the user"))
		}
		return model.ErrRWarehousePostgresAddWarehouseProductStock.New(err)
	}

	var warehouseProductStock model.WarehouseProductStock
	err = r.db.WithContext(ctx).Where("product_id = ? AND warehouse_id = ?", productId, warehouse.Id).First(&warehouseProductStock).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrRWarehousePostgresAddWarehouseProductStock.New(err)
		}
		warehouseProductStock = model.WarehouseProductStock{
			ProductId:   productId,
			WarehouseId: warehouse.Id,
			Quantity:    quantity,
		}
		if err = r.db.WithContext(ctx).Create(&warehouseProductStock).Error; err != nil {
			return model.ErrRWarehousePostgresAddWarehouseProductStock.New(err)
		}
		return nil
	}

	warehouseProductStock.Quantity += quantity
	if err = r.db.WithContext(ctx).Where("id = ?", warehouseProductStock.Id).Updates(&warehouseProductStock).Error; err != nil {
		return model.ErrRWarehousePostgresAddWarehouseProductStock.New(err)
	}

	return nil
}

func (r *warehouseResourcePostgresImpl) TransferWarehouseProductStock(ctx context.Context, shopUserId, fromId, toId, productId uint64, quantity int) (*model.WarehouseProductTransfer, error) {
	// validate src warehouse ownership
	var srcWarehouse model.Warehouse
	err := r.db.WithContext(ctx).Model(&srcWarehouse).Where("shop_user_id = ? AND id = ?", shopUserId, fromId).First(&srcWarehouse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(fmt.Errorf("source warehouse is not found or not belongs to the user"))
		}
		return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(err)
	}

	// validate dest warehouse ownership
	var destWarehouse model.Warehouse
	err = r.db.WithContext(ctx).Model(&destWarehouse).Where("shop_user_id = ? AND id = ?", shopUserId, toId).First(&destWarehouse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(fmt.Errorf("destination warehouse is not found or not belongs to the user"))
		}
		return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(err)
	}

	// get and validate warehouse product stock in src warehouse
	var srcWarehouseProductStock model.WarehouseProductStock
	err = r.db.WithContext(ctx).Where("product_id = ? AND warehouse_id = ?", productId, fromId).First(&srcWarehouseProductStock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(fmt.Errorf("warehouse product stock at source warehouse is not found"))
		}
		return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(err)
	}

	if srcWarehouseProductStock.Quantity-quantity < 0 {
		return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(fmt.Errorf("warehouse product stock at source warehouse is not enough"))
	}

	// create new transaction
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// reduce warehouse product stock from src warehouse
	srcWarehouseProductStock.Quantity -= quantity
	if err = tx.Where("id = ?", srcWarehouseProductStock.Id).Updates(&srcWarehouseProductStock).Error; err != nil {
		return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(err)
	}

	// add or create warehouse product stock in dest warehouse
	var destWarehouseProductStock model.WarehouseProductStock
	err = r.db.WithContext(ctx).Where("product_id = ? AND warehouse_id = ?", productId, toId).First(&destWarehouseProductStock).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(err)
		}
		destWarehouseProductStock = model.WarehouseProductStock{
			ProductId:   productId,
			WarehouseId: toId,
			Quantity:    quantity,
		}
		if err = tx.Create(&destWarehouseProductStock).Error; err != nil {
			return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(err)
		}
	} else {
		destWarehouseProductStock.Quantity += quantity
		if err = tx.Where("id = ?", destWarehouseProductStock.Id).Updates(&destWarehouseProductStock).Error; err != nil {
			return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(err)
		}
	}

	// create warehouse transfer log
	warehouseProductTransfer := model.WarehouseProductTransfer{
		ProductId:       productId,
		FromWarehouseId: fromId,
		ToWarehouseId:   toId,
		Quantity:        quantity,
	}

	err = tx.Create(&warehouseProductTransfer).Error
	if err != nil {
		return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(err)
	}

	// commit transaction
	err = tx.Commit().Error
	if err != nil {
		return nil, model.ErrRWarehousePostgresTransferWarehouseProductStock.New(err)
	}

	return &warehouseProductTransfer, nil
}
