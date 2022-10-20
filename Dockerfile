FROM golang:alpine AS builder

WORKDIR /app

COPY . ./

RUN go build -o main main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 4000

CMD ["/app/main"]