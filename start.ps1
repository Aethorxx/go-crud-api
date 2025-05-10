# Остановка и удаление существующих контейнеров
Write-Host "Останавливаем существующие контейнеры..." -ForegroundColor Yellow
docker-compose down

# Удаление старых образов
Write-Host "Удаляем старые образы..." -ForegroundColor Yellow
docker-compose rm -f
docker system prune -f

# Сборка и запуск контейнеров
Write-Host "Запускаем приложение..." -ForegroundColor Green
docker-compose up --build -d

# Ожидание запуска сервисов
Write-Host "Ожидаем запуска сервисов..." -ForegroundColor Yellow
Start-Sleep -Seconds 5

# Проверка статуса контейнеров
Write-Host "Статус контейнеров:" -ForegroundColor Cyan
docker-compose ps

Write-Host "`nПриложение запущено!" -ForegroundColor Green
Write-Host "API доступно по адресу: http://localhost:8080" -ForegroundColor Cyan
Write-Host "База данных доступна по адресу: localhost:5432" -ForegroundColor Cyan 