# Используем официальный образ Go как базовый
FROM golang:1.24-alpine as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники приложения
COPY . .

# Собираем основное приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Собираем миграции
RUN CGO_ENABLED=0 GOOS=linux go build -o migrate ./migration

# Начинаем новую стадию сборки на основе минимального образа
FROM alpine:latest

# Устанавливаем ca-certificates для HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем исполняемые файлы из builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

# Создаем entrypoint скрипт
RUN echo '#!/bin/sh' > /entrypoint.sh && \
    echo 'set -e' >> /entrypoint.sh && \
    echo 'echo "Running database migrations..."' >> /entrypoint.sh && \
    echo './migrate' >> /entrypoint.sh && \
    echo 'echo "Migrations completed successfully"' >> /entrypoint.sh && \
    echo 'echo "Starting application..."' >> /entrypoint.sh && \
    echo 'exec ./main' >> /entrypoint.sh && \
    chmod +x /entrypoint.sh

# Определяем переменные окружения с значениями по умолчанию
ENV DADATA_API_KEY="" \
    DADATA_SECRET_KEY="" \
    JWT_SECRET="" \
    DSN=""

# Открываем порт 8080
EXPOSE 8080

# Используем entrypoint для запуска миграций и приложения
ENTRYPOINT ["/entrypoint.sh"]
