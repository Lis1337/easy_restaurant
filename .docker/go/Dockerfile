FROM golang:alpine AS builder
LABEL authors="lis4991@gmail.com"
WORKDIR /app
#COPY go.mod go.sum .
COPY go.mod .
COPY . .
RUN go build -o main .

FROM ubuntu:latest

WORKDIR /app
COPY --from=builder /app/main .
RUN chmod +x ./main

CMD ["./main"]
