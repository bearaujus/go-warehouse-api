package http

import (
	"context"
	"encoding/json"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func (h *warehouseHandlerHTTPImpl) GetWarehousesByShopUserId(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHWarehouseHTTPGetWarehousesByShopUserId.New(model.ErrCommonInvalidAuthToken))
		return
	}

	warehouses, err := h.uWarehouse.GetWarehousesByShopUserId(ctx, userId)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusOK, warehouses)
}

func (h *warehouseHandlerHTTPImpl) CreateWarehouse(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHWarehouseHTTPCreateWarehouse.New(model.ErrCommonInvalidAuthToken))
		return
	}

	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPCreateWarehouse.New(model.ErrCommonInvalidRequestBody))
		return
	}

	warehouse := model.Warehouse{}
	err = json.Unmarshal(data, &warehouse)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPCreateWarehouse.New(model.ErrCommonInvalidRequestBody))
		return
	}

	warehouse.ShopUserId = userId
	id, err := h.uWarehouse.CreateWarehouse(ctx, &warehouse)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusCreated, &model.Warehouse{Id: id})
}

func (h *warehouseHandlerHTTPImpl) UpdateWarehouse(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHWarehouseHTTPUpdateWarehouse.New(model.ErrCommonInvalidAuthToken))
		return
	}

	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPUpdateWarehouse.New(model.ErrCommonInvalidRequestBody))
		return
	}

	id, err := pkg.StringToUint64(rCtx.Param("id"))
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPUpdateWarehouse.New(model.ErrCommonInvalidRequestURL))
		return
	}

	warehouse := model.Warehouse{}
	err = json.Unmarshal(data, &warehouse)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPUpdateWarehouse.New(model.ErrCommonInvalidRequestBody))
		return
	}

	warehouse.Id = id
	warehouse.ShopUserId = userId
	err = h.uWarehouse.UpdateWarehouse(ctx, &warehouse)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusOK, warehouse)
}

func (h *warehouseHandlerHTTPImpl) AddWarehouseProductStock(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHWarehouseHTTPAddWarehouseProductStock.New(model.ErrCommonInvalidAuthToken))
		return
	}

	id, err := pkg.StringToUint64(rCtx.Param("id"))
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPAddWarehouseProductStock.New(model.ErrCommonInvalidRequestURL))
		return
	}

	productId, err := pkg.StringToUint64(rCtx.Param("product_id"))
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPAddWarehouseProductStock.New(model.ErrCommonInvalidRequestURL))
		return
	}

	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPAddWarehouseProductStock.New(model.ErrCommonInvalidRequestBody))
		return
	}

	type addWarehouseProductStockReq struct {
		Quantity int `json:"quantity"`
	}

	addWarehouseProductStock := addWarehouseProductStockReq{}
	err = json.Unmarshal(data, &addWarehouseProductStock)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPAddWarehouseProductStock.New(model.ErrCommonInvalidRequestBody))
		return
	}

	err = h.uWarehouse.AddWarehouseProductStock(ctx, userId, id, productId, addWarehouseProductStock.Quantity)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteEmptyResponse(rCtx, http.StatusOK)
}

func (h *warehouseHandlerHTTPImpl) TransferWarehouseProductStock(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHWarehouseHTTPTransferWarehouseProductStock.New(model.ErrCommonInvalidAuthToken))
		return
	}

	fromId, err := pkg.StringToUint64(rCtx.Param("id"))
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPTransferWarehouseProductStock.New(model.ErrCommonInvalidRequestURL))
		return
	}

	productId, err := pkg.StringToUint64(rCtx.Param("product_id"))
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPTransferWarehouseProductStock.New(model.ErrCommonInvalidRequestURL))
		return
	}

	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPTransferWarehouseProductStock.New(model.ErrCommonInvalidRequestBody))
		return
	}

	type transferWarehouseProductStockReq struct {
		DestinationWarehouseId uint64 `json:"destination_warehouse_id"`
		Quantity               int    `json:"quantity"`
	}

	transferWarehouseProductStock := transferWarehouseProductStockReq{}
	err = json.Unmarshal(data, &transferWarehouseProductStock)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPTransferWarehouseProductStock.New(model.ErrCommonInvalidRequestBody))
		return
	}

	warehouseProductTransfer, err := h.uWarehouse.TransferWarehouseProductStock(ctx, userId, fromId, transferWarehouseProductStock.DestinationWarehouseId, productId, transferWarehouseProductStock.Quantity)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusOK, warehouseProductTransfer)
}
