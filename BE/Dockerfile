FROM golang:1.22-alpine AS build
WORKDIR /app

# Install required system dependencies (combined into one RUN to reduce layers)
RUN apk add --no-cache \
    git \
    sqlite \
    nodejs \
    npm \
    gcc \
    musl-dev

# Set environment variables
ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOOSE_DRIVER=sqlite3 \
    GOOSE_DBSTRING=/app/tmp/app.db \
    GOOSE_MIGRATION_DIR=/app/sqlite/migration

# Install Go tools
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/a-h/templ/cmd/templ@latest && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Create necessary directories
RUN mkdir -p /app/tmp /app/sqlite/migration

# Copy the rest of the code
COPY . .

# Generate, migrate, and build
RUN sqlc generate && \
    goose up && \
    go build -buildvcs=false -o /app/tmp/main

EXPOSE 8080
ENTRYPOINT ["/app/tmp/main"]