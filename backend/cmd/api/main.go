package main

import (
	"fmt"
	"github.com/joho/godotenv"
	transport "github.com/samallen659/invoices/backend/internal"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"github.com/samallen659/invoices/backend/internal/user"
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
	cognitoClientID := os.Getenv("COGNITO_CLIENT_ID")
	cognitoClientSecret := os.Getenv("COGNITO_CLIENT_SECRET")

	postgresConnStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=localhost",
		postgresUser, postgresDB, postgresPass)

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
	usrAuth, err := user.NewCognitoAuthentication(cognitoClientID, cognitoClientSecret)
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
