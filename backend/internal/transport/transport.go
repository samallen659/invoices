package transport

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"github.com/samallen659/invoices/backend/internal/user"
)

type Server struct {
	invHandler  *invoice.Handler
	Router      *mux.Router
	methods     handlers.CORSOption
	credentials handlers.CORSOption
	origins     handlers.CORSOption
}

func NewServer(invHandler *invoice.Handler, usrHandler *user.Handler) (*Server, error) {
	router := mux.NewRouter()

	router.HandleFunc("/invoice/{id}", invHandler.HandleGetByID).Methods(http.MethodGet)
	router.HandleFunc("/invoice/{id}", invHandler.HandleUpdate).Methods(http.MethodPut)
	router.HandleFunc("/invoice/{id}", invHandler.HandleDelete).Methods(http.MethodDelete)
	router.HandleFunc("/invoice", invHandler.HandleGetAll).Methods(http.MethodGet)
	router.HandleFunc("/invoice", invHandler.HandleStore).Methods(http.MethodPost)

	router.HandleFunc("/user/login", usrHandler.HandleLogin).Methods(http.MethodGet)
	router.HandleFunc("/user/callback", usrHandler.HandleCallback).Methods(http.MethodGet)

	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{"localhost:5173"})

	gob.Register(map[string]any{})

	return &Server{
		invHandler:  invHandler,
		Router:      router,
		methods:     methods,
		credentials: credentials,
		origins:     origins}, nil
}

func (s *Server) Serve(port string) error {
	if err := http.ListenAndServe(port, handlers.CORS(s.credentials, s.methods, s.origins)(s.Router)); err != nil {
		return err
	}
	return nil
}
