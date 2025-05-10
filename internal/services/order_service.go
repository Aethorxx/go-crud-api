package services

import (
	"errors"

	"go-crud-api/internal/models"
	"go-crud-api/internal/repository"
)

// OrderService содержит бизнес-логику для работы с заказами
// Включает валидацию данных и проверку прав доступа
type OrderService struct {
	orderRepo *repository.OrderRepository
	userRepo  *repository.UserRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, userRepo *repository.UserRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

// Create создает новый заказ
// Проверяет существование пользователя и валидирует данные
func (s *OrderService) Create(order *models.Order) error {
	// Проверяем существование пользователя
	_, err := s.userRepo.GetByID(order.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	return s.orderRepo.Create(order)
}

// GetByID получает информацию о заказе по ID
// Возвращает ошибку если заказ не найден
func (s *OrderService) GetByID(id uint) (*models.Order, error) {
	return s.orderRepo.GetByID(id)
}

// GetByUserID получает все заказы пользователя
// Возвращает пустой слайс если заказов нет
func (s *OrderService) GetByUserID(userID uint) ([]models.Order, error) {
	// Проверяем существование пользователя
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return s.orderRepo.GetByUserID(userID)
}

// Update обновляет данные заказа
// Проверяет существование заказа и права доступа
func (s *OrderService) Update(order *models.Order) error {
	existingOrder, err := s.orderRepo.GetByID(order.ID)
	if err != nil {
		return err
	}

	// Проверяем, что заказ принадлежит пользователю
	if existingOrder.UserID != order.UserID {
		return errors.New("order does not belong to user")
	}

	return s.orderRepo.Update(order)
}

// Delete удаляет заказ по ID
// Проверяет существование заказа и права доступа
func (s *OrderService) Delete(id uint) error {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.orderRepo.Delete(order.ID)
}

// List возвращает список всех заказов с пагинацией
// Поддерживает фильтрацию и сортировку
func (s *OrderService) List(page, limit int) ([]models.Order, int64, error) {
	return s.orderRepo.List(page, limit)
}
