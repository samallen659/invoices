package user

import (
	"context"
	"errors"
	// "fmt"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"os"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func NewAuthenticator() (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://cognito-idp.eu-west-2.amazonaws.com/eu-west-2_HeJpvCPvH",
		// fmt.Sprintf(os.Getenv("COGNITO_DOMAIN")),
	)
	if err != nil {
		return nil, err
	}

	config := oauth2.Config{
		ClientID:     os.Getenv("COGNITO_CLIENT_ID"),
		ClientSecret: os.Getenv("COGNITO_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("COGNITO_CALLBACK_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "aws.cognito.signin.user.admin", "profile", "email", "openid"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   config,
	}, nil
}

func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
