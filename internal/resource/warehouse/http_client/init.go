package http_client

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/warehouse"
	"net/http"
	"time"
)

type warehouseResourceHTTPClientImpl struct {
	host       string
	httpClient *http.Client
}

func NewWarehouseResourceHTTPClient(host string, timeout time.Duration) warehouse.WarehouseResource {
	return &warehouseResourceHTTPClientImpl{host: host, httpClient: &http.Client{Timeout: timeout}}
}
