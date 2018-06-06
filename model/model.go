package model

import (
	"time"
)

type Model struct {
	ID        uint64     `json:"id" gorm:"primary_key" json:"id,omitempty`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
