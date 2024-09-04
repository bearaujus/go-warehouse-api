package postgres

import (
	"context"
	"github.com/bearaujus/go-warehouse-api/internal/model"
)

func (r *userResourcePostgresImpl) GetUserById(ctx context.Context, id uint64) (*model.User, error) {
	var ret model.User
	err := r.db.WithContext(ctx).Model(&ret).Where("id = ?", id).First(&ret).Error
	if err != nil {
		return nil, model.ErrRUserPostgresGetUserById.New(err)
	}
	return &ret, nil
}

func (r *userResourcePostgresImpl) GetUserByLogin(ctx context.Context, login, passwordHash string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Model(&user).Where("(email = ? OR phone = ?) AND password_hash = ?", login, login, passwordHash).First(&user).Error
	if err != nil {
		return nil, model.ErrRUserPostgresGetUserByLogin.New(err)
	}
	return &user, nil
}

func (r *userResourcePostgresImpl) CreateUser(ctx context.Context, user *model.User) (uint64, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return 0, model.ErrRUserPostgresCreateUser.New(err)
	}
	return user.Id, nil
}
