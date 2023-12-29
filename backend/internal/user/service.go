package user

import ()

type Service struct {
	auth *Authenticator
}

func NewService(auth *Authenticator) (*Service, error) {
	return &Service{auth: auth}, nil
}
