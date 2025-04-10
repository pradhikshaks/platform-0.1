FROM golang:latest


WORKDIR /app/microservices
COPY . .
RUN go mod download
ENTRYPOINT go run cmd/main.go
EXPOSE 8080