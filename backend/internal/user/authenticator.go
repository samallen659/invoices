package user

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type Authenticator interface {
	SignUp(email string, firstName string, lastName string, password string) (error, string)
}

type CognitoAuthenticator struct {
	AppClientID   string
	cognitoClient *cip.Client
}

func NewCognitoAuthentication(cognitoClientID string) (*CognitoAuthenticator, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	return &CognitoAuthenticator{
		AppClientID:   cognitoClientID,
		cognitoClient: cip.NewFromConfig(cfg),
	}, nil
}
