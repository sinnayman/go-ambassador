package models

import (
	"time"

	"gorm.io/gorm"
)

type ModelWrite struct {
	gorm.Model
}

type ModelRead struct {
	ID        int            `json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
