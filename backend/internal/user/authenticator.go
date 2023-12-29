package user

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	AUTH_URL_PATH  = "/oauth2/authorize"
	TOKEN_URL_PATH = "/oauth2/token"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type CognitoAuthenticator struct {
	ClientID     string
	ClientSecret string
	Domain       string
	CallbackURL  string
}

func NewCognitoAuthenticator() (*CognitoAuthenticator, error) {
	cognitoClientID := os.Getenv("COGNITO_CLIENT_ID")
	cognitoClientSecret := os.Getenv("COGNITO_CLIENT_SECRET")
	cognitoDomain := os.Getenv("COGNITO_DOMAIN")
	cognitoCallbackURL := os.Getenv("COGNITO_CALLBACK_URL")

	return &CognitoAuthenticator{
		ClientID:     cognitoClientID,
		ClientSecret: cognitoClientSecret,
		Domain:       cognitoDomain,
		CallbackURL:  cognitoCallbackURL,
	}, nil
}

func (c *CognitoAuthenticator) GetLoginURL() string {
	return fmt.Sprintf("%s%s?client_id=%s&scope=aws.cognito.signin.user.admin+email+profile+openid&response_type=code&redirect_uri=%s", c.Domain, AUTH_URL_PATH, c.ClientID, c.CallbackURL)
}

func (c *CognitoAuthenticator) GetAccessToken(ctx context.Context, authCode string) (string, error) {
	base64Auth := computeBase64Authorization(c.ClientID, c.ClientSecret)
	fmt.Println(base64Auth)

	client := http.Client{}
	data := url.Values{}
	data.Add("client_id", c.ClientID)
	data.Add("grant_type", "authorization_code")
	// data.Add("scope", "profile")
	data.Add("redirect_uri", c.CallbackURL)
	data.Add("code", authCode)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Domain, TOKEN_URL_PATH), strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64Auth))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	var tokenResp TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		fmt.Println("here")
		return "", err
	}

	fmt.Printf("AccessToken: %s\n", tokenResp.AccessToken)
	fmt.Printf("IDToken: %s\n", tokenResp.IDToken)
	fmt.Printf("RefreshToken: %s\n", tokenResp.RefreshToken)
	fmt.Printf("TokenType: %s\n", tokenResp.TokenType)
	fmt.Printf("ExpiresIn: %d\n", tokenResp.ExpiresIn)

	return "", nil
}

func computeSecretHash(clientSecret string, email string, clientID string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(email + clientID))

	hash := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	fmt.Println(hash)
	return hash
}

func computeBase64Authorization(clientID string, clientSecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)))
}
