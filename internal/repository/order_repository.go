package repository

import (
	"context"
	"errors"

	"go-crud-api/internal/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Create создает новый заказ
func (r *OrderRepository) Create(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

// GetByID получает заказ по ID
func (r *OrderRepository) GetByID(ctx context.Context, id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.WithContext(ctx).First(&order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}

// Update обновляет данные заказа
func (r *OrderRepository) Update(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}

// Delete удаляет заказ
func (r *OrderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Order{}, id).Error
}

// List получает список всех заказов
func (r *OrderRepository) List(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.WithContext(ctx).Find(&orders).Error
	return orders, err
}

// GetByUserID получает все заказы пользователя
func (r *OrderRepository) GetByUserID(ctx context.Context, userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
