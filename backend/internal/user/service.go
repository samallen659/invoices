package user

type Service struct {
	repo Repository
	auth Authenticator
}

func NewService(repo Repository, auth Authenticator) (*Service, error) {
	return &Service{repo: repo, auth: auth}, nil
}

func (s *Service) SignUp() error {
	return nil
}
