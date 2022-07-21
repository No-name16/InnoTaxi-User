package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/No-name16/InnoTaxi-User/internal/entity"
	"time"
)

const (
	salt = "digsfs1224iojoisj34"
)

type Repository interface {
	CreateUser(user entity.User) (int, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (service *Service) CreateUser(user entity.User) (int, error) {
	dt := time.Now()
	dtResult := dt.Format("2006-01-06 15:04:05")
	user.Password = service.GeneratePasswordHash(user.Password)
	user.UpdatedAt = dtResult
	user.CreatedAt = dtResult
	return service.repo.CreateUser(user)
}
func (service *Service) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
