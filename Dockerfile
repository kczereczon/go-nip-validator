FROM golang:1.24.1-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o main main.go

FROM alpine:latest

COPY --from=builder /app/main /app/main

CMD ["/app/main"]