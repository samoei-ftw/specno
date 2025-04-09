FROM golang:1.24.1 AS project_build

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod tidy

COPY ./common /app/common
COPY ./gateways /app/gateways
COPY ./database /app/database
COPY ./wait-for-it.sh /wait-for-it.sh

WORKDIR /app/task_service
COPY ./task_service /app/task_service

ENV GOARCH=amd64
ENV GOOS=linux

RUN go build -o /task-api ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache bash ca-certificates curl

COPY --from=task_build /task-api  /task-api 
COPY --from=task_build /wait-for-it.sh /wait-for-it.sh
WORKDIR /
RUN chmod +x ./wait-for-it.sh

EXPOSE 8083
CMD ["bash", "-c", "./wait-for-it.sh project-service:8082 -- /task-api"]