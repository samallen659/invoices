#!make
include .env

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go
# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Start postgres container
postgres:
	@docker-compose up -d

# SQL Migrate up
migrate:
	@migrate -source file://db/migrations -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable up

.PHONY: all build run test clean postgres migrate
