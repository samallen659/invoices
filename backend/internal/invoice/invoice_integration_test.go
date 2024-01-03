package invoice_test

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/samallen659/invoices/backend/internal/db"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"github.com/samallen659/invoices/backend/internal/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	DB_NAME = "inv_test"
	DB_USER = "postgres"
	DB_PASS = "test_password"
)

func TestInvoiceDomain(t *testing.T) {
	ctx := context.Background()

	connStr := test.SetupDB(t, ctx, DB_PASS, DB_USER, DB_NAME)

	invHandler, err := setupInvoiceDomain(connStr)
	if err != nil {
		t.Errorf("failed setting up invoice domain for integration test: %s", err.Error())
	}

	server := newServer(invHandler)

	invReq := createInvoiceRequest(t)
	var id uuid.UUID
	t.Run("HandleGetByID returns status code 400 for bad ID path variable", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := test.CreateHTTPRequest(t, http.MethodGet, "/invoice/123", nil)
		server.Router.ServeHTTP(w, req)

		test.AssertHTTPStatusCode(t, w.Code, http.StatusBadRequest)

		var receivedError invoice.ErrorResponse
		json.Unmarshal(w.Body.Bytes(), &receivedError)

		if receivedError.Error != "invalid UUID length: 3" {
			t.Errorf("Unexpected error received in body")
		}
	})
	t.Run("HandleStore saves new invoice", func(t *testing.T) {
		w := httptest.NewRecorder()
		invReqJson, _ := json.Marshal(invReq)
		req := test.CreateHTTPRequest(t, http.MethodPost, "/invoice", &invReqJson)
		server.Router.ServeHTTP(w, req)

		test.AssertHTTPStatusCode(t, w.Code, http.StatusOK)

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
		req := test.CreateHTTPRequest(t, http.MethodPut, fmt.Sprintf("/invoice/%s", id.String()), &invReqJson)
		server.Router.ServeHTTP(w, req)

		test.AssertHTTPStatusCode(t, w.Code, http.StatusOK)

		body := unmarshalResponse(t, w.Body.Bytes())

		if body.Invoice[0].Description != "Changed Description" {
			t.Error("Returned Invoice has not had the description correctly edited")
		}
	})
	t.Run("HandleGetByID returns correct Invoice", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := test.CreateHTTPRequest(t, http.MethodGet, fmt.Sprintf("/invoice/%s", id.String()), nil)
		server.Router.ServeHTTP(w, req)

		test.AssertHTTPStatusCode(t, w.Code, http.StatusOK)

		body := unmarshalResponse(t, w.Body.Bytes())

		if body.Invoice[0].ID.String() != id.String() {
			t.Error("Returned Invoice ID is incorrect")
		}
	})
	t.Run("HandleDelete deletes Invoice", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := test.CreateHTTPRequest(t, http.MethodDelete, fmt.Sprintf("/invoice/%s", id.String()), nil)
		server.Router.ServeHTTP(w, req)

		test.AssertHTTPStatusCode(t, w.Code, http.StatusOK)

		//check deleted invoice not returned
		w = httptest.NewRecorder()
		req = test.CreateHTTPRequest(t, http.MethodGet, fmt.Sprintf("/invoice/%s", id.String()), nil)

		server.Router.ServeHTTP(w, req)

		test.AssertHTTPStatusCode(t, w.Code, http.StatusBadRequest)
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
	conn, err := db.ConnectDB(connStr)
	if err != nil {
		return nil, err
	}
	repo := invoice.NewPostgresRespository(conn)

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

func unmarshalResponse(t testing.TB, bodyBytes []byte) invoice.InvoiceResponse {
	t.Helper()

	var body invoice.InvoiceResponse
	err := json.Unmarshal(bodyBytes, &body)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err.Error())
	}

	return body
}
