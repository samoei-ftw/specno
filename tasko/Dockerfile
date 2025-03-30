FROM golang:1.23 AS build

WORKDIR /app

RUN go version
COPY go.mod ./

COPY . .
RUN go build -o /specno-be cmd/user-service/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /specno-be /specno-be

RUN chmod +x /specno-be

EXPOSE 8080
RUN ls -l /
CMD ["/specno-be"]