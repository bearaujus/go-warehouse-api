package warehouse

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

type WarehouseResource interface {
	GetWarehousesByShopUserId(ctx context.Context, shopUserId uint64) ([]*model.Warehouse, error)
	GetActiveWarehouseProductStocksByProductId(ctx context.Context, productId uint64) ([]*model.WarehouseProductStock, error)
	GetWarehouseProductStocksByShopUserIdAndProductId(ctx context.Context, shopUserId, productId uint64) ([]*model.WarehouseProductStock, error)
	CreateWarehouse(ctx context.Context, warehouse *model.Warehouse) (uint64, error)
	UpdateWarehouse(ctx context.Context, warehouse *model.Warehouse) error
	AddWarehouseProductStock(ctx context.Context, shopUserId, id, productId uint64, quantity int) error
	TransferWarehouseProductStock(ctx context.Context, shopUserId, fromId, toId, productId uint64, quantity int) (*model.WarehouseProductTransfer, error)
}
