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

func (h *orderHandlerHTTPImpl) GetOrdersByUserIdAndStatus(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHOrderHTTPGetOrdersByUserIdAndStatus.New(model.ErrCommonInvalidAuthToken))
		return
	}

	orders, err := h.uOrder.GetOrdersByUserIdAndStatus(ctx, userId, model.OrderStatus(rCtx.Query("status")))
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	for i := range orders {
		orders[i].UserId = 0
		for j := range orders[i].OrderItems {
			orders[i].OrderItems[j].Id = 0
			orders[i].OrderItems[j].OrderId = 0
		}
	}

	httputil.WriteResponse(rCtx, http.StatusOK, orders)
}

func (h *orderHandlerHTTPImpl) CreateOrder(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHOrderHTTPCreateOrder.New(model.ErrCommonInvalidAuthToken))
		return
	}

	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHOrderHTTPCreateOrder.New(model.ErrCommonInvalidRequestBody))
		return
	}

	type createOrderReq struct {
		ProductId uint64 `json:"product_id,omitempty"`
		Quantity  int    `json:"quantity,omitempty"`
	}

	var createOrder []*createOrderReq
	err = json.Unmarshal(data, &createOrder)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHOrderHTTPCreateOrder.New(model.ErrCommonInvalidRequestBody))
		return
	}

	var orderItems []*model.OrderItem
	for _, item := range createOrder {
		orderItems = append(orderItems, &model.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	order, err := h.uOrder.CreateOrder(ctx, userId, orderItems)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	order.UserId = 0
	for i := range order.OrderItems {
		order.OrderItems[i].Id = 0
		order.OrderItems[i].OrderId = 0
	}

	httputil.WriteResponse(rCtx, http.StatusCreated, order)
}

func (h *orderHandlerHTTPImpl) CompleteOrder(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHOrderHTTPCompleteOrder.New(model.ErrCommonInvalidAuthToken))
		return
	}

	id, err := pkg.StringToUint64(rCtx.Param("id"))
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHOrderHTTPCompleteOrder.New(model.ErrCommonInvalidRequestURL))
		return
	}

	order, err := h.uOrder.CompleteOrder(ctx, userId, id)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	order.UserId = 0
	for i := range order.OrderItems {
		order.OrderItems[i].Id = 0
		order.OrderItems[i].OrderId = 0
	}

	httputil.WriteResponse(rCtx, http.StatusOK, order)
}
