package shop

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (u *shopUsecaseImpl) CreateShop(ctx context.Context, shop *model.Shop) error {
	err := shop.Validate()
	if err != nil {
		return model.ErrUShopCreateShop.New(err)
	}

	return u.rShop.CreateShop(ctx, shop)
}
