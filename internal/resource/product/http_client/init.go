package http_client

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/product"
	"net/http"
	"time"
)

type productResourceHTTPClientImpl struct {
	host       string
	httpClient *http.Client
}

func NewProductResourceHTTPClient(host string, timeout time.Duration) product.ProductResource {
	return &productResourceHTTPClientImpl{host: host, httpClient: &http.Client{Timeout: timeout}}
}
