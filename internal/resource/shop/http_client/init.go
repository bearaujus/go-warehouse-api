package http_client

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/shop"
	"net/http"
	"time"
)

type shopResourceHTTPClientImpl struct {
	host       string
	httpClient *http.Client

	serviceName    string
	serviceAuthKey string
	serviceAuthTTL time.Duration
}

func NewShopResourceHTTPClient(host string, timeout time.Duration, serviceName, serviceAuthKey string, serviceAuthTTL time.Duration) shop.ShopResource {
	return &shopResourceHTTPClientImpl{
		host:           host,
		httpClient:     &http.Client{Timeout: timeout},
		serviceName:    serviceName,
		serviceAuthKey: serviceAuthKey,
		serviceAuthTTL: serviceAuthTTL,
	}
}
