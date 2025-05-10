package repository

import (
	"go-crud-api/internal/models"

	"gorm.io/gorm"
)

// UserRepository отвечает за работу с данными пользователей в БД
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create сохраняет нового пользователя в БД
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByID находит пользователя по ID
// Возвращает ошибку если пользователь не найден
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// GetByEmail находит пользователя по email
// Используется при аутентификации
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Update обновляет данные пользователя
// Обновляет только непустые поля
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete удаляет пользователя по ID
// Использует мягкое удаление (soft delete)
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// List возвращает список пользователей с пагинацией
// Поддерживает сортировку и фильтрацию
func (r *UserRepository) List(page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	offset := (page - 1) * limit

	err := r.db.Model(&models.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(limit).Find(&users).Error
	return users, total, err
}

// CheckExists проверяет существование пользователя по email
// Используется при регистрации для проверки уникальности email
func (r *UserRepository) CheckExists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) List(params models.PaginationParams) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.db.Model(&models.User{})

	if params.MinAge > 0 {
		query = query.Where("age >= ?", params.MinAge)
	}
	if params.MaxAge > 0 {
		query = query.Where("age <= ?", params.MaxAge)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (params.Page - 1) * params.Limit
	err = query.Offset(offset).Limit(params.Limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *UserRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *UserRepository) GetUserOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
