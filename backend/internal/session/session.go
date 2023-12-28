package session

import (
	"net/http"
	"os"

	gssessions "github.com/gorilla/sessions"
)

var store = gssessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func Get(req *http.Request) (*gssessions.Session, error) {
	return store.Get(req, "session-name")
}

func GetNamed(req *http.Request, name string) (*gssessions.Session, error) {
	return store.Get(req, name)
}
