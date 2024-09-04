package http_client

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *productResourceHTTPClientImpl) GetProductsWithStockByUser(ctx context.Context, userId uint64) ([]*model.ProductWithStock, error) {
	return nil, model.ErrRProductHTTPClientGetProductsWithStockByUser.New(model.ErrCommonNotImplemented)
}

func (r *productResourceHTTPClientImpl) CreateProduct(ctx context.Context, product *model.Product) (uint64, error) {
	return 0, model.ErrRProductHTTPClientCreateProduct.New(model.ErrCommonNotImplemented)
}
