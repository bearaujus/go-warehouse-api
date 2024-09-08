package user

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/resource/shop"
	"github.com/bearaujus/go-warehouse-api/internal/resource/user"
	"time"
)

type UserUsecase interface {
	GetUserById(ctx context.Context, id uint64) (*model.User, error)
	Register(ctx context.Context, user *model.User, passwordRaw string, shop *model.Shop) (uint64, error)
	Login(ctx context.Context, login, passwordRaw string) (string, model.UserRole, error)
}

type userUsecaseImpl struct {
	rUser user.UserResource
	rShop shop.ShopResource

	authSecretKey string
	authTTL       time.Duration
}

func NewUserUsecase(rUser user.UserResource, rShop shop.ShopResource, authSecretKey string, authTTL time.Duration) UserUsecase {
	return &userUsecaseImpl{rUser: rUser, rShop: rShop, authSecretKey: authSecretKey, authTTL: authTTL}
}
