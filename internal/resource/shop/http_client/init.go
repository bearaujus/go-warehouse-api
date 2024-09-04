package http_client

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/shop"
	"net/http"
	"time"
)

type shopResourceHTTPClientImpl struct {
	host       string
	httpClient *http.Client
}

func NewShopResourceHTTPClient(host string, timeout time.Duration) shop.ShopResource {
	return &shopResourceHTTPClientImpl{host: host, httpClient: &http.Client{Timeout: timeout}}
}
