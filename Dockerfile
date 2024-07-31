#Build stage
FROM golang:1.22.5-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/main.go

#Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
RUN mkdir migrations
COPY --from=builder /app/internal/database/migrations /app/migrations
COPY --from=builder /app/.env .



EXPOSE 8080