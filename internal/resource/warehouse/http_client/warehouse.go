package http_client

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *warehouseResourceHTTPClientImpl) GetWarehousesByUser(ctx context.Context, userId uint64) ([]*model.Warehouse, error) {
	return nil, model.ErrRWarehouseHTTPClientGetWarehousesByUser.New(model.ErrCommonNotImplemented)
}

func (r *warehouseResourceHTTPClientImpl) GetWarehousesByUserAndShop(ctx context.Context, userId, shopId uint64) ([]*model.Warehouse, error) {
	return nil, model.ErrRWarehouseHTTPClientGetWarehousesByUserAndShop.New(model.ErrCommonNotImplemented)
}

func (r *warehouseResourceHTTPClientImpl) CreateWarehouse(ctx context.Context, userId uint64, warehouse *model.Warehouse) (uint64, error) {
	return 0, model.ErrRWarehouseHTTPClientCreateWarehouse.New(model.ErrCommonNotImplemented)
}

func (r *warehouseResourceHTTPClientImpl) CreateWarehouseInboundTransaction(ctx context.Context, userId, id, productId uint64, quantity int) error {
	return model.ErrRWarehouseHTTPClientCreateWarehouseInboundTransaction.New(model.ErrCommonNotImplemented)
}
