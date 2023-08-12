FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd/server && go build -o /app/main


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD [ "./main" ]