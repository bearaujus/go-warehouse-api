package warehouse

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/resource/warehouse"
)

type WarehouseUsecase interface {
	GetWarehousesByUser(ctx context.Context, userId uint64) ([]*model.Warehouse, error)
	GetWarehousesByUserAndShop(ctx context.Context, userId, shopId uint64) ([]*model.Warehouse, error)
	CreateWarehouse(ctx context.Context, userId uint64, warehouse *model.Warehouse) (uint64, error)
	CreateWarehouseInboundTransaction(ctx context.Context, userId, id, productId uint64, quantity int) error
}

type warehouseUsecaseImpl struct {
	rWarehouse warehouse.WarehouseResource
}

func NewWarehouseUsecase(rWarehouse warehouse.WarehouseResource) WarehouseUsecase {
	return &warehouseUsecaseImpl{rWarehouse: rWarehouse}
}
