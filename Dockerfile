# 1. Этап сборки (используем полный образ Go)
FROM golang:1.25-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь остальной код
COPY . .

# Собираем бинарный файл. 
# Важно: мы указываем путь к main.go (cmd/main.go)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# 2. Финальный этап (используем минимальный образ alpine)
FROM alpine:latest

# Устанавливаем сертификаты (нужны для работы с БД)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем собранный бинарник из первого этапа
COPY --from=builder /app/main .

# Копируем папку миграций (она обязательна для запуска)
COPY --from=builder /app/cmd/migrations ./cmd/migrations

# Открываем порт 9091
EXPOSE 9091

# Запускаем приложение
CMD ["./main"]