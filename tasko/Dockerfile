FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o /go-app cmd/user-service/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /go-app /go-app

EXPOSE 8080

CMD ["/go-app"]