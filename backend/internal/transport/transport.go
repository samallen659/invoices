package transport

import (
	"encoding/gob"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/samallen659/invoices/backend/internal/auth"
	"github.com/samallen659/invoices/backend/internal/invoice"
	"github.com/samallen659/invoices/backend/internal/session"
	"github.com/samallen659/invoices/backend/internal/user"
	"net/http"
)

type Server struct {
	invHandler  *invoice.Handler
	Router      *mux.Router
	methods     handlers.CORSOption
	credentials handlers.CORSOption
	origins     handlers.CORSOption
}

var authenticator *auth.Authenticator

type handler func(http.ResponseWriter, *http.Request)

func NewServer(invHandler *invoice.Handler, usrHandler *user.Handler, authen *auth.Authenticator) (*Server, error) {
	router := mux.NewRouter()

	authenticator = authen

	router.HandleFunc("/invoice/{id}", invHandler.HandleGetByID).Methods(http.MethodGet)
	router.HandleFunc("/invoice/{id}", invHandler.HandleUpdate).Methods(http.MethodPut)
	router.HandleFunc("/invoice/{id}", invHandler.HandleDelete).Methods(http.MethodDelete)
	router.HandleFunc("/invoice", authMiddleware(invHandler.HandleGetAll)).Methods(http.MethodGet)
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

func authMiddleware(next handler) handler {
	return (func(w http.ResponseWriter, r *http.Request) {
		ses, err := session.Get(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if ses.Values["profile"] == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	})
}
