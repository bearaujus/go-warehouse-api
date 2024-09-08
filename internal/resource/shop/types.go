package shop

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

type ShopResource interface {
	CreateShop(ctx context.Context, shop *model.Shop) error
}
