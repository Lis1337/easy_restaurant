FROM golang:alpine AS builder
LABEL authors="lis4991@gmail.com"
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .
RUN go build -o main .

FROM ubuntu:latest

WORKDIR /app
COPY --from=builder /app/main .
RUN chmod +x ./main

CMD ["./main"]
