package http

import (
	"context"
	"encoding/json"
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

func (h *userHandlerHTTPImpl) RegisterUser(ctx context.Context, rCtx *app.RequestContext) {
	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHUserHTTPRegisterUser.New(model.ErrCommonInvalidRequestBody))
		return
	}

	type registerUserReq struct {
		Email       string `json:"email"`
		Phone       string `json:"phone"`
		PasswordRaw string `json:"password"`
		ShopName    string `json:"shop_name"`
		ShopDesc    string `json:"shop_desc"`
	}

	registerUser := registerUserReq{}
	err = json.Unmarshal(data, &registerUser)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHUserHTTPRegisterUser.New(model.ErrCommonInvalidRequestBody))
		return
	}

	user := model.User{
		Email: registerUser.Email,
		Phone: registerUser.Phone,
	}

	_, err = h.uUser.Register(ctx, &user, registerUser.PasswordRaw, &model.Shop{
		Name:        registerUser.ShopName,
		Description: registerUser.ShopDesc,
	})
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	user.PasswordHash = ""
	httputil.WriteResponse(rCtx, http.StatusCreated, user)
}

func (h *userHandlerHTTPImpl) LoginUser(ctx context.Context, rCtx *app.RequestContext) {
	data, err := rCtx.Body()
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHUserHTTPLoginUser.New(model.ErrCommonInvalidRequestBody))
		return
	}

	type loginUserReq struct {
		Login       string `json:"login"`
		PasswordRaw string `json:"password"`
	}

	loginUser := loginUserReq{}
	err = json.Unmarshal(data, &loginUser)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, model.ErrHUserHTTPLoginUser.New(model.ErrCommonInvalidRequestBody))
		return
	}

	userAuthToken, userRole, err := h.uUser.Login(ctx, loginUser.Login, loginUser.PasswordRaw)
	if err != nil {
		httputil.WriteErrorResponseAndAbort(rCtx, http.StatusBadRequest, err)
		return
	}

	httputil.WriteResponse(rCtx, http.StatusOK, utils.H{
		"token": userAuthToken,
		"role":  userRole,
	})
}
