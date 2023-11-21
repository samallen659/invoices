package transport

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samallen659/invoices/backend/internal/invoice"
)

type Server struct {
	invHandler *invoice.Handler
	router     *mux.Router
}

func NewServer(invHandler *invoice.Handler) (*Server, error) {
	router := mux.NewRouter()

	router.HandleFunc("/invoice/{id}", invHandler.HandleGetByID)

	return &Server{invHandler: invHandler, router: router}, nil
}

func (s *Server) Serve(port string) error {
	if err := http.ListenAndServe(port, s.router); err != nil {
		return err
	}
	return nil
}
