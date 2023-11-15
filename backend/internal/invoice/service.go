package invoice

type Service struct {
	repo Repository
}

func NewService(repo Repository) (*Service, error) {
	return &Service{repo: repo}, nil
}
