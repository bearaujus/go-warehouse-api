package user

import (
	"context"
	"fmt"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/authutil"
)

func (u *userUsecaseImpl) GetUserById(ctx context.Context, id uint64) (*model.User, error) {
	return u.rUser.GetUserById(ctx, id)
}

func (u *userUsecaseImpl) Register(ctx context.Context, user *model.User, passwordRaw string, shop *model.Shop) (uint64, error) {
	user.Role = model.UserRoleBuyer
	if shop.Name != "" || shop.Description != "" {
		user.Role = model.UserRoleSeller
	}

	err := user.Validate(passwordRaw)
	if err != nil {
		return 0, model.ErrUUserRegister.New(err)
	}

	user.PasswordHash, err = pkg.ToSHA256Hash(passwordRaw)
	if err != nil {
		return 0, model.ErrUUserRegister.New(err)
	}

	user.Id = 0
	user.CreatedAt = nil
	user.Id, err = u.rUser.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	if user.Role == model.UserRoleSeller {
		shop.UserId = user.Id
		err = u.rShop.CreateShop(ctx, shop)
		if err != nil {
			_ = u.rUser.DeleteUser(ctx, user.Id)
			return 0, err
		}
	}

	return user.Id, nil
}

func (u *userUsecaseImpl) Login(ctx context.Context, login, passwordRaw string) (string, model.UserRole, error) {
	passwordHash, err := pkg.ToSHA256Hash(passwordRaw)
	if err != nil {
		return "", "", model.ErrUUserLogin.New(err)
	}

	user, err := u.rUser.GetUserByLogin(ctx, login, passwordHash)
	if err != nil {
		return "", "", err
	}

	token, err := authutil.GenerateAuthToken(fmt.Sprint(user.Id), u.authSecretKey, u.authTTL)
	if err != nil {
		return "", "", model.ErrUUserLogin.New(err)
	}

	return token, user.Role, nil
}
