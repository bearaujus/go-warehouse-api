package shop

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/resource/shop"
)

type ShopUsecase interface {
	CreateShop(ctx context.Context, shop *model.Shop) error
}

type shopUsecaseImpl struct {
	rShop shop.ShopResource
}

func NewShopUsecase(rShop shop.ShopResource) ShopUsecase {
	return &shopUsecaseImpl{rShop: rShop}
}
