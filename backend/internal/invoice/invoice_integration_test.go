package invoice_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	transport "github.com/samallen659/invoices/backend/internal"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

const (
	DB_NAME = "inv_test"
	DB_USER = "postgres"
	DB_PASS = "test_password"
)

func TestInvoiceDomain(t *testing.T) {
	ctx := context.Background()
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

	err = migrateDB(ctx, hostPort)
	if err != nil {
		t.Errorf("Failed created DB for integration test: %s", err.Error())
	}

	invHandler, err := setupInvoiceDomain(ctx, connStr)
	if err != nil {
		t.Errorf("failed setting up invoice domain for integration test: %s", err.Error())
	}

	server, err := transport.NewServer(invHandler)
	if err != nil {
		log.Fatal(err)
	}

	t.Run("HandleGetByID returns status code 400 for bad ID path variable", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/invoice/123", nil)
		if err != nil {
			t.Fatalf("failed to create http request: %s", err.Error())
		}
		server.Router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Unexpected error code, received: %d expected: %d", w.Code, http.StatusBadRequest)
		}
		var receivedError invoice.ErrorResponse
		json.Unmarshal(w.Body.Bytes(), &receivedError)

		if receivedError.Error != "invalid UUID length: 3" {
			t.Errorf("Unexpected error received in body")
		}
	})

}

func setupInvoiceDomain(ctx context.Context, connStr string) (*invoice.Handler, error) {
	repo, err := invoice.NewPostgresRespository(connStr)
	if err != nil {
		return nil, err
	}

	svc, err := invoice.NewService(repo)
	if err != nil {
		return nil, err
	}

	handler, err := invoice.NewHandler(svc)
	if err != nil {
		return nil, err
	}

	return handler, nil
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

func migrateDB(ctx context.Context, hostPort string) error {
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
