package user

import (
	"github.com/samallen659/invoices/backend/internal/session"
	"math/rand"
	"net/http"
	"net/url"
	"os"
)

const letterBytes = "abcdefghijklmnopqrstuvmxyzADCDEFGHIJKLMNOPQRSTUVQXYZ"

var cognitoDomain string
var frontendHost string

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
	cognitoDomain = os.Getenv("COGNITO_DOMAIN")
	frontendHost = os.Getenv("FRONTEND_HOST")
	return &Handler{svc: svc}, nil
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	state := generateRandomState(32)

	ses, err := session.Get(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ses.Values["state"] = state
	if err := ses.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, h.svc.auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func (h *Handler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	ses, err := session.Get(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if ses.Values["state"] != state {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	token, err := h.svc.auth.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	idToken, err := h.svc.auth.VerifyIDToken(r.Context(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var profile map[string]any
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ses.Values["access_token"] = token.AccessToken
	ses.Values["profile"] = profile
	if err := ses.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, frontendHost, http.StatusPermanentRedirect)
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	ses, err := session.Get(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ses.Options.MaxAge = -1
	err = ses.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl, err := url.Parse(cognitoDomain + "/logout")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logout_uri, err := url.Parse(frontendHost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := url.Values{}
	params.Add("logout_uri", logout_uri.String())
	params.Add("client_id", os.Getenv("COGNITO_CLIENT_ID"))
	logoutUrl.RawQuery = params.Encode()

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}

func generateRandomState(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}
