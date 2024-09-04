package postgres

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/user"
	"gorm.io/gorm"
)

type userResourcePostgresImpl struct {
	db *gorm.DB
}

func NewUserResourcePostgres(db *gorm.DB) user.UserResource {
	return &userResourcePostgresImpl{db: db}
}
