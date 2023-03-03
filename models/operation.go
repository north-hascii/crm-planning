package models

type Operation struct {
	Id                int    `json:"-" db:"id"`
	OperationName     string `json:"operation_name" db:"operation_name" binding:"required"`
	MaterialsListId   int    `json:"materials_list_id" db:"materials_list_id" binding:"required"`
	SpecialtiesListId int    `json:"specialties_list_id" db:"specialties_list_id" binding:"required"`
}
