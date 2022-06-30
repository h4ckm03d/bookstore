package models

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition
type CommonTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
