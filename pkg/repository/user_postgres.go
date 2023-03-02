package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/north-hascii/crm-planning/models"
	"github.com/siruspen/logrus"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUser(userId int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, email, first_name, second_name, third_name, user_role, status FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, query, userId)

	logrus.Printf("Level: repos; func getUser(): models.User=%v", user)

	return user, err
}
