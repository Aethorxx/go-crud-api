package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-crud-api/internal/config"
	"go-crud-api/internal/handlers"
	"go-crud-api/internal/middleware"
	"go-crud-api/internal/repository"
	"go-crud-api/internal/services"
	"go-crud-api/internal/utils"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Применение миграций
	if err := utils.RunMigrations(); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(cfg.DB.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Инициализация зависимостей
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)

	// Настройка маршрутизатора
	router := gin.Default()

	// Публичные маршруты
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	// Защищенные маршруты
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/users", userHandler.GetUsers)
		authorized.GET("/users/:id", userHandler.GetUser)
		authorized.PUT("/users/:id", userHandler.UpdateUser)
		authorized.DELETE("/users/:id", userHandler.DeleteUser)
	}

	// Запуск сервера
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
