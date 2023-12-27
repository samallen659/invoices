package user

import (
	"encoding/json"
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

func (h *Handler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	var signUpReq SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&signUpReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.svc.SignUp(signUpReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
