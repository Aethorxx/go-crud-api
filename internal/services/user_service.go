package services

import (
	"context"
	"errors"

	"go-crud-api/internal/models"
	"go-crud-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Register регистрирует нового пользователя
func (s *UserService) Register(ctx context.Context, user *models.User) error {
	// Проверяем, не существует ли уже пользователь с таким email
	existingUser, err := s.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("пользователь с таким email уже существует")
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	// Создаем пользователя
	return s.repo.Create(ctx, user)
}

// Login выполняет вход пользователя
func (s *UserService) Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("пользователь не найден")
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("неверный пароль")
	}

	return user, nil
}

// GetByID получает пользователя по ID
func (s *UserService) GetByID(ctx context.Context, id uint) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

// Update обновляет данные пользователя
func (s *UserService) Update(ctx context.Context, user *models.User) error {
	// Если пароль был изменен, хешируем его
	if user.PasswordHash != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.PasswordHash = string(hashedPassword)
	}
	return s.repo.Update(ctx, user)
}

// Delete удаляет пользователя
func (s *UserService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// List получает список всех пользователей
func (s *UserService) List(ctx context.Context) ([]models.User, error) {
	return s.repo.List(ctx)
}
