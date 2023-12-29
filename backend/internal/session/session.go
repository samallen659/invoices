package session

import (
	gssessions "github.com/gorilla/sessions"
	"net/http"
)

var store *gssessions.CookieStore

func New(sessionSecret string) {
	store = gssessions.NewCookieStore([]byte(sessionSecret))
}

func Get(req *http.Request) (*gssessions.Session, error) {
	return store.Get(req, "session")
}

func GetNamed(req *http.Request, name string) (*gssessions.Session, error) {
	return store.Get(req, name)
}
