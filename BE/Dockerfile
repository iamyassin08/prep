FROM golang:1.22-alpine

WORKDIR /app

# Install only the necessary dependencies
RUN apk add --no-cache sqlite gcc musl-dev

# Enable CGO for SQLite
ENV CGO_ENABLED=1

# Create necessary directories
RUN mkdir -p /app/tmp /app/sqlite/migration

# Copy everything
COPY . .

# Download dependencies
RUN go mod download

# Build and prepare the app
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest && \
    
RUN sqlc generate
ENV GOOSE_DRIVER=sqlite3 
ENV GOOSE_DBSTRING=/app/tmp/app.db 
RUN goose -dir /app/sqlite/migration up
RUN go build -o /app/tmp/main

EXPOSE 8080

CMD ["/app/tmp/main"]