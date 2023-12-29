package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"github.com/samallen659/invoices/backend/internal/session"
	"github.com/samallen659/invoices/backend/internal/transport"
	"github.com/samallen659/invoices/backend/internal/user"
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

	session.New(os.Getenv("SESSION_SECRET"))

	//Invoice setup
	invRepo, err := invoice.NewPostgresRespository(postgresConnStr)
	if err != nil {
		log.Fatal(err)
	}
	invSvc, err := invoice.NewService(invRepo)
	if err != nil {
		log.Fatal(err)
	}
	invHandler, err := invoice.NewHandler(invSvc)
	if err != nil {
		log.Fatal(err)
	}

	//User setup
	usrAuth, err := user.NewAuthenticator()
	if err != nil {
		log.Fatal(err)
	}
	usrSvc, err := user.NewService(usrAuth)
	if err != nil {
		log.Fatal(err)
	}
	usrHandler, err := user.NewHandler(usrSvc)
	if err != nil {
		log.Fatal(err)
	}

	server, err := transport.NewServer(invHandler, usrHandler)
	if err != nil {
		log.Fatal(err)
	}

	server.Serve(":8080")
}
