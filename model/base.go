package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID uint `gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}