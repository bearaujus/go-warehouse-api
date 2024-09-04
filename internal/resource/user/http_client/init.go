package http_client

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/user"
	"net/http"
	"time"
)

type userResourceHTTPClientImpl struct {
	host       string
	httpClient *http.Client

	serviceName    string
	serviceAuthKey string
	serviceAuthTTL time.Duration
}

func NewUserResourceHTTPClient(host string, timeout time.Duration, serviceName, serviceAuthKey string, serviceAuthTTL time.Duration) user.UserResource {
	return &userResourceHTTPClientImpl{
		host:           host,
		httpClient:     &http.Client{Timeout: timeout},
		serviceName:    serviceName,
		serviceAuthKey: serviceAuthKey,
		serviceAuthTTL: serviceAuthTTL,
	}
}
