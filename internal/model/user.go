package model

import (
	"gorm.io/plugin/optimisticlock"
	"time"
)

type User struct {
	Id           uint64                 `json:"id,omitempty"`
	Email        string                 `json:"email,omitempty"`
	Phone        string                 `json:"phone,omitempty"`
	PasswordHash string                 `json:"-"`
	Role         UserRole               `json:"role,omitempty"`
	CreatedAt    *time.Time             `json:"created_at,omitempty"`
	Version      optimisticlock.Version `json:"-"`
}

type UserRole string

const (
	UserRoleSeller UserRole = "seller"
	UserRoleBuyer  UserRole = "buyer"
)
