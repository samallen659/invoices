package transport

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samallen659/invoices/backend/internal/invoice"
)

type Server struct {
	invHandler *invoice.Handler
	Router     *mux.Router
}

func NewServer(invHandler *invoice.Handler) (*Server, error) {
	router := mux.NewRouter()

	router.HandleFunc("/invoice/{id}", invHandler.HandleGetByID).Methods(http.MethodGet)
	router.HandleFunc("/invoice/{id}", invHandler.HandleUpdate).Methods(http.MethodPut)
	router.HandleFunc("/invoice/{id}", invHandler.HandleDelete).Methods(http.MethodDelete)
	router.HandleFunc("/invoice", invHandler.HandleGetAll).Methods(http.MethodGet)
	router.HandleFunc("/invoice", invHandler.HandleStore).Methods(http.MethodPost)

	return &Server{invHandler: invHandler, Router: router}, nil
}

func (s *Server) Serve(port string) error {
	if err := http.ListenAndServe(port, s.Router); err != nil {
		return err
	}
	return nil
}
