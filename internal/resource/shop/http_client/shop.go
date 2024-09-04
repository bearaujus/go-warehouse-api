package http_client

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *shopResourceHTTPClientImpl) GetShopsByUser(ctx context.Context, userId uint64) ([]*model.Shop, error) {
	return nil, model.ErrRShopHTTPClientGetShopsByUser.New(model.ErrCommonNotImplemented)
}

func (r *shopResourceHTTPClientImpl) CreateShop(ctx context.Context, shop *model.Shop) (uint64, error) {
	return 0, model.ErrRShopHTTPClientCreateShop.New(model.ErrCommonNotImplemented)
}
