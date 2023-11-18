package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	// "github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/samallen659/invoices/backend/internal/invoice"
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

	svc, err := invoice.NewService(repo)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	paymentDue, _ := time.Parse("2006-01-02", "2023-11-20")
	invRq := invoice.InvoiceRequest{
		PaymentDue:   paymentDue,
		Description:  "testing",
		PaymentTerms: 2,
		ClientName:   "Sam Allen",
		ClientEmail:  "sam@email.com",
		Status:       "pending",
		ClientAddress: struct {
			Street   string `json:"street"`
			City     string `json:"city"`
			PostCode string `json:"postCode"`
			Country  string `json:"country"`
		}{
			Street:   "15 Eastfield Road",
			City:     "Doncaster",
			PostCode: "DN9 1JQ",
			Country:  "United Kingdom",
		},
		SenderAddress: struct {
			Street   string `json:"street"`
			City     string `json:"city"`
			PostCode string `json:"postCode"`
			Country  string `json:"country"`
		}{
			Street:   "77 High Street",
			City:     "Doncaster",
			PostCode: "DN9 1JS",
			Country:  "United Kingdom",
		},
		Items: []struct {
			Name     string  `json:"name"`
			Quantity int     `json:"quantity"`
			Price    float64 `json:"price"`
		}{{
			Name:     "Web Design",
			Quantity: 1,
			Price:    1499.99,
		}, {
			Name:     "Web Development",
			Quantity: 1,
			Price:    3649.99,
		}},
	}

	inv, err := svc.NewInvoice(ctx, invRq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inv)

}
