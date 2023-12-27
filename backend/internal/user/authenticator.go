package user

import (
	"context"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func NewAuthenticator(auth0_domain, auth0_client_id, auth0_client_secret, auth0_callback_url string) (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+auth0_domain+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     auth0_client_id,
		ClientSecret: auth0_client_secret,
		RedirectURL:  auth0_callback_url,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}
