package http_client

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *warehouseResourceHTTPClientImpl) GetWarehousesByShopUserId(ctx context.Context, shopUserId uint64) ([]*model.Warehouse, error) {
	return nil, model.ErrRWarehouseHTTPClientGetWarehousesByShopUserId.New(model.ErrCommonNotImplemented)
}

func (r *warehouseResourceHTTPClientImpl) GetActiveWarehouseProductStocksByProductId(ctx context.Context, productId uint64) ([]*model.WarehouseProductStock, error) {
	return nil, model.ErrRWarehouseHTTPClientGetActiveWarehouseProductStocksByProductId.New(model.ErrCommonNotImplemented)
}

func (r *warehouseResourceHTTPClientImpl) GetWarehouseProductStocksByShopUserIdAndProductId(ctx context.Context, shopUserId, productId uint64) ([]*model.WarehouseProductStock, error) {
	return nil, model.ErrRWarehouseHTTPClientGetWarehouseProductStocksByShopUserIdAndProductId.New(model.ErrCommonNotImplemented)
}

func (r *warehouseResourceHTTPClientImpl) CreateWarehouse(ctx context.Context, warehouse *model.Warehouse) (uint64, error) {
	return 0, model.ErrRWarehouseHTTPClientCreateWarehouse.New(model.ErrCommonNotImplemented)
}

func (r *warehouseResourceHTTPClientImpl) UpdateWarehouse(ctx context.Context, warehouse *model.Warehouse) error {
	return model.ErrRWarehouseHTTPClientUpdateWarehouse.New(model.ErrCommonNotImplemented)
}

func (r *warehouseResourceHTTPClientImpl) AddWarehouseProductStock(ctx context.Context, shopUserId, id, productId uint64, quantity int) error {
	return model.ErrRWarehouseHTTPClientAddWarehouseProductStock.New(model.ErrCommonNotImplemented)
}

func (r *warehouseResourceHTTPClientImpl) TransferWarehouseProductStock(ctx context.Context, shopUserId, fromId, toId, productId uint64, quantity int) (*model.WarehouseProductTransfer, error) {
	return nil, model.ErrRWarehouseHTTPClientTransferWarehouseProductStock.New(model.ErrCommonNotImplemented)
}
