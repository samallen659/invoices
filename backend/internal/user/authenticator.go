package user

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	cipTypes "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoAuthenticator struct {
	AppClientID     string
	AppClientSecret string
	cognitoClient   *cip.Client
}

func NewCognitoAuthentication(cognitoClientID string, AppClientSecret string) (*CognitoAuthenticator, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	return &CognitoAuthenticator{
		AppClientID:   cognitoClientID,
		cognitoClient: cip.NewFromConfig(cfg),
	}, nil
}

func (c *CognitoAuthenticator) SignUp(ctx context.Context, email, firstName, lastName, password string) (*cip.SignUpOutput, error) {
	secretHash := computeSecretHash(c.AppClientSecret, email, c.AppClientID)
	awsReq := &cip.SignUpInput{
		ClientId: aws.String(c.AppClientID),
		Password: aws.String(password),
		Username: aws.String(email),
		UserAttributes: []cipTypes.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
			{
				Name:  aws.String("given_name"),
				Value: aws.String(firstName),
			},
			{
				Name:  aws.String("family_name"),
				Value: aws.String(lastName),
			},
		},
		SecretHash: aws.String(secretHash),
	}

	out, err := c.cognitoClient.SignUp(ctx, awsReq)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func computeSecretHash(clientSecret string, email string, clientID string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(email + clientID))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
