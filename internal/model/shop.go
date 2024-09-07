package model

import (
	"fmt"
	"gorm.io/plugin/optimisticlock"
)

type Shop struct {
	UserId      uint64                 `json:"user_id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Version     optimisticlock.Version `json:"-"`
}

func (m *Shop) Validate() error {
	// validate user id
	if m.UserId == 0 {
		return fmt.Errorf("shop user id is required")
	}

	// validate name
	if m.Name == "" {
		return fmt.Errorf("shop name is required")
	}

	// validate description
	if m.Description == "" {
		return fmt.Errorf("shop description is required")
	}

	if len(m.Description) < 20 {
		return fmt.Errorf("shop description must be at least 20 characters")
	}

	return nil
}
