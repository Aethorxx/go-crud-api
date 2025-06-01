
## О проекте
Это базовый REST API с CRUD-операциями для управления пользователями и их заказами. Использует Gin, GORM, PostgreSQL и JWT для аутентификации, ну и конечно же Docker.

## Как запустить?

### Через Docker
1. Скопируйте `.env.example` в `.env`
2. Запустите:
```bash
docker-compose up --build
```

### Тестирование API через Postman
1. Импортируйте коллекцию `go-crud-api.postman_collection.json` в Postman
2. Создайте пользователя.
3. Войдите.
4. Тестируйте остальные эндпоинты.

### Доступные эндпоинты
```
POST /auth/register - Регистрация
POST /auth/login - Вход
GET    /users - Список пользователей
POST   /users - Создание пользователя
GET    /users/:id - Получение пользователя
PUT    /users/:id - Обновление пользователя
DELETE /users/:id - Удаление пользователя
GET    /users/:id/orders - Заказы пользователя
POST   /users/:id/orders - Создание заказа
GET    /users/:id/orders/:order_id - Получение заказа
PUT    /users/:id/orders/:order_id - Обновление заказа
DELETE /users/:id/orders/:order_id - Удаление заказа
```

### Важно
- Все запросы (кроме регистрации и входа) требуют JWT токен
- Пользователь может работать только со своими заказами
- При попытке доступа к чужим заказам получите ошибку доступа
- ID пользователя в URL должен совпадать с ID в токене

### Логи
```bash
docker-compose logs -f api
``` 
