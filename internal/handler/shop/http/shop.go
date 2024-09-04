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

func (h *shopHandlerHTTPImpl) GetShops(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHShopHTTPGetShops.New(model.ErrCommonInvalidAuthToken))
		return
	}

	shops, err := h.uShop.GetShopsByUser(ctx, userId)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusOK, shops)
}

func (h *shopHandlerHTTPImpl) CreateShop(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHShopHTTPCreateShop.New(model.ErrCommonInvalidAuthToken))
		return
	}

	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHShopHTTPCreateShop.New(model.ErrCommonInvalidRequestBody))
		return
	}

	shop := model.Shop{}
	err = json.Unmarshal(data, &shop)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHShopHTTPCreateShop.New(model.ErrCommonInvalidRequestBody))
		return
	}

	shop.UserId = userId
	id, err := h.uShop.CreateShop(ctx, &shop)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusCreated, &model.Shop{Id: id})
}
