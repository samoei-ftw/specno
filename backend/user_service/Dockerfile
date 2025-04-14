FROM golang:1.24.1 AS user_build

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod tidy

COPY ./common /app/common
COPY ./gateways /app/gateways
COPY ./database /app/database

WORKDIR /app/user_service
COPY ./user_service /app/user_service

ENV GOARCH=amd64
ENV GOOS=linux

RUN go build -o /user-api ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache bash ca-certificates curl

COPY --from=user_build /user-api /user-api
RUN chmod +x /user-api

EXPOSE 8080

CMD ["/user-api"]