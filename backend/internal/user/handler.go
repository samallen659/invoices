package user

import (
	"fmt"
	"github.com/samallen659/invoices/backend/internal/session"
	"net/http"
)

type Handler struct {
	svc *Service
}

type SignUpRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewHandler(svc *Service) (*Handler, error) {
	return &Handler{svc: svc}, nil
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	loginUrl := h.svc.GetLoginURL()
	http.Redirect(w, r, loginUrl, http.StatusSeeOther)
}

func (h *Handler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	session, err := session.Get(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(session)
	code := r.URL.Query().Get("code")

	token, err := h.svc.GetAccessToken(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed", http.StatusBadRequest)
	}

	fmt.Println(token)
}
