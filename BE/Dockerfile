FROM golang:1.22-alpine AS build
WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux

COPY go.mod go.sum ./
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/cosmtrek/air@v1.52.0

COPY . .
RUN go mod download

RUN swag init --parseDependency
RUN swag fmt
ENTRYPOINT ["air"]
