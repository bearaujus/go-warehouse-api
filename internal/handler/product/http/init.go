package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/handler"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/product"
)

type productHandlerHTTPImpl struct {
	uProduct product.ProductUsecase
}

func NewProductHandlerHTTP(uProduct product.ProductUsecase) handler.HTTPHandler {
	return &productHandlerHTTPImpl{uProduct: uProduct}
}
