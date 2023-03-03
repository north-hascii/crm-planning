package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/north-hascii/crm-planning/models"
	"github.com/siruspen/logrus"
)

type OperationPostgres struct {
	db *sqlx.DB
}

func NewOperationPostgres(db *sqlx.DB) *OperationPostgres {
	return &OperationPostgres{db: db}
}

func (o *OperationPostgres) GetOperationById(userId int) (models.Operation, error) {
	var operation models.Operation
	query := fmt.Sprintf("SELECT id, operation_name, materials_list_id, specialties_list_id FROM %s WHERE id = $1", operationsTable)
	err := o.db.Get(&operation, query, userId)

	logrus.Printf("Level: repos; func GetOperationById(): models.Operation=%v", operation)

	return operation, err
}

func (o *OperationPostgres) GetAllOperations() ([]models.Operation, error) {
	var operations []models.Operation
	query := fmt.Sprintf("SELECT id, operation_name, materials_list_id, specialties_list_id FROM %s", operationsTable)
	err := o.db.Select(&operations, query)

	return operations, err
}

func (o *OperationPostgres) CreateOperation(operation *models.Operation) (int, error) {
	var id int
	createOperationQuery := fmt.Sprintf("INSERT INTO %s (operation_name, materials_list_id, specialties_list_id) VALUES ($1, $2, $3) RETURNING id", operationsTable)
	row := o.db.QueryRow(createOperationQuery, operation.OperationName, operation.MaterialsListId, operation.SpecialtiesListId)

	logrus.Printf("Level: repos; func CreateOperation(): models.Operation=%v", operation)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (o *OperationPostgres) UpdateOperation(operation *models.Operation) error {
	clearOperationQuery := fmt.Sprintf(
		"UPDATE %s SET operation_name = %s, material_list_id = %d, position_list_id = %d  WHERE user_id = %d",
		operationsTable,
		operation.OperationName,
		operation.MaterialsListId,
		operation.SpecialtiesListId,
		operation.Id,
	)

	_, err := o.db.Exec(clearOperationQuery)
	return err
}
