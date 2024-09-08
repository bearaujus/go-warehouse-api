package http

import (
	"context"
	"encoding/json"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/auth"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"time"
)

func (h *productHandlerHTTPImpl) GetSellerProductsByShopUserId(ctx context.Context, rCtx *app.RequestContext) {
	userId, err := auth.GetUserIdFromContext(rCtx)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrHProductHTTPGetSellerProductsByShopUserId.New(model.ErrCommonInvalidAuthToken))
		return
	}

	activeProducts, err := h.uProduct.GetProductsByShopUserIdAndWarehouseStatus(ctx, userId, model.WarehouseStatusActive)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
	}

	inactiveProducts, err := h.uProduct.GetProductsByShopUserIdAndWarehouseStatus(ctx, userId, model.WarehouseStatusInactive)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
	}

	type getSellerProductsByShopUserIdResp struct {
		Id                uint64                         `json:"id,omitempty"`
		Name              string                         `json:"name,omitempty"`
		Description       string                         `json:"description,omitempty"`
		Price             float64                        `json:"price,omitempty"`
		CreatedAt         *time.Time                     `json:"created_at,omitempty"`
		Stock             int                            `json:"stock"`
		StockData         []*model.WarehouseProductStock `json:"stock_data"`
		InactiveStock     int                            `json:"inactive_stock"`
		InactiveStockData []*model.WarehouseProductStock `json:"inactive_stock_data"`
	}

	tmpFunc := func(m map[uint64]*getSellerProductsByShopUserIdResp, products []*model.Product, isActive bool) {
		for _, product := range products {
			d, ok := m[product.Id]
			if !ok {
				d = &getSellerProductsByShopUserIdResp{
					Id:                product.Id,
					Name:              product.Name,
					Description:       product.Description,
					Price:             product.Price,
					CreatedAt:         product.CreatedAt,
					Stock:             0,
					StockData:         make([]*model.WarehouseProductStock, 0),
					InactiveStock:     0,
					InactiveStockData: make([]*model.WarehouseProductStock, 0),
				}
				m[product.Id] = d
			}
			for _, warehouseProductStock := range product.WarehouseProductStocks {
				if warehouseProductStock.Quantity == 0 {
					continue
				}
				warehouseProductStock.Id = 0
				warehouseProductStock.ProductId = 0
				if isActive {
					d.Stock += warehouseProductStock.Quantity
					d.StockData = append(d.StockData, warehouseProductStock)
				} else {
					d.InactiveStock += warehouseProductStock.Quantity
					d.InactiveStockData = append(d.InactiveStockData, warehouseProductStock)
				}
			}
		}
	}

	tmp := map[uint64]*getSellerProductsByShopUserIdResp{}
	tmpFunc(tmp, activeProducts, true)
	tmpFunc(tmp, inactiveProducts, false)

	var resp []*getSellerProductsByShopUserIdResp
	for _, d := range tmp {
		resp = append(resp, d)
	}

	httputil.WriteResponse(rCtx, http.StatusOK, resp)
}

func (h *productHandlerHTTPImpl) GetBuyerProductsByShopUserId(ctx context.Context, rCtx *app.RequestContext) {
	products, err := h.uProduct.GetProductsByShopUserIdAndWarehouseStatus(ctx, 0, model.WarehouseStatusActive)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
	}

	type getBuyerProductsByShopUserIdResp struct {
		Id          uint64  `json:"id,omitempty"`
		ShopId      uint64  `json:"shop_id,omitempty"`
		Name        string  `json:"name,omitempty"`
		Description string  `json:"description,omitempty"`
		Price       float64 `json:"price,omitempty"`
		Stock       int     `json:"stock"`
	}

	tmp := map[uint64]*getBuyerProductsByShopUserIdResp{}
	for _, product := range products {
		d, ok := tmp[product.Id]
		if !ok {
			d = &getBuyerProductsByShopUserIdResp{
				Id:          product.Id,
				ShopId:      product.ShopUserId,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Stock:       0,
			}
			tmp[product.Id] = d
		}
		for _, warehouseProductStock := range product.WarehouseProductStocks {
			d.Stock += warehouseProductStock.Quantity
		}
	}

	var resp []*getBuyerProductsByShopUserIdResp
	for _, d := range tmp {
		resp = append(resp, d)
	}

	httputil.WriteResponse(rCtx, http.StatusOK, resp)
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

	product.ShopUserId = userId
	id, err := h.uProduct.CreateProduct(ctx, &product)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusCreated, &model.Product{Id: id})
}
