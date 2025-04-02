FROM golang:1.24.1 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy 

COPY . .
WORKDIR /app/user_service
COPY ./common /app/common

# Set environment variables for cross-compilation
ENV GOARCH=amd64
ENV GOOS=linux
RUN go build -o /user-api ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates


COPY --from=build /user-api /user-api


COPY .env /app/.env

RUN chmod +x /user-api

EXPOSE 8080

CMD ["/user-api"]