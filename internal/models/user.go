package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `gorm:"not null" json:"name"`
	Email        string         `gorm:"unique;not null" json:"email"`
	Age          int            `gorm:"not null" json:"age"`
	PasswordHash string         `gorm:"not null" json:"-"`
}
