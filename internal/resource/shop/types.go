package shop

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

type ShopResource interface {
	GetShopsByUser(ctx context.Context, userId uint64) ([]*model.Shop, error)
	CreateShop(ctx context.Context, shop *model.Shop) (uint64, error)
}
