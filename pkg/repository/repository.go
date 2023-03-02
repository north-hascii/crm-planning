package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/north-hascii/crm-planning/models"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUserId(email, password string) (int, error)
}

type User interface {
	GetUser(userId int) (models.User, error)
}

type Repository struct {
	Authorization
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
	}
}
