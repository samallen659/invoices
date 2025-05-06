package user

import (
	"encoding/json"
	"net/http"

	"github.com/samallen659/invoices/backend/internal/utils"
)

type Handler struct {
	svc *Service
}

type SignUpRequest struct {
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type DeleteUserRequest struct {
	UserName string `json:"userName"`
}

func NewHandler(svc *Service) (*Handler, error) {
	return &Handler{svc: svc}, nil
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	success, err := h.svc.LoginUser(r.Context(), loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !success {
		http.Error(w, "Login Failed", http.StatusBadRequest)
		return
	}

	jwt, err := generateJWT(loginReq.UserName)
	if err != nil {
		http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
	}

	utils.WriteJson(w, 200, `{"token":"`+jwt+`"}`)
}

func (h *Handler) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var signupReq SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&signupReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	success, err := h.svc.SignUpUser(r.Context(), signupReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if !success {
		http.Error(w, "Failed to sign up user", http.StatusInternalServerError)
	}

	jwt, err := generateJWT(signupReq.UserName)
	if err != nil {
		http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
	}

	utils.WriteJson(w, 201, `{"token":"`+jwt+`"}`)
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var deleteReq DeleteUserRequest

	err := json.NewDecoder(r.Body).Decode(&deleteReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	err = h.svc.DeleteUser(r.Context(), deleteReq.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(202)
}

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	var getUserReq GetUserRequest
	err := json.NewDecoder(r.Body).Decode(&getUserReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	user, err := h.svc.GetUser(r.Context(), getUserReq.Email, getUserReq.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	utils.WriteJson(w, http.StatusOK, user)
}
