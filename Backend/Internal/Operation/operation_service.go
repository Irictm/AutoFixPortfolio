package operation

import (
	"fmt"

	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type Operation = data.Operation

type IOperationRepository interface {
	SaveOperation(Operation) (*Operation, error)
	GetOperationById(uint32) (*Operation, error)
	GetOperationVehicleMotorType(Operation) (string, error)
	GetAllOperations() ([]Operation, error)
	GetAllOperationsByRepair(uint32) ([]Operation, error)
	UpdateOperation(Operation) error
	DeleteOperationById(uint32) error
}

type ITariffService interface {
	GetOperationTariffValue(string, uint32) (int32, error)
}

type Service struct {
	Repository IOperationRepository
	TarService ITariffService
}

func (serv *Service) SaveOperation(op Operation) (*Operation, error) {
	motorType, err := serv.Repository.GetOperationVehicleMotorType(op)
	if err != nil {
		return nil, err
	}

	value, err := serv.calculateBaseCost(op, motorType)
	if err != nil {
		err := fmt.Errorf("failed saving operation, could not get cost from tariff: - %w", err)
		return nil, err
	}
	op.Cost = value
	return serv.Repository.SaveOperation(op)
}

func (serv *Service) GetOperationById(id uint32) (*Operation, error) {
	return serv.Repository.GetOperationById(id)
}

func (serv *Service) GetAllOperations() ([]Operation, error) {
	return serv.Repository.GetAllOperations()
}

func (serv *Service) GetAllOperationsByRepair(id_repair uint32) ([]Operation, error) {
	return serv.Repository.GetAllOperationsByRepair(id_repair)
}

func (serv *Service) UpdateOperation(op Operation) error {
	return serv.Repository.UpdateOperation(op)
}

func (serv *Service) DeleteOperationById(id uint32) error {
	return serv.Repository.DeleteOperationById(id)
}

func (serv *Service) calculateBaseCost(op Operation, typeOfMotor string) (int32, error) {
	cost, err := serv.TarService.GetOperationTariffValue(typeOfMotor, op.Id_operation_type)
	if err != nil {
		return 0, err
	}
	return int32(cost), nil
}

func (serv *Service) CalculateTotalBaseCost(id_repair uint32, typeOfMotor string) (int32, error) {
	var totalCost int32 = 0

	operations, err := serv.Repository.GetAllOperationsByRepair(id_repair)
	if err != nil {
		return 0, err
	}

	for _, op := range operations {
		cost, err := serv.TarService.GetOperationTariffValue(typeOfMotor, op.Id_operation_type)
		if err != nil {
			return 0, err
		}
		totalCost += cost
	}
	return totalCost, nil
}
