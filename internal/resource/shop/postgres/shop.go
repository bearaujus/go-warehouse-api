package postgres

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *shopResourcePostgresImpl) CreateShop(ctx context.Context, shop *model.Shop) error {
	err := r.db.WithContext(ctx).Create(shop).Error
	if err != nil {
		return model.ErrRShopPostgresCreateShop.New(err)
	}
	return nil
}
