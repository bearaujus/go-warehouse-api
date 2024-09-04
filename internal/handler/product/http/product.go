package http

import (
	"context"
	"encoding/json"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func (h *productHandlerHTTPImpl) GetProductsWithStock(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHProductHTTPGetProductsWithStock.New(model.ErrCommonInvalidAuthToken))
		return
	}

	products, err := h.uProduct.GetProductsWithStockByUser(ctx, userId)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusOK, products)
}

func (h *productHandlerHTTPImpl) CreateProduct(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHProductHTTPCreateProduct.New(model.ErrCommonInvalidAuthToken))
		return
	}

	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHProductHTTPCreateProduct.New(model.ErrCommonInvalidRequestBody))
		return
	}

	product := model.Product{}
	err = json.Unmarshal(data, &product)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHProductHTTPCreateProduct.New(model.ErrCommonInvalidRequestBody))
		return
	}

	product.UserId = userId
	id, err := h.uProduct.CreateProduct(ctx, &product)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusCreated, &model.Product{Id: id})
}
