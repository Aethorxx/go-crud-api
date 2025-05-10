package repository

import (
	"go-crud-api/internal/models"

	"gorm.io/gorm"
)

// OrderRepository отвечает за работу с данными заказов в БД
type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Create сохраняет новый заказ в БД
func (r *OrderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

// GetByID находит заказ по ID
// Возвращает ошибку если заказ не найден
func (r *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, id).Error
	return &order, err
}

// GetByUserID находит все заказы пользователя
// Возвращает пустой слайс если заказов нет
func (r *OrderRepository) GetByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

// Update обновляет данные заказа
// Обновляет только непустые поля
func (r *OrderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

// Delete удаляет заказ по ID
// Использует мягкое удаление (soft delete)
func (r *OrderRepository) Delete(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}

// List возвращает список всех заказов с пагинацией
// Поддерживает сортировку и фильтрацию
func (r *OrderRepository) List(page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * limit

	err := r.db.Model(&models.Order{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(limit).Find(&orders).Error
	return orders, total, err
}
