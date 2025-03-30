FROM golang:1.23 AS build

WORKDIR /app

RUN go version
COPY go.mod ./
COPY . .

# Set environment variables for cross-compilation
ENV GOARCH=amd64
ENV GOOS=linux
RUN go build -o /specno-be cmd/user-service/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates


COPY --from=build /specno-be /specno-be


COPY .env .env

RUN chmod +x /specno-be

EXPOSE 8080

CMD ["/specno-be"]