package http

import (
	"context"
	"encoding/json"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func (h *shopHandlerHTTPImpl) CreateShop(ctx context.Context, rCtx *app.RequestContext) {
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

	err = h.uShop.CreateShop(ctx, &shop)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteEmptyResponse(rCtx, http.StatusCreated)
}
