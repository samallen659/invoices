package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	ctx := context.Background()
	inv, err := repo.GetInvoiceByID(ctx, "c1aa569f-87ff-4b5c-9c27-45fd24e29642")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inv)

	inv2 := invoice.NewInvoice()
	inv2.SetPaymentDue(time.Now().Add(240 * time.Hour))
	inv2.SetDescription("Landing Page Design")
	inv2.SetPaymentTerms(30)
	inv2.Client.ClientName = "Thomas Wayne"
	inv2.Client.ClientEmail = "thomas@dc.com"
	inv2.SetStatus(invoice.STATUS_PENDING)
	inv2.SenderAddress.Street = "19 Union Terrace"
	inv2.SenderAddress.City = "London"
	inv2.SenderAddress.PostCode = "E1 3EZ"
	inv2.SenderAddress.Country = "United Kingdom"
	inv2.ClientAddress.Street = "3964 Queens Lane"
	inv2.ClientAddress.City = "Gotham"
	inv2.ClientAddress.PostCode = "60457"
	inv2.ClientAddress.Country = "United States of America"
	inv2.InvoiceItems = append(inv2.InvoiceItems, invoice.InvoiceItem{
		Item:     invoice.Item{Name: "Web Design", Price: 6155.91},
		Quantity: 1,
		Total:    6155.91,
	})
	inv2.Total = 6155.91

	err = repo.StoreInvoice(ctx, *inv2)
	if err != nil {
		log.Fatal(err)
	}

	invs, err := repo.GetAllInvoices(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, in := range invs {
		fmt.Println("invoice....")
		fmt.Println(*in)
	}
}
