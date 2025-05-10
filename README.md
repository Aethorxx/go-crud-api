# Go CRUD API

REST API на Go для управления пользователями и их заказами.

## Требования

- Go 1.21 или выше
- PostgreSQL 15 или выше
- Make (опционально)

## Установка и запуск

1. Клонируйте репозиторий:
2. Установите зависимости:
```bash
go mod download
```

3. Создайте базу данных PostgreSQL:
```sql
CREATE DATABASE go_crud_api;
```

4. Создайте файл .env в корне проекта:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_crud_api
DB_SSL_MODE=disable

JWT_SECRET=secret-key
SERVER_PORT=8080
```

5. Запустите приложение:
```bash
go run cmd/api/main.go
```

## API Endpoints

### Публичные маршруты
- `POST /auth/login` - авторизация
- `POST /users` - создание пользователя

### Защищенные маршруты (требуют JWT токен)
- `GET /users` - получение списка пользователей
- `GET /users/:id` - получение пользователя по ID
- `PUT /users/:id` - обновление пользователя
- `DELETE /users/:id` - удаление пользователя
- `POST /users/:user_id/orders` - создание заказа
- `GET /users/:user_id/orders` - получение списка заказов пользователя

## Примеры запросов

### Создание пользователя
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30,
    "password": "securepassword"
  }'
```

### Авторизация
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

### Создание заказа (с токеном)
```bash
curl -X POST http://localhost:8080/users/1/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "product": "Laptop",
    "quantity": 1,
    "price": 1200.50
  }'
``` 