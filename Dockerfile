FROM golang:alpine AS builder
LABEL authors="lis4991@gmail.com"
WORKDIR /src

#COPY go.mod go.sum .
COPY go.mod .
COPY . .
RUN go build -o src . && go build -C app

FROM ubuntu:latest
COPY --from=builder /src /src
RUN chmod +x /src

CMD ["./src/main"]
#ENV PORT=8080
#EXPOSE $PORT
