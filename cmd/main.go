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
	orderRepo := repository.NewOrderRepository(db)
	userService := services.NewUserService(userRepo)
	orderService := services.NewOrderService(orderRepo, userRepo)
	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Настройка маршрутизатора
	router := gin.Default()

	// Public routes
	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/register", authHandler.Register)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		users := protected.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)

			// Order routes
			orders := users.Group("/:id/orders")
			{
				orders.GET("", orderHandler.GetUserOrders)
				orders.POST("", orderHandler.CreateOrder)
				orders.GET("/:order_id", orderHandler.GetOrder)
				orders.PUT("/:order_id", orderHandler.UpdateOrder)
				orders.DELETE("/:order_id", orderHandler.DeleteOrder)
			}
		}
	}

	// Запуск сервера
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
