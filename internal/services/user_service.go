package services

import (
	"errors"
	"go-crud-api/internal/models"
	"go-crud-api/internal/repository"
	"go-crud-api/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

// UserService содержит бизнес-логику для работы с пользователями
// Включает валидацию данных и обработку ошибок
type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(user *models.User) error {
	// Проверяем, существует ли пользователь с таким email
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user with this email already exists")
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	// Создаем пользователя
	return s.userRepo.Create(user)
}

// Login аутентифицирует пользователя
// Проверяет email и пароль, возвращает пользователя
func (s *UserService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) Update(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) Delete(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) List() ([]models.User, error) {
	return s.userRepo.List()
}

// CreateUser создает нового пользователя
// Проверяет уникальность email и хеширует пароль
func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.UserResponse, error) {
	exists, err := s.userRepo.CheckExists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user with this email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		Age:          req.Age,
		PasswordHash: hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}, nil
}

// GetUser получает информацию о пользователе по ID
// Возвращает ошибку если пользователь не найден
func (s *UserService) GetUser(id uint) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}, nil
}

// UpdateUser обновляет данные пользователя
// Проверяет существование пользователя и валидирует данные
func (s *UserService) UpdateUser(id uint, req models.UpdateUserRequest) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		exists, err := s.userRepo.CheckExists(req.Email)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("user with this email already exists")
		}
		user.Email = req.Email
	}
	if req.Age != 0 {
		user.Age = req.Age
	}
	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = hashedPassword
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}, nil
}

// DeleteUser удаляет пользователя по ID
// Возвращает ошибку если пользователь не найден
func (s *UserService) DeleteUser(id uint) error {
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(id)
}

// ListUsers возвращает список пользователей с пагинацией
// Поддерживает фильтрацию и сортировку
func (s *UserService) ListUsers(params models.PaginationParams) (*models.PaginatedResponse, error) {
	users, total, err := s.userRepo.List(params.Page, params.Limit)
	if err != nil {
		return nil, err
	}

	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Age:       user.Age,
			CreatedAt: user.CreatedAt,
		}
	}

	totalPages := (int(total) + params.Limit - 1) / params.Limit

	return &models.PaginatedResponse{
		Data:       userResponses,
		Total:      total,
		Page:       params.Page,
		Limit:      params.Limit,
		TotalPages: totalPages,
	}, nil
}

func (s *UserService) CreateOrder(userID uint, req models.CreateOrderRequest) (*models.Order, error) {
	// Проверяем существование пользователя
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	order := &models.Order{
		UserID:   userID,
		Product:  req.Product,
		Quantity: req.Quantity,
		Price:    req.Price,
	}

	if err := s.userRepo.CreateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *UserService) GetUserOrders(userID uint) ([]models.Order, error) {
	// Проверяем существование пользователя
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return s.userRepo.GetUserOrders(userID)
}
