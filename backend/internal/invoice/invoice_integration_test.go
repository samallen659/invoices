package invoice_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	DB_NAME = "inv_test"
	DB_USER = "postgres"
	DB_PASS = "test_password"
)

func TestInvoiceDomain(t *testing.T) {
	ctx := context.Background()

	connStr := setupInvoiceDB(t, ctx)

	invHandler, err := setupInvoiceDomain(connStr)
	if err != nil {
		t.Errorf("failed setting up invoice domain for integration test: %s", err.Error())
	}

	server := newServer(invHandler)

	invReq := createInvoiceRequest(t)
	var id uuid.UUID
	t.Run("HandleGetByID returns status code 400 for bad ID path variable", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := createHTTPRequest(t, http.MethodGet, "/invoice/123", nil)
		server.Router.ServeHTTP(w, req)

		assertHTTPStatusCode(t, w.Code, http.StatusBadRequest)

		var receivedError invoice.ErrorResponse
		json.Unmarshal(w.Body.Bytes(), &receivedError)

		if receivedError.Error != "invalid UUID length: 3" {
			t.Errorf("Unexpected error received in body")
		}
	})
	t.Run("HandleStore saves new invoice", func(t *testing.T) {
		w := httptest.NewRecorder()
		invReqJson, _ := json.Marshal(invReq)
		req := createHTTPRequest(t, http.MethodPost, "/invoice", &invReqJson)
		server.Router.ServeHTTP(w, req)

		assertHTTPStatusCode(t, w.Code, http.StatusOK)

		body := unmarshalResponse(t, w.Body.Bytes())
		id = body.Invoice[0].ID

		if body.Invoice[0].Description != invReq.Description {
			t.Error("Returned Invoice does not match supplied invoice request")
		}
	})
	t.Run("HandleUpdate edits the invoice", func(t *testing.T) {
		w := httptest.NewRecorder()
		invReq.Description = "Changed Description"
		invReqJson, _ := json.Marshal(invReq)
		req := createHTTPRequest(t, http.MethodPut, fmt.Sprintf("/invoice/%s", id.String()), &invReqJson)
		server.Router.ServeHTTP(w, req)

		assertHTTPStatusCode(t, w.Code, http.StatusOK)

		body := unmarshalResponse(t, w.Body.Bytes())

		if body.Invoice[0].Description != "Changed Description" {
			t.Error("Returned Invoice has not had the description correctly edited")
		}
	})
	t.Run("HandleGetByID returns correct Invoice", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := createHTTPRequest(t, http.MethodGet, fmt.Sprintf("/invoice/%s", id.String()), nil)
		server.Router.ServeHTTP(w, req)

		assertHTTPStatusCode(t, w.Code, http.StatusOK)

		body := unmarshalResponse(t, w.Body.Bytes())

		if body.Invoice[0].ID.String() != id.String() {
			t.Error("Returned Invoice ID is incorrect")
		}
	})
	t.Run("HandleDelete deletes Invoice", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := createHTTPRequest(t, http.MethodDelete, fmt.Sprintf("/invoice/%s", id.String()), nil)
		server.Router.ServeHTTP(w, req)

		assertHTTPStatusCode(t, w.Code, http.StatusOK)

		//check deleted invoice not returned
		w = httptest.NewRecorder()
		req = createHTTPRequest(t, http.MethodGet, fmt.Sprintf("/invoice/%s", id.String()), nil)

		server.Router.ServeHTTP(w, req)

		assertHTTPStatusCode(t, w.Code, http.StatusBadRequest)
	})
}

type Server struct {
	Router *mux.Router
}

func newServer(invHandler *invoice.Handler) *Server {
	router := mux.NewRouter()
	router.HandleFunc("/invoice/{id}", invHandler.HandleGetByID).Methods(http.MethodGet)
	router.HandleFunc("/invoice/{id}", invHandler.HandleUpdate).Methods(http.MethodPut)
	router.HandleFunc("/invoice/{id}", invHandler.HandleDelete).Methods(http.MethodDelete)
	router.HandleFunc("/invoice", invHandler.HandleGetAll).Methods(http.MethodGet)
	router.HandleFunc("/invoice", invHandler.HandleStore).Methods(http.MethodPost)

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

func setupInvoiceDomain(connStr string) (*invoice.Handler, error) {
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

func setupInvoiceDB(t testing.TB, ctx context.Context) string {
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

func unmarshalResponse(t testing.TB, bodyBytes []byte) invoice.InvoiceResponse {
	t.Helper()

	var body invoice.InvoiceResponse
	err := json.Unmarshal(bodyBytes, &body)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err.Error())
	}

	return body
}
