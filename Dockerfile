FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD [ "go", "run",  "cmd/server/main.go"]