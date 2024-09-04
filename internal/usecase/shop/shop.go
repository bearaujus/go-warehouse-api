package shop

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (u *shopUsecaseImpl) GetShopsByUser(ctx context.Context, userId uint64) ([]*model.Shop, error) {
	return u.rShop.GetShopsByUser(ctx, userId)
}

func (u *shopUsecaseImpl) CreateShop(ctx context.Context, shop *model.Shop) (uint64, error) {
	// validate name
	if shop.Name == "" {
		return 0, model.ErrUShopCreateShop.New("name is required")
	}

	// validate description
	if shop.Description == "" {
		return 0, model.ErrUShopCreateShop.New("description is required")
	}

	if len(shop.Description) < 20 {
		return 0, model.ErrUShopCreateShop.New("description must be at least 20 characters")
	}

	shop.Id = 0
	return u.rShop.CreateShop(ctx, shop)
}
