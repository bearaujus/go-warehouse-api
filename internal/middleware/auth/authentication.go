package auth

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/authutil"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

const ctxKeyAuthUser string = "XXX-ctx-key-auth-user"
const ctxKeyAuthService string = "XXX-ctx-key-auth-service"

func (mw *MiddlewareAuth) AuthenticateUser() app.HandlerFunc {
	return func(ctx context.Context, rCtx *app.RequestContext) {
		data, err := authutil.ReadAuthTokenDataFromHTTPRequest(rCtx, mw.authUserSecretKey)
		userId, err := pkg.StringToUint64(data)
		if err != nil {
			httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrMWAuthAuthenticateUser.New())
			return
		}

		u, err := mw.rUser.GetUserById(ctx, userId)
		if err != nil {
			httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrMWAuthAuthenticateUser.New())
			return
		}

		rCtx.Set(ctxKeyAuthUser, u.Id)
		rCtx.Next(ctx)
	}
}

func (mw *MiddlewareAuth) AuthenticateService() app.HandlerFunc {
	return func(ctx context.Context, rCtx *app.RequestContext) {
		for serviceName, authServiceSecretKey := range mw.authAllowedServices {
			data, err := authutil.ReadAuthTokenDataFromHTTPRequest(rCtx, authServiceSecretKey)
			if err != nil {
				continue
			}

			if data != serviceName {
				break
			}

			rCtx.Next(ctx)
			return
		}

		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrMWAuthAuthenticateService.New())
	}
}
