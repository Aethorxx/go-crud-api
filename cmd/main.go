package main

import (
	"log"
	"net/http"

	"go-crud-api/internal/config"
	"go-crud-api/internal/handlers"
	"go-crud-api/internal/repository"
	"go-crud-api/internal/router"
	"go-crud-api/internal/services"
	"go-crud-api/internal/utils"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Подключаемся к базе данных
	db, err := utils.InitDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Инициализируем репозитории
	userRepo := repository.NewUserRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	// Инициализируем сервисы
	userService := services.NewUserService(userRepo)
	orderService := services.NewOrderService(orderRepo, userRepo)

	// Инициализируем обработчики
	userHandler := handlers.NewUserHandler(userService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Настраиваем маршруты
	r := router.SetupRouter(userHandler, orderHandler)

	// Запускаем сервер
	log.Printf("Сервер запущен на порту %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
