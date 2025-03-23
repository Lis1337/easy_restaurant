FROM golang:alpine AS builder
LABEL authors="lis4991@gmail.com"

ADD . /dockerdev
WORKDIR /dockerdev

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go build -gcflags="all=-N -l" -o /server

FROM ubuntu:latest

EXPOSE 8000 40000

WORKDIR /
COPY --from=builder /go/bin/dlv /
COPY --from=builder /server /

COPY --from=builder /go/bin/goose /usr/local/bin/goose

USER 1000:1000

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]
