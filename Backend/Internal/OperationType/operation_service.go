package operationType

import (
	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type OperationType = data.OperationType

type IOperationTypeRepository interface {
	SaveOperationType(OperationType) (*OperationType, error)
	GetOperationTypeById(int64) (*OperationType, error)
	GetAllOperationTypes() ([]OperationType, error)
	UpdateOperationType(OperationType) error
	DeleteOperationTypeById(int64) error
}

type Service struct {
	Repository IOperationTypeRepository
}

func (serv *Service) SaveOperationType(opType OperationType) (*OperationType, error) {
	return serv.Repository.SaveOperationType(opType)
}

func (serv *Service) GetOperationTypeById(id int64) (*OperationType, error) {
	return serv.Repository.GetOperationTypeById(id)
}

func (serv *Service) GetAllOperationTypes() ([]OperationType, error) {
	return serv.Repository.GetAllOperationTypes()
}

func (serv *Service) UpdateOperationType(opType OperationType) error {
	return serv.Repository.UpdateOperationType(opType)
}

func (serv *Service) DeleteOperationTypeById(id int64) error {
	return serv.Repository.DeleteOperationTypeById(id)
}
