package services

import (
	"context"
	"errors"

	"go-crud-api/internal/models"
	"go-crud-api/internal/repository"
)

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
func (s *OrderService) Create(ctx context.Context, order *models.Order) error {
	// Проверяем существование пользователя
	user, err := s.userRepo.GetByID(ctx, order.UserID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("пользователь не найден")
	}

	return s.orderRepo.Create(ctx, order)
}

// GetByID получает заказ по ID
func (s *OrderService) GetByID(ctx context.Context, id uint) (*models.Order, error) {
	return s.orderRepo.GetByID(ctx, id)
}

// Update обновляет данные заказа
func (s *OrderService) Update(ctx context.Context, order *models.Order) error {
	// Проверяем существование заказа
	existingOrder, err := s.orderRepo.GetByID(ctx, order.ID)
	if err != nil {
		return err
	}
	if existingOrder == nil {
		return errors.New("заказ не найден")
	}

	// Проверяем существование пользователя, если ID пользователя изменился
	if existingOrder.UserID != order.UserID {
		user, err := s.userRepo.GetByID(ctx, order.UserID)
		if err != nil {
			return err
		}
		if user == nil {
			return errors.New("пользователь не найден")
		}
	}

	return s.orderRepo.Update(ctx, order)
}

// Delete удаляет заказ
func (s *OrderService) Delete(ctx context.Context, id uint) error {
	return s.orderRepo.Delete(ctx, id)
}

// List получает список всех заказов
func (s *OrderService) List(ctx context.Context) ([]models.Order, error) {
	return s.orderRepo.List(ctx)
}

// GetByUserID получает все заказы пользователя
func (s *OrderService) GetByUserID(ctx context.Context, userID uint) ([]models.Order, error) {
	// Проверяем существование пользователя
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("пользователь не найден")
	}

	return s.orderRepo.GetByUserID(ctx, userID)
}
