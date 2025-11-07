# Swagger Documentation для ProxyDB User Repository

## Описание

Swagger документация для API управления пользователями в ProxyDB.

## Доступные эндпоинты

### Users API

#### GET /api/users
Получить список всех пользователей с пагинацией
- **Query параметры:**
  - `limit` (required): Количество записей (например, 10)
  - `offset` (required): Смещение (например, 0)
- **Responses:**
  - 200: Список пользователей + header `X-Total-Count`
  - 400: Некорректные параметры запроса
  - 500: Внутренняя ошибка сервера

#### GET /api/users/{id}
Получить пользователя по ID
- **Path параметры:**
  - `id` (required): ID пользователя
- **Responses:**
  - 200: Информация о пользователе
  - 404: Пользователь не найден

#### POST /api/users
Создать нового пользователя
- **Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Responses:**
  - 200: Пользователь создан
  - 400: Некорректные данные

#### PUT /api/users/{id}
Обновить данные пользователя
- **Path параметры:**
  - `id` (required): ID пользователя
- **Body:**
  ```json
  {
    "email": "newemail@example.com",
    "password": "newpassword123"
  }
  ```
- **Responses:**
  - 200: Пользователь обновлен
  - 400: Некорректные данные

#### DELETE /api/users/{id}
Удалить пользователя
- **Path параметры:**
  - `id` (required): ID пользователя
- **Responses:**
  - 200: Пользователь удален
  - 400: Ошибка при удалении

## Модели данных

### User
```json
{
  "id": 1,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z",
  "deleted_at": "2023-01-01T00:00:00Z",
  "email": "user@example.com",
  "password": "password123"
}
```

### ListResponse
```json
{
  "total": 100,
  "users": [
    {
      "id": 1,
      "email": "user@example.com",
      "password": "password123"
    }
  ]
}
```

## Просмотр документации

1. Запустите сервер:
   ```bash
   go run cmd/main.go
   ```

2. Откройте в браузере:
   ```
   http://localhost:8080/swagger/index.html
   ```

## Регенерация документации

После внесения изменений в аннотации Swagger, выполните:

```bash
swag init -g cmd/main.go -o docs --parseDependency --parseInternal
```

## Файлы документации

- `docs.go` - Сгенерированный Go код
- `swagger.json` - JSON спецификация
- `swagger.yaml` - YAML спецификация
