package http_client

import (
	"context"
	"fmt"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/authutil"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"net/http"
)

func (r *userResourceHTTPClientImpl) GetUserById(ctx context.Context, id uint64) (*model.User, error) {
	url := fmt.Sprintf("%v/internal/users/%v", r.host, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, model.ErrRUserHTTPClientGetUserById.New(err)
	}

	err = authutil.GenerateAndSetAuthTokenForHTTPRequestHeader(req, r.serviceName, r.serviceAuthKey, r.serviceAuthTTL)
	if err != nil {
		return nil, model.ErrRUserHTTPClientGetUserById.New(err)
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, model.ErrRUserHTTPClientGetUserById.New(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, model.ErrRUserHTTPClientGetUserById.New(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}

	var user model.User
	err = httputil.DecodeUnmarshalResponseBody(resp.Body, &user)
	if err != nil {
		return nil, model.ErrRUserHTTPClientGetUserById.New(err)
	}

	return &user, nil
}

func (r *userResourceHTTPClientImpl) GetUserByLogin(ctx context.Context, login, passwordHash string) (*model.User, error) {
	return nil, model.ErrRUserHTTPClientGetUserByLogin.New(model.ErrCommonNotImplemented)
}

func (r *userResourceHTTPClientImpl) CreateUser(ctx context.Context, user *model.User) (uint64, error) {
	return 0, model.ErrRUserHTTPClientCreateUser.New(model.ErrCommonNotImplemented)
}
