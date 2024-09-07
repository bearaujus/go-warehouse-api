package postgres

import (
	"context"
	"errors"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"strings"
)

func (r *userResourcePostgresImpl) GetUserById(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, model.ErrRUserPostgresGetUserById.New(err)
	}
	return &user, nil
}

func (r *userResourcePostgresImpl) GetUserByLogin(ctx context.Context, login, passwordHash string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("(email = ? OR phone = ?) AND password_hash = ?", login, login, passwordHash).First(&user).Error
	if err != nil {
		return nil, model.ErrRUserPostgresGetUserByLogin.New("invalid login or password")
	}
	return &user, nil
}

func (r *userResourcePostgresImpl) CreateUser(ctx context.Context, user *model.User) (uint64, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			if strings.Contains(strings.ToLower(err.Error()), "users_email_key") {
				return 0, model.ErrRUserPostgresCreateUser.New("user email is already registered")
			}
			if strings.Contains(strings.ToLower(err.Error()), "users_phone_key") {
				return 0, model.ErrRUserPostgresCreateUser.New("user phone is already registered")
			}
		}
		return 0, model.ErrRUserPostgresCreateUser.New(err)
	}
	return user.Id, nil
}

func (r *userResourcePostgresImpl) DeleteUser(ctx context.Context, id uint64) error {
	q := r.db.WithContext(ctx).Delete(&model.User{Id: id})
	err := q.Error
	if err != nil {
		return model.ErrRUserPostgresDeleteUser.New(err)
	}
	if q.RowsAffected == 0 {
		return model.ErrRUserPostgresDeleteUser.New(gorm.ErrRecordNotFound)
	}
	return nil
}
