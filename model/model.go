package model

import (
	"time"
)

type Model struct {
	ID        uint64     `json:"id" gorm:"primary_key" json:"id,omitempty`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
