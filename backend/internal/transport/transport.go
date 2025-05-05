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

type handler func(http.ResponseWriter, *http.Request)

func NewServer(invHandler *invoice.Handler, usrHandler *user.Handler) (*Server, error) {
	router := mux.NewRouter()

	router.HandleFunc("/invoice/{id}", user.JwtMiddleware(invHandler.HandleGetByID)).Methods(http.MethodGet)
	router.HandleFunc("/invoice/{id}", user.JwtMiddleware(invHandler.HandleUpdate)).Methods(http.MethodPut)
	router.HandleFunc("/invoice/{id}", user.JwtMiddleware(invHandler.HandleDelete)).Methods(http.MethodDelete)
	router.HandleFunc("/invoice", user.JwtMiddleware(invHandler.HandleGetAll)).Methods(http.MethodGet)
	router.HandleFunc("/invoice", user.JwtMiddleware(invHandler.HandleStore)).Methods(http.MethodPost)

	router.HandleFunc("/user/login", usrHandler.HandleLogin).Methods(http.MethodGet)
	router.HandleFunc("/user", user.JwtMiddleware(usrHandler.HandleGetUser)).Methods(http.MethodGet)

	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{"localhost:5173"})

	gob.Register(map[string]any{})

	return &Server{
		invHandler:  invHandler,
		Router:      router,
		methods:     methods,
		credentials: credentials,
		origins:     origins,
	}, nil
}

func (s *Server) Serve(port string) error {
	if err := http.ListenAndServe(port, handlers.CORS(s.credentials, s.methods, s.origins)(s.Router)); err != nil {
		return err
	}
	return nil
}
