package user

import (
	"fmt"
	"github.com/samallen659/invoices/backend/internal/session"
	"net/http"
	"net/url"
	"os"
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
	state, err := generateRandomState()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
		fmt.Printf("Ses State: %s Url State: %s", ses.Values["state"], state)
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

	http.Redirect(w, r, "http://localhost:5173", http.StatusPermanentRedirect)
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

	fmt.Println(ses.Options)

	logoutUrl, err := url.Parse(os.Getenv("COGNITO_DOMAIN") + "/logout")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logout_uri, err := url.Parse("http://localhost:5173")
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

func generateRandomState() (string, error) {
	// b := make([]byte, 32)
	// _, err := rand.Read(b)
	// if err != nil {
	// 	return "", err
	// }

	// state := base64.StdEncoding.EncodeToString(b)

	return "statestring", nil
}
