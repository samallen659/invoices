package main

import (
	"fmt"
	"github.com/joho/godotenv"
	transport "github.com/samallen659/invoices/backend/internal"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPass := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")

	postgresConnStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=localhost",
		postgresUser, postgresDB, postgresPass)

	repo, err := invoice.NewPostgresRespository(postgresConnStr)
	if err != nil {
		log.Fatal(err)
	}

	svc, err := invoice.NewService(repo)
	if err != nil {
		log.Fatal(err)
	}

	invHandler, err := invoice.NewHandler(svc)
	if err != nil {
		log.Fatal(err)
	}

	server, err := transport.NewServer(invHandler)
	if err != nil {
		log.Fatal(err)
	}

	server.Serve(":8080")
}
