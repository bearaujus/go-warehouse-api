package warehouse

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (u *warehouseUsecaseImpl) GetWarehousesByShopUserId(ctx context.Context, shopUserId uint64) ([]*model.Warehouse, error) {
	return u.rWarehouse.GetWarehousesByShopUserId(ctx, shopUserId)
}

func (u *warehouseUsecaseImpl) GetActiveWarehouseProductStocksByProductId(ctx context.Context, productId uint64) ([]*model.WarehouseProductStock, error) {
	return u.rWarehouse.GetActiveWarehouseProductStocksByProductId(ctx, productId)
}

func (u *warehouseUsecaseImpl) GetWarehouseProductStocksByShopUserIdAndProductId(ctx context.Context, shopUserId, productId uint64) ([]*model.WarehouseProductStock, error) {
	return u.rWarehouse.GetWarehouseProductStocksByShopUserIdAndProductId(ctx, shopUserId, productId)
}

func (u *warehouseUsecaseImpl) CreateWarehouse(ctx context.Context, warehouse *model.Warehouse) (uint64, error) {
	warehouse.Status = model.WarehouseStatusActive
	err := warehouse.Validate()
	if err != nil {
		return 0, model.ErrUWarehouseCreateWarehouse.New(err)
	}

	warehouse.Id = 0
	warehouse.CreatedAt = nil
	return u.rWarehouse.CreateWarehouse(ctx, warehouse)
}

func (u *warehouseUsecaseImpl) UpdateWarehouse(ctx context.Context, warehouse *model.Warehouse) error {
	err := warehouse.Validate()
	if err != nil {
		return model.ErrUWarehouseUpdateWarehouse.New(err)
	}

	warehouse.CreatedAt = nil
	return u.rWarehouse.UpdateWarehouse(ctx, warehouse)
}

func (u *warehouseUsecaseImpl) AddWarehouseProductStock(ctx context.Context, shopUserId, id, productId uint64, quantity int) error {
	return u.rWarehouse.AddWarehouseProductStock(ctx, shopUserId, id, productId, quantity)
}

func (u *warehouseUsecaseImpl) TransferWarehouseProductStock(ctx context.Context, shopUserId, fromId, toId, productId uint64, quantity int) (*model.WarehouseProductTransfer, error) {
	return u.rWarehouse.TransferWarehouseProductStock(ctx, shopUserId, fromId, toId, productId, quantity)
}
