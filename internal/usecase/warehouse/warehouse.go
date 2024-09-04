package warehouse

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (u *warehouseUsecaseImpl) GetWarehousesByUser(ctx context.Context, userId uint64) ([]*model.Warehouse, error) {
	return u.rWarehouse.GetWarehousesByUser(ctx, userId)
}

func (u *warehouseUsecaseImpl) GetWarehousesByUserAndShop(ctx context.Context, userId, shopId uint64) ([]*model.Warehouse, error) {
	return u.rWarehouse.GetWarehousesByUserAndShop(ctx, userId, shopId)
}

func (u *warehouseUsecaseImpl) CreateWarehouse(ctx context.Context, userId uint64, warehouse *model.Warehouse) (uint64, error) {
	// validate name
	if warehouse.Name == "" {
		return 0, model.ErrUWarehouseCreateWarehouse.New("name is required")
	}

	// validate location
	if warehouse.Location == "" {
		return 0, model.ErrUWarehouseCreateWarehouse.New("location is required")
	}

	warehouse.Id = 0
	warehouse.Status = model.WarehouseStatusActive
	return u.rWarehouse.CreateWarehouse(ctx, userId, warehouse)
}

func (u *warehouseUsecaseImpl) CreateWarehouseInboundTransaction(ctx context.Context, userId, id, productId uint64, quantity int) error {
	return u.rWarehouse.CreateWarehouseInboundTransaction(ctx, userId, id, productId, quantity)
}
