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
	// validate shop user id
	if shopUserId == 0 {
		return model.ErrUWarehouseAddWarehouseProductStock.New("warehouse shop user id is required")
	}

	// validate warehouse id
	if id == 0 {
		return model.ErrUWarehouseAddWarehouseProductStock.New("warehouse id is required")
	}

	// validate product id
	if productId == 0 {
		return model.ErrUWarehouseAddWarehouseProductStock.New("product id is required")
	}

	// validate quantity
	if quantity == 0 {
		return model.ErrUWarehouseAddWarehouseProductStock.New("quantity id is required")
	}

	if quantity < 0 {
		return model.ErrUWarehouseAddWarehouseProductStock.New("quantity is invalid")
	}

	return u.rWarehouse.AddWarehouseProductStock(ctx, shopUserId, id, productId, quantity)
}

func (u *warehouseUsecaseImpl) TransferWarehouseProductStock(ctx context.Context, shopUserId, fromId, toId, productId uint64, quantity int) (*model.WarehouseProductTransfer, error) {
	// validate shop user id
	if shopUserId == 0 {
		return nil, model.ErrUWarehouseTransferWarehouseProductStock.New("warehouse shop user id is required")
	}

	// validate from id
	if fromId == 0 {
		return nil, model.ErrUWarehouseTransferWarehouseProductStock.New("source warehouse id is required")
	}

	// validate to id
	if toId == 0 {
		return nil, model.ErrUWarehouseTransferWarehouseProductStock.New("destination warehouse id is required")
	}

	// validate product id
	if productId == 0 {
		return nil, model.ErrUWarehouseTransferWarehouseProductStock.New("product id is required")
	}

	// validate quantity
	if quantity == 0 {
		return nil, model.ErrUWarehouseTransferWarehouseProductStock.New("quantity id is required")
	}

	if quantity < 0 {
		return nil, model.ErrUWarehouseTransferWarehouseProductStock.New("quantity is invalid")
	}

	// validate source and destination id
	if fromId == toId {
		return nil, model.ErrUWarehouseTransferWarehouseProductStock.New("transfer between the same warehouse is not allowed")
	}
	return u.rWarehouse.TransferWarehouseProductStock(ctx, shopUserId, fromId, toId, productId, quantity)
}
