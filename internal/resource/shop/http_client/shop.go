package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/authutil"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"net/http"
)

func (r *shopResourceHTTPClientImpl) CreateShop(ctx context.Context, shop *model.Shop) error {
	reqBody, err := json.Marshal(shop)
	if err != nil {
		return model.ErrRShopHTTPClientCreateShop.New(err)
	}

	url := fmt.Sprintf("%v/internal/shops", r.host)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return model.ErrRShopHTTPClientCreateShop.New(err)
	}

	err = authutil.GenerateAndSetAuthTokenForHTTPRequestHeader(req, r.serviceName, r.serviceAuthKey, r.serviceAuthTTL)
	if err != nil {
		return model.ErrRShopHTTPClientCreateShop.New(err)
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return model.ErrRShopHTTPClientCreateShop.New(err)
	}
	defer resp.Body.Close()

	err = httputil.CheckResponseBodyError(resp.Body)
	if err != nil {
		return model.ErrRShopHTTPClientCreateShop.New(err)
	}

	if resp.StatusCode != http.StatusCreated {
		return model.ErrRShopHTTPClientCreateShop.New(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}

	return nil
}
