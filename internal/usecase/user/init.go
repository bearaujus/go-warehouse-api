package user

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/resource/user"
	"time"
)

type UserUsecase interface {
	GetUserById(ctx context.Context, id uint64) (*model.User, error)
	Register(ctx context.Context, email, phone, roleRaw, passwordRaw string) (uint64, error)
	Login(ctx context.Context, login, passwordRaw string) (string, error)
}

type userUsecaseImpl struct {
	rUser user.UserResource

	authSecretKey string
	authTTL       time.Duration
}

func NewUserUsecase(rUser user.UserResource, authSecretKey string, authTTL time.Duration) UserUsecase {
	return &userUsecaseImpl{rUser: rUser, authSecretKey: authSecretKey, authTTL: authTTL}
}
