package models

import (
	"time"

	"gorm.io/gorm"
)

// Order представляет модель заказа в системе
// Связан с пользователем через UserID
type Order struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	Product   string         `json:"product" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`             // Мягкое удаление
	User      *User          `json:"-" gorm:"foreignKey:UserID"` // Скрываем данные пользователя из JSON
}

// CreateOrderRequest представляет данные для создания нового заказа
// Используется при создании заказа
type CreateOrderRequest struct {
	Product  string  `json:"product" binding:"required"`
	Quantity int     `json:"quantity" binding:"required,gt=0"`
	Price    float64 `json:"price" binding:"required,gt=0"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	return nil
}

func (o *Order) BeforeUpdate(tx *gorm.DB) error {
	o.UpdatedAt = time.Now()
	return nil
}
