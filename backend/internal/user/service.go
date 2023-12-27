package user

import "fmt"

type Service struct {
	auth Authenticator
}

func NewService(auth Authenticator) (*Service, error) {
	return &Service{auth: auth}, nil
}

func (s *Service) SignUp(req SignUpRequest) error {
	fmt.Println(req)
	return nil
}
