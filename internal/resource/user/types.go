package user

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

type UserResource interface {
	GetUserById(ctx context.Context, id uint64) (*model.User, error)
	GetUserByLogin(ctx context.Context, login, passwordHash string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (uint64, error)
}
