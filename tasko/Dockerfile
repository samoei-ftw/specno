FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .
RUN go build -o /specno-be cmd/user-service/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /specno-be /specno-be

EXPOSE 8080

CMD ["/specno-be"]