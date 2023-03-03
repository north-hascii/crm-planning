package models

type Machine struct {
	Id          int    `json:"-" db:"id"`
	MachineName string `json:"machine_name" db:"machine_name" binding:"required"`
}
