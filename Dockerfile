FROM golang:1.21-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN cd cmd/server && go build -o /app/main

EXPOSE 8080

CMD [ "./main" ]