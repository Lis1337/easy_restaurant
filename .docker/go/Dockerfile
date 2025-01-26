FROM golang:alpine AS builder
LABEL authors="lis4991@gmail.com"
WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
COPY . ./

RUN go build -o main .

FROM ubuntu:latest
WORKDIR /app

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/main .

RUN chmod +x ./main

CMD ["./main"]
