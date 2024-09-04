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

func (h *warehouseHandlerHTTPImpl) GetWarehouses(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHWarehouseHTTPGetWarehouses.New(model.ErrCommonInvalidAuthToken))
		return
	}

	shopId, _ := pkg.StringToUint64(rCtx.Query("shop_id"))
	if shopId != 0 {
		warehouses, err := h.uWarehouse.GetWarehousesByUserAndShop(ctx, userId, shopId)
		if err != nil {
			httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
			return
		}

		httputil.WriteResponse(rCtx, http.StatusOK, warehouses)
		return
	}

	warehouses, err := h.uWarehouse.GetWarehousesByUser(ctx, userId)
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

	id, err := h.uWarehouse.CreateWarehouse(ctx, userId, &warehouse)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusCreated, &model.Warehouse{Id: id})
}

func (h *warehouseHandlerHTTPImpl) CreateWarehouseInboundTransaction(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHWarehouseHTTPCreateWarehouseInboundTransaction.New(model.ErrCommonInvalidAuthToken))
		return
	}

	warehouseId, err := pkg.StringToUint64(rCtx.Param("id"))
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPCreateWarehouseInboundTransaction.New(model.ErrCommonInvalidRequestURL))
		return
	}

	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPCreateWarehouseInboundTransaction.New(model.ErrCommonInvalidRequestBody))
		return
	}

	productStock := model.ProductStock{}
	err = json.Unmarshal(data, &productStock)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHWarehouseHTTPCreateWarehouseInboundTransaction.New(model.ErrCommonInvalidRequestBody))
		return
	}

	err = h.uWarehouse.CreateWarehouseInboundTransaction(ctx, userId, warehouseId, productStock.ProductId, productStock.Quantity)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteEmptyResponse(rCtx, http.StatusCreated)
}
