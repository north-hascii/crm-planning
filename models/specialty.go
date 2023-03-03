package models

type Specialty struct {
	Id            int    `json:"-" db:"id"`
	SpecialtyName string `json:"specialty_name" db:"specialty_name" binding:"required"`
}
