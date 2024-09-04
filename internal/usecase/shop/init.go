package shop

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/resource/shop"
)

type ShopUsecase interface {
	GetShopsByUser(ctx context.Context, userId uint64) ([]*model.Shop, error)
	CreateShop(ctx context.Context, shop *model.Shop) (uint64, error)
}

type shopUsecaseImpl struct {
	rShop shop.ShopResource
}

func NewShopUsecase(rShop shop.ShopResource) ShopUsecase {
	return &shopUsecaseImpl{rShop: rShop}
}
