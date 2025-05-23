package models

import (
	"time"

	"gorm.io/gorm"
)

// User представляет модель пользователя в системе
// Содержит основную информацию о пользователе и его учетных данных
type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"not null"`
	Email        string         `json:"email" gorm:"unique;not null"`
	Age          int            `json:"age" gorm:"not null"`
	PasswordHash string         `json:"-" gorm:"not null"` // Скрываем хеш пароля из JSON
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"` // Мягкое удаление
	Orders       []Order        `json:"orders,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

// UserResponse представляет данные пользователя для ответа API
// Не включает конфиденциальную информацию
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateUserRequest представляет данные для создания нового пользователя
// Используется при регистрации
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gte=0,lte=130"`
	Password string `json:"password" binding:"required,min=6"`
}

// UpdateUserRequest представляет данные для обновления пользователя
// Все поля опциональны
type UpdateUserRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Age      int    `json:"age,omitempty" binding:"omitempty,gte=0,lte=130"`
	Password string `json:"password,omitempty" binding:"omitempty,min=6"`
}
