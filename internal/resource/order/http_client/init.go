package http_client

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/order"
	"net/http"
	"time"
)

type orderResourceHTTPClientImpl struct {
	host       string
	httpClient *http.Client
}

func NewOrderResourceHTTPClient(host string, timeout time.Duration) order.OrderResource {
	return &orderResourceHTTPClientImpl{host: host, httpClient: &http.Client{Timeout: timeout}}
}
