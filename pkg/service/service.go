package service

import (
	"github.com/north-hascii/crm-planning/models"
	"github.com/north-hascii/crm-planning/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) error
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUser(userId int) (models.User, error)
}

type Service struct {
	Authorization
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
	}
}
