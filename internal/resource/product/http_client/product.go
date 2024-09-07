package http_client

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *productResourceHTTPClientImpl) GetProductsByShopUserIdAndWarehouseStatus(ctx context.Context, shopUserId uint64, status model.WarehouseStatus) ([]*model.Product, error) {
	return nil, model.ErrRProductHTTPClientGetProductsByShopUserIdAndWarehouseStatus.New(model.ErrCommonNotImplemented)
}

func (r *productResourceHTTPClientImpl) CreateProduct(ctx context.Context, product *model.Product) (uint64, error) {
	return 0, model.ErrRProductHTTPClientCreateProduct.New(model.ErrCommonNotImplemented)
}
