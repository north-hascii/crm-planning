package service

import (
	"github.com/north-hascii/crm-planning/models"
	"github.com/north-hascii/crm-planning/pkg/repository"
)

type OperationService struct {
	repo repository.Operation
}

func NewOperationService(repo repository.Operation) *OperationService {
	return &OperationService{repo: repo}
}

func (s *OperationService) GetOperationById(operationId int) (models.Operation, error) {
	return s.repo.GetOperationById(operationId)
}

func (s *OperationService) CreateOperation(operation models.Operation) (int, error) {
	return s.repo.CreateOperation(operation)
}

func (s *OperationService) GetAllOperations() ([]models.Operation, error) {
	return s.repo.GetAllOperations()
}
