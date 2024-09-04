package model

import (
	"gorm.io/plugin/optimisticlock"
	"time"
)

type Shop struct {
	Id          uint64                 `gorm:"primaryKey" json:"id,omitempty"`
	UserId      uint64                 `gorm:"index" json:"user_id,omitempty"`
	Name        string                 `gorm:"size:255;not null" json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	CreatedAt   *time.Time             `gorm:"autoCreateTime" json:"created_at,omitempty"`
	Warehouses  []*Warehouse           `json:"warehouses,omitempty"`
	Version     optimisticlock.Version `json:"-"`
}
