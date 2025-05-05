#Build
FROM golang:1.23-alpine AS builder

WORKDIR /invoices

COPY . .

RUN go mod download

EXPOSE 80

RUN go build -o /api ./cmd/api

#Deploy
FROM alpine:latest

WORKDIR /

COPY --from=builder /api /api

EXPOSE 80

ENTRYPOINT ["/api"]
