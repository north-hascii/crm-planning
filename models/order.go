package models

type Order struct {
	Id              int    `json:"-" db:"id"`
	OrderTitle      string `json:"order_title" db:"order_title" binding:"required"`
	Status          string `json:"status" db:"status" binding:"required"`
	CustomerName    string `json:"customer_name" db:"customer_name" binding:"required"`
	EndDate         string `json:"end_date" db:"end_date" binding:"required"`
	PhoneNumber     string `json:"phone_number" db:"phone_number" binding:"required"`
	ManagerId       int    `json:"manager_id" db:"manager_id" binding:"required"`
	OperationListId int    `json:"operation_list_id" db:"operation_list_id" binding:"required"`
	// ФПП
}
