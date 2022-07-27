package service

import (
	"time"

	"github.com/No-name16/InnoTaxi-User/internal/entity"
)

const (
	salt       = "digsfs1224iojoisj34"
	tokenTTL   = 12 * time.Hour
	signingKey = "siefjeoiUOJiosjdiJ#ihuh#difjsi"
)

type Repository interface {
	CreateUser(user entity.User) (int, error)
	GetUser(phoneNumber, password string) (entity.User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
