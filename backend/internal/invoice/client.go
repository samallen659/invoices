package invoice

import "errors"

type Client struct {
	ClientName  string
	ClientEmail string
}

func NewClient(name string, email string) (*Client, error) {
	if name == "" {
		return nil, errors.New("Name cannot be empty")
	}
	if email == "" {
		return nil, errors.New("Email cannot be empty")
	}

	return &Client{ClientName: name, ClientEmail: email}, nil
}
