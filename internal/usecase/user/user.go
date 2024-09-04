package user

import (
	"context"
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/authutil"
	"github.com/nyaruka/phonenumbers"
)

func (u *userUsecaseImpl) GetUserById(ctx context.Context, id uint64) (*model.User, error) {
	return u.rUser.GetUserById(ctx, id)
}

func (u *userUsecaseImpl) Register(ctx context.Context, email, phone, roleRaw, passwordRaw string) (uint64, error) {
	// validate email
	if email == "" {
		return 0, model.ErrUUserRegister.New("email is required")
	}

	err := checkmail.ValidateFormat(email)
	if err != nil {
		return 0, model.ErrUUserRegister.New("email is invalid")
	}

	// validate phone
	if phone == "" {
		return 0, model.ErrUUserRegister.New("phone is required")
	}

	_, err = phonenumbers.Parse(phone, "ID")
	if err != nil {
		return 0, model.ErrUUserRegister.New("phone is invalid")
	}

	// validate password
	if passwordRaw == "" {
		return 0, model.ErrUUserRegister.New("password is required")
	}

	if len(passwordRaw) < 6 {
		return 0, model.ErrUUserRegister.New("password must be at least 6 characters")
	}

	// validate role
	if roleRaw != string(model.UserRoleSeller) && roleRaw != string(model.UserRoleBuyer) {
		return 0, model.ErrUUserRegister.New("role is invalid")
	}

	// hash the raw password
	passwordHash, err := pkg.ToSHA256Hash(passwordRaw)
	if err != nil {
		return 0, model.ErrUUserRegister.New(err)
	}

	return u.rUser.CreateUser(ctx, &model.User{
		Email:        email,
		Phone:        phone,
		PasswordHash: passwordHash,
		Role:         model.UserRole(roleRaw),
	})
}

func (u *userUsecaseImpl) Login(ctx context.Context, login, passwordRaw string) (string, error) {
	passwordHash, err := pkg.ToSHA256Hash(passwordRaw)
	if err != nil {
		return "", model.ErrUUserLogin.New(err)
	}

	user, err := u.rUser.GetUserByLogin(ctx, login, passwordHash)
	if err != nil {
		return "", err
	}

	token, err := authutil.GenerateAuthToken(fmt.Sprint(user.Id), u.authSecretKey, u.authTTL)
	if err != nil {
		return "", model.ErrUUserLogin.New(err)
	}

	return token, nil
}
