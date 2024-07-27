FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/main.go



FROM alpine:latest

COPY --from=builder /app/main /app/main

COPY .env /app/.env
COPY entrypoint.sh /app/entrypoint.sh

WORKDIR /app

RUN chmod +x /app/entrypoint.sh

ENV PORT=${WEB_PORT}

EXPOSE ${WEB_PORT}

CMD ["/app/entrypoint.sh"]
