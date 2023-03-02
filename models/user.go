package models

type User struct {
	Id         int    `json:"-" db:"id"`
	Email      string `json:"email" db:"email" binding:"required"`
	Password   string `json:"password" db:"password_hash" binding:"required"`
	FirstName  string `json:"first_name" db:"first_name" binding:"required"`
	SecondName string `json:"second_name" db:"second_name" binding:"required"`
	ThirdName  string `json:"third_name" db:"third_name" binding:"required"`
	UserRole   string `json:"user_role" db:"user_role" binding:"required"`
	Status     string `json:"status" db:"status" binding:"required"`
}
