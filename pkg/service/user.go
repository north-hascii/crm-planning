package service

import (
	"github.com/north-hascii/crm-planning/models"
	"github.com/north-hascii/crm-planning/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(userId int) (models.User, error) {
	return s.repo.GetUser(userId)
}
