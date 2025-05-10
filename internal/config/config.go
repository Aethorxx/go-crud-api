package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config содержит все настройки приложения
// Загружается из переменных окружения
type Config struct {
	DB     DBConfig
	Server ServerConfig
	JWT    JWTConfig
}

// DBConfig содержит настройки подключения к базе данных
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// ServerConfig содержит настройки HTTP сервера
type ServerConfig struct {
	Port int
}

// JWTConfig содержит настройки JWT токенов
type JWTConfig struct {
	Secret string
}

// Load загружает конфигурацию из .env файла
// Использует значения по умолчанию если переменные не заданы
func Load() (*Config, error) {
	// Загружаем .env файл если он существует
	_ = godotenv.Load()

	// Загружаем настройки базы данных
	dbConfig := DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvAsInt("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "go_crud_api"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}

	// Загружаем настройки сервера
	serverConfig := ServerConfig{
		Port: getEnvAsInt("SERVER_PORT", 8080),
	}

	// Загружаем настройки JWT
	jwtConfig := JWTConfig{
		Secret: getEnv("JWT_SECRET", "your-secret-key"),
	}

	return &Config{
		DB:     dbConfig,
		Server: serverConfig,
		JWT:    jwtConfig,
	}, nil
}

// GetDSN возвращает строку подключения к PostgreSQL
func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

// getEnv получает значение переменной окружения
// Возвращает значение по умолчанию если переменная не задана
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvAsInt получает целочисленное значение переменной окружения
// Возвращает значение по умолчанию если переменная не задана или не является числом
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}
