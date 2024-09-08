package http

import (
	"github.com/bearaujus/go-warehouse-api/internal/handler"
	"github.com/bearaujus/go-warehouse-api/internal/usecase/shop"
)

type shopHandlerHTTPImpl struct {
	uShop shop.ShopUsecase
}

func NewShopHandlerHTTP(uShop shop.ShopUsecase) handler.HTTPHandler {
	return &shopHandlerHTTPImpl{uShop: uShop}
}
