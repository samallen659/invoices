#Build
FROM golang:1.21-alpine AS BUILDER

WORKDIR /invoices

COPY . .

RUN go mod download

EXPOSE 8080

RUN go build -o /api ./cmd/api

#Deploy
FROM alpine:latest

WORKDIR /

COPY --from=BUILDER /api /api

EXPOSE 8080

ENTRYPOINT ["/api"]
