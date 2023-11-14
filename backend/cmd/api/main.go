package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
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
	// postgresPort := os.Getenv("POSTGRES_PORT")

	postgresConnStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=localhost",
		postgresUser, postgresDB, postgresPass)

	repo, err := invoice.NewPostgresRespository(postgresConnStr)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	inv, err := repo.GetInvoiceByID(ctx, "c1aa569f-87ff-4b5c-9c27-45fd24e29642")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inv)

}
