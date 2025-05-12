package utils

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	// Создаем файл для логов
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	// Инициализируем логгеры
	InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogOperation логирует операцию с пользователем
func LogOperation(operation string, userID uint, details string) {
	InfoLogger.Printf("[%s] UserID: %d, Details: %s", operation, userID, details)
}

// LogError логирует ошибку
func LogError(operation string, err error) {
	ErrorLogger.Printf("[%s] Error: %v", operation, err)
}

// LogOrderOperation логирует операцию с заказом
func LogOrderOperation(operation string, orderID uint, userID uint, details string) {
	InfoLogger.Printf("[%s] OrderID: %d, UserID: %d, Details: %s", operation, orderID, userID, details)
}
