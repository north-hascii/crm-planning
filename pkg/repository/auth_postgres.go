package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/north-hascii/crm-planning/models"
	"github.com/siruspen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) CreateUser(user models.User) error {
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash, first_name, second_name, third_name, user_role, status) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", usersTable)
	err := p.db.QueryRow(query, user.Email, user.Password, user.FirstName, user.SecondName, user.ThirdName, user.UserRole, user.Status)

	logrus.Printf("Level: repos; func CreateUser(): user.id=%d, err=%v", user.Id, err)

	return err.Err()
}

func (p *AuthPostgres) GetUserId(email, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := p.db.Get(&id, query, email, password)

	logrus.Printf("Level: repos; func GetUserId(): user.id=%d, err=%v", id, err)

	return id, err
}
