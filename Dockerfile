FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o main cmd/app/main.go

FROM alpine:latest

COPY --from=builder app/.env /.env
COPY --from=builder app/main /main

ENTRYPOINT ["/main"]