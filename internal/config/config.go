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
// Возвращает ошибку если какие-то переменные не заданы
func Load() (*Config, error) {
	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Загружаем настройки базы данных
	dbPort, err := getEnvAsInt("DB_PORT")
	if err != nil {
		return nil, fmt.Errorf("error loading DB_PORT: %v", err)
	}

	dbConfig := DBConfig{
		Host:     getEnv("DB_HOST"),
		Port:     dbPort,
		User:     getEnv("DB_USER"),
		Password: getEnv("DB_PASSWORD"),
		DBName:   getEnv("DB_NAME"),
		SSLMode:  getEnv("DB_SSL_MODE"),
	}

	// Загружаем настройки сервера
	serverPort, err := getEnvAsInt("SERVER_PORT")
	if err != nil {
		return nil, fmt.Errorf("error loading SERVER_PORT: %v", err)
	}

	serverConfig := ServerConfig{
		Port: serverPort,
	}

	// Загружаем настройки JWT
	jwtConfig := JWTConfig{
		Secret: getEnv("JWT_SECRET"),
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
// Возвращает ошибку если переменная не задана
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s is not set", key))
	}
	return value
}

// getEnvAsInt получает целочисленное значение переменной окружения
// Возвращает ошибку если переменная не задана или не является числом
func getEnvAsInt(key string) (int, error) {
	valueStr := getEnv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("environment variable %s is not a valid integer: %v", key, err)
	}
	return value, nil
}
