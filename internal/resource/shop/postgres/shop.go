package postgres

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *shopResourcePostgresImpl) GetShopsByUser(ctx context.Context, userId uint64) ([]*model.Shop, error) {
	var shops []*model.Shop
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&shops).Error
	if err != nil {
		return nil, model.ErrRShopPostgresGetShopsByUser.New(err)
	}
	return shops, nil
}

func (r *shopResourcePostgresImpl) CreateShop(ctx context.Context, shop *model.Shop) (uint64, error) {
	err := r.db.WithContext(ctx).Create(shop).Error
	if err != nil {
		return 0, model.ErrRShopPostgresCreateShop.New(err)
	}
	return shop.Id, nil
}
