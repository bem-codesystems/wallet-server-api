FROM golang:1.16-alpine3.13 AS builder

ADD . /go/src/wallet-server

WORKDIR /go/src/wallet-server

RUN go mod init wallet-server

RUN go mod tidy

RUN go build -o wallet-server main.go

EXPOSE 4000

CMD ["./wallet-server"]
