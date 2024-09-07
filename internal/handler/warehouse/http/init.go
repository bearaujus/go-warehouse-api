package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/handler"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/warehouse"
)

type warehouseHandlerHTTPImpl struct {
	uWarehouse warehouse.WarehouseUsecase
}

func NewWarehouseHandlerHTTP(uWarehouse warehouse.WarehouseUsecase) handler.HTTPHandler {
	return &warehouseHandlerHTTPImpl{uWarehouse: uWarehouse}
}
