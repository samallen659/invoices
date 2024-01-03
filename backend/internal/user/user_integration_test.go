package user_test

import (
	"context"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/samallen659/invoices/backend/internal/auth"
	"github.com/samallen659/invoices/backend/internal/db"
	"github.com/samallen659/invoices/backend/internal/test"
	"github.com/samallen659/invoices/backend/internal/user"
	"net/http"
	"testing"
)

const (
	DB_NAME = "usr_test"
	DB_USER = "postgres"
	DB_PASS = "test_password"
)

func TestUserDomain(t *testing.T) {
	ctx := context.Background()

	connStr := test.SetupDB(t, ctx, DB_PASS, DB_USER, DB_NAME)

	usrHandler, err := setupUserDomain(connStr)
	if err != nil {
		t.Errorf("failed setting up user domain for integration test: %s", err.Error())
	}

	server := newServer(usrHandler)
	fmt.Println(server)
	t.Fatal("stop")
}

func setupUserDomain(connStr string) (*user.Handler, error) {
	conn, err := db.ConnectDB(connStr)
	if err != nil {
		return nil, err
	}
	repo := user.NewPostgresRepository(conn)
	auth, err := auth.NewAuthenticator()
	if err != nil {
		return nil, err
	}
	svc, err := user.NewService(auth, repo)
	if err != nil {
		return nil, err
	}
	handler, err := user.NewHandler(svc)
	if err != nil {
		return nil, err
	}

	return handler, nil
}

type Server struct {
	Router *mux.Router
}

func newServer(usrHandler *user.Handler) *Server {
	router := mux.NewRouter()
	router.HandleFunc("/user/login", usrHandler.HandleLogin).Methods(http.MethodGet)
	router.HandleFunc("/user/callback", usrHandler.HandleCallback).Methods(http.MethodGet)
	router.HandleFunc("/user/logout", usrHandler.HandleLogout).Methods(http.MethodGet)
	router.HandleFunc("/user", usrHandler.HandleGetUser).Methods(http.MethodGet)
	return &Server{
		Router: router,
	}
}

func (s *Server) Serve(port string) error {
	if err := http.ListenAndServe(port, s.Router); err != nil {
		return err
	}
	return nil
}
