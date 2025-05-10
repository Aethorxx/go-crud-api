# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Установка необходимых зависимостей
RUN apk add --no-cache gcc musl-dev

# Копируем файлы зависимостей
COPY go.mod ./
COPY go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Копируем бинарный файл из builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"] 