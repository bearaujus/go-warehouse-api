package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func (mw *MiddlewareAuth) AuthorizeSeller() app.HandlerFunc {
	return func(ctx context.Context, rCtx *app.RequestContext) {
		userId, err := GetUserIdFromContext(rCtx)
		if err != nil {
			httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrMWAuthAuthorizeSeller.New())
			return
		}

		usr, err := mw.rUser.GetUserById(ctx, userId)
		if err != nil {
			httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, err)
			return
		}

		if usr.Role != model.UserRoleSeller {
			httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrMWAuthAuthorizeSeller.New())
			return
		}

		rCtx.Next(ctx)
	}
}

func (mw *MiddlewareAuth) AuthorizeBuyer() app.HandlerFunc {
	return func(ctx context.Context, rCtx *app.RequestContext) {
		userId, err := GetUserIdFromContext(rCtx)
		if err != nil {
			httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrMWAuthAuthorizeBuyer.New())
			return
		}

		usr, err := mw.rUser.GetUserById(ctx, userId)
		if err != nil {
			httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, err)
			return
		}

		if usr.Role != model.UserRoleBuyer {
			httputil.WriteErrorResponseAndAbort(rCtx, http.StatusUnauthorized, model.ErrMWAuthAuthorizeBuyer.New())
			return
		}

		rCtx.Next(ctx)
	}
}

// GetUserIdFromContext should be called only after the request has passed through the AuthenticateUser middleware.
// If called before, it will always return an error.
func GetUserIdFromContext(rCtx *app.RequestContext) (uint64, error) {
	rawUserId, exist := rCtx.Get(ctxKeyAuthUser)
	if !exist {
		return 0, errors.New("not authenticated")
	}
	return pkg.StringToUint64(fmt.Sprintf("%v", rawUserId))
}

// GetServiceNameFromContext should be called only after the request has passed through the AuthenticateService middleware.
// If called before, it will always return an error.
func GetServiceNameFromContext(rCtx *app.RequestContext) (uint64, error) {
	rawUserId, exist := rCtx.Get(ctxKeyAuthUser)
	if !exist {
		return 0, errors.New("not authenticated")
	}
	return pkg.StringToUint64(fmt.Sprintf("%v", rawUserId))
}
