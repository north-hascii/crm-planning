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

type Operation interface {
	GetOperationById(userId int) (models.Operation, error)
	GetAllOperations() ([]models.Operation, error)
	CreateOperation(operation models.Operation) (int, error)
}

type Repository struct {
	Authorization
	User
	Operation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Operation:     NewOperationPostgres(db),
	}
}
