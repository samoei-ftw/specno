FROM golang:1.24.1 AS project_build

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod tidy

COPY ./common /app/common
COPY ./gateways /app/gateways
COPY ./user_service /app/user_service
COPY ./database /app/database
COPY ./wait-for-it.sh /wait-for-it.sh


WORKDIR /app/project_service
COPY ./project_service /app/project_service

ENV GOARCH=amd64
ENV GOOS=linux

RUN go build -o /project-api ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache bash ca-certificates curl

COPY --from=project_build /project-api /project-api
COPY --from=project_build /wait-for-it.sh /wait-for-it.sh
WORKDIR /
RUN chmod +x ./wait-for-it.sh

EXPOSE 8082
CMD ["bash", "-c", "./wait-for-it.sh user-service:8080 -- /project-api -run-migrations=false"]