package main

import (
	"fmt"
	"github.com/samallen659/invoices/backend/internal/auth"
	"github.com/samallen659/invoices/backend/internal/db"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"github.com/samallen659/invoices/backend/internal/session"
	"github.com/samallen659/invoices/backend/internal/transport"
	"github.com/samallen659/invoices/backend/internal/user"
	"log"
	"os"
)

func main() {
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPass := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresHost := os.Getenv("POSTGRES_HOST")

	postgresConnStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s",
		postgresUser, postgresDB, postgresPass, postgresHost)

	session.New(os.Getenv("SESSION_SECRET"))

	conn, err := db.ConnectDB(postgresConnStr)
	if err != nil {
		log.Fatal(err)
	}

	usrAuth, err := auth.NewAuthenticator()
	if err != nil {
		log.Fatal(err)
	}

	//Invoice setup
	invRepo := invoice.NewPostgresRespository(conn)
	invSvc, err := invoice.NewService(invRepo)
	if err != nil {
		log.Fatal(err)
	}
	invHandler, err := invoice.NewHandler(invSvc)
	if err != nil {
		log.Fatal(err)
	}

	//User setup
	usrRepo := user.NewPostgresRepository(conn)
	usrSvc, err := user.NewService(usrAuth, usrRepo)
	if err != nil {
		log.Fatal(err)
	}
	usrHandler, err := user.NewHandler(usrSvc)
	if err != nil {
		log.Fatal(err)
	}

	server, err := transport.NewServer(invHandler, usrHandler, usrAuth)
	if err != nil {
		log.Fatal(err)
	}

	server.Serve(":8080")
}
