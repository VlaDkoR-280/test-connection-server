# Этап сборки
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем go.mod и загружаем зависимости
COPY go.mod ./
RUN go mod download

# Копируем исходный код
COPY . ./

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o health-monitor ./cmd/server

# Финальный этап
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из этапа сборки
COPY --from=builder /app/health-monitor .

# Создаем непривилегированного пользователя
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup && \
    chown appuser:appgroup /root/health-monitor

USER appuser

# Порт по умолчанию
EXPOSE 8080

# Запускаем приложение
CMD ["./health-monitor"]