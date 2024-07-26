# Используем официальный образ Golang версии 1.21
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем весь исходный код в контейнер
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/main.go

# Используем минимальный образ Alpine для запуска
FROM alpine:latest

# Копируем бинарник из builder контейнера
COPY --from=builder /app/main /app/main

# Копируем файл .env и entrypoint.sh в рабочую директорию контейнера
COPY .env /app/.env
COPY entrypoint.sh /app/entrypoint.sh

# Устанавливаем рабочую директорию
WORKDIR /app

# Даем права на выполнение скрипту entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Определяем переменные окружения
ENV PORT=${WEB_PORT}

# Открываем порт
EXPOSE ${WEB_PORT}

# Команда для запуска скрипта entrypoint.sh
CMD ["/app/entrypoint.sh"]
