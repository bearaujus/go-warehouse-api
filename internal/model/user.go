package model

import (
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/nyaruka/phonenumbers"
	"gorm.io/plugin/optimisticlock"
	"time"
)

type UserRole string

const (
	UserRoleSeller UserRole = "seller"
	UserRoleBuyer  UserRole = "buyer"
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

func (m *User) Validate(passwordRaw string) error {
	// validate email
	if m.Email == "" {
		return fmt.Errorf("user email is required")
	}

	err := checkmail.ValidateFormat(m.Email)
	if err != nil {
		return fmt.Errorf("user email is invalid")
	}

	// validate phone
	if m.Phone == "" {
		return fmt.Errorf("user phone is required")
	}

	_, err = phonenumbers.Parse(m.Phone, "ID")
	if err != nil {
		return fmt.Errorf("user phone is invalid")
	}

	// validate password
	if passwordRaw == "" {
		return fmt.Errorf("user password is required")
	}

	if len(passwordRaw) < 6 {
		return fmt.Errorf("user password must be at least 6 characters")
	}

	// validate role
	if m.Role != "" && m.Role != UserRoleSeller && m.Role != UserRoleBuyer {
		return fmt.Errorf("invalid user role. user role must be one of [%v, %v]", UserRoleSeller, UserRoleBuyer)
	}

	return nil
}
