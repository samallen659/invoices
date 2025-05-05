package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samallen659/invoices/backend/internal/db"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"github.com/samallen659/invoices/backend/internal/transport"
	"github.com/samallen659/invoices/backend/internal/user"
)

func main() {
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPass := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresHost := os.Getenv("POSTGRES_HOST")

	postgresConnStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s",
		postgresUser, postgresDB, postgresPass, postgresHost)

	conn, err := db.ConnectDB(postgresConnStr)
	if err != nil {
		log.Fatal(err)
	}

	// Invoice setup
	invRepo := invoice.NewPostgresRespository(conn)
	invSvc, err := invoice.NewService(invRepo)
	if err != nil {
		log.Fatal(err)
	}
	invHandler, err := invoice.NewHandler(invSvc)
	if err != nil {
		log.Fatal(err)
	}

	// User setup
	usrRepo := user.NewPostgresRepository(conn)
	usrSvc, err := user.NewService(usrRepo)
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
