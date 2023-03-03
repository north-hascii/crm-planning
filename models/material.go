package models

type Material struct {
	Id           int    `json:"-" db:"id"`
	MaterialName string `json:"material_name" db:"material_name" binding:"required"`
}
