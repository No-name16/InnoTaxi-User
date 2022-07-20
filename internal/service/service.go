package service

type Repository interface {
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (service *Service) Register(login, password string) error {

	return nil
}
