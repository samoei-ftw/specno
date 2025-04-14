FROM golang:1.24.1 AS project_build

ADD https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz /tmp/
RUN tar -C /usr/local/bin -xzvf /tmp/dockerize-linux-amd64-v0.6.1.tar.gz

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod tidy

COPY ./common /app/common
COPY ./gateways /app/gateways
COPY ./user_service /app/user_service
COPY ./database /app/database
WORKDIR /app/project_service
COPY ./project_service /app/project_service

ENV GOARCH=amd64
ENV GOOS=linux

RUN go build -o /project-api ./cmd/main.go
FROM alpine:latest

RUN apk add --no-cache bash ca-certificates curl

COPY --from=project_build /project-api /project-api
COPY --from=project_build /usr/local/bin/dockerize /usr/local/bin/dockerize

EXPOSE 8082

CMD ["sh", "-c", "dockerize -wait tcp://user-service:8080 -timeout 30s /project-api -run-migrations=false"]