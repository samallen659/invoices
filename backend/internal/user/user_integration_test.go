package user_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/samallen659/invoices/backend/internal/auth"
	"github.com/samallen659/invoices/backend/internal/db"
	"github.com/samallen659/invoices/backend/internal/user"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

const (
	DB_NAME = "usr_test"
	DB_USER = "postgres"
	DB_PASS = "test_password"
)

func TestUserDomain(t *testing.T) {
	ctx := context.Background()

	connStr := setupDB(t, ctx)

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

func setupDB(t testing.TB, ctx context.Context) string {
	t.Helper()

	pgContainer, err := createDBContainer(ctx)
	if err != nil {
		t.Errorf("Failed created DB for integration test: %s", err.Error())
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Errorf("Failed created DB for integration test: %s", err.Error())
	}
	ports, err := pgContainer.Container.Ports(ctx)
	if err != nil {
		t.Errorf("Failed created DB for integration test: %s", err.Error())
	}

	hostPort := ports["5432/tcp"][0].HostPort

	err = migrateDB(hostPort)
	if err != nil {
		t.Errorf("Failed created DB for integration test: %s", err.Error())
	}

	return connStr
}

func createDBContainer(ctx context.Context) (*postgres.PostgresContainer, error) {
	pgContainer, err := postgres.RunContainer(ctx, testcontainers.WithImage("postgres:15-alpine"),
		postgres.WithDatabase(DB_NAME), postgres.WithUsername(DB_USER), postgres.WithPassword(DB_PASS),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	return pgContainer, nil
}

func migrateDB(hostPort string) error {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get path")
	}
	pathToMigrationFile := filepath.Clean(filepath.Join(filepath.Dir(path), "../../db/migrations/"))

	databaseURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", DB_USER, DB_PASS, hostPort, DB_NAME)
	m, err := migrate.New(fmt.Sprintf("file:%s", pathToMigrationFile), databaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	log.Println("migration done")
	return nil
}

func createHTTPRequest(t testing.TB, method string, route string, body *[]byte) *http.Request {
	t.Helper()

	var req *http.Request
	var err error
	if body != nil {
		b := bytes.NewReader(*body)
		req, err = http.NewRequest(method, route, b)
		if err != nil {
			t.Fatalf("failed creating http request: %s", err.Error())
		}
	} else {
		req, err = http.NewRequest(method, route, nil)
		if err != nil {
			t.Fatalf("failed creating http request: %s", err.Error())
		}
	}

	return req
}

func assertHTTPStatusCode(t testing.TB, desiredCode int, returnedCode int) {
	t.Helper()

	if returnedCode != desiredCode {
		t.Errorf("Unexpected error code, received: %d expected: %d", returnedCode, desiredCode)
	}
}
