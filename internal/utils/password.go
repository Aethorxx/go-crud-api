package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword создает хеш пароля с использованием bcrypt
// Использует стандартную стоимость хеширования (10 раундов)
// Возвращает хеш в виде строки
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash сравнивает пароль с его хешем
// Возвращает true если пароль совпадает с хешем
// Использует безопасное сравнение для предотвращения timing-атак
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
