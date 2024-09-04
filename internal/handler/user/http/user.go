package http

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/httputil"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
)

func (h *userHandlerHTTPImpl) GetUserById(ctx context.Context, rCtx *app.RequestContext) {
	id, err := pkg.StringToUint64(rCtx.Param("id"))
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHUserHTTPGetUserById.New(model.ErrCommonInvalidRequestURL))
		return
	}

	user, err := h.uUser.GetUserById(ctx, id)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusOK, user)
}

func (h *userHandlerHTTPImpl) Register(ctx context.Context, rCtx *app.RequestContext) {
	userEmail, _ := rCtx.GetPostForm("email")
	userPhone, _ := rCtx.GetPostForm("phone")
	userRawRole, _ := rCtx.GetPostForm("role")
	userRawPassword, _ := rCtx.GetPostForm("password")

	userId, err := h.uUser.Register(ctx, userEmail, userPhone, userRawRole, userRawPassword)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusCreated, &model.User{Id: userId})
}

func (h *userHandlerHTTPImpl) Login(ctx context.Context, rCtx *app.RequestContext) {
	userLogin, _ := rCtx.GetPostForm("login")
	userRawPassword, _ := rCtx.GetPostForm("password")

	userAuthToken, err := h.uUser.Login(ctx, userLogin, userRawPassword)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusOK, utils.H{"token": userAuthToken})
}
