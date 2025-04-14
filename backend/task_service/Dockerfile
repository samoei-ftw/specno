FROM golang:1.24.1 AS task_build

ADD https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz /tmp/
RUN tar -C /usr/local/bin -xzvf /tmp/dockerize-linux-amd64-v0.6.1.tar.gz

WORKDIR /app

COPY ./go.mod ./go.sum ./
COPY ./common /app/common
COPY ./gateways /app/gateways
COPY ./user_service /app/user_service
COPY ./project_service /app/project_service
COPY ./database /app/database
COPY ./task_service /app/task_service

RUN go mod tidy

WORKDIR /app/task_service

ENV GOARCH=amd64
ENV GOOS=linux

RUN go build -o /task-api ./cmd/main.go
FROM alpine:latest

RUN apk add --no-cache bash ca-certificates curl

COPY --from=task_build /task-api /task-api
COPY --from=task_build /usr/local/bin/dockerize /usr/local/bin/dockerize

EXPOSE 8083

CMD ["sh", "-c", "dockerize -wait tcp://project-service:8082 -timeout 30s /task-api -run-migrations=false"]