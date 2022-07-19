package service

import "github.com/No-name16/InnoTaxi-User/InnoTaxi-User/internal/repository"

type Authorization interface {
}

type Service struct {
	repo Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
