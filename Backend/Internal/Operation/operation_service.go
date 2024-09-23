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
	GetOperationTariffValue(string, string) (int32, error)
}

type OperationService struct {
	Repository IOperationRepository
	TarService ITariffService
}

func (serv *OperationService) SaveOperation(op Operation) (*Operation, error) {
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

func (serv *OperationService) GetOperationById(id uint32) (*Operation, error) {
	return serv.Repository.GetOperationById(id)
}

func (serv *OperationService) GetAllOperations() ([]Operation, error) {
	return serv.Repository.GetAllOperations()
}

func (serv *OperationService) GetAllOperationsByRepair(id_repair uint32) ([]Operation, error) {
	return serv.Repository.GetAllOperationsByRepair(id_repair)
}

func (serv *OperationService) UpdateOperation(op Operation) error {
	return serv.Repository.UpdateOperation(op)
}

func (serv *OperationService) DeleteOperationById(id uint32) error {
	return serv.Repository.DeleteOperationById(id)
}

func (serv *OperationService) calculateBaseCost(op Operation, typeOfMotor string) (int32, error) {
	cost, err := serv.TarService.GetOperationTariffValue(typeOfMotor, op.Type)
	if err != nil {
		return 0, err
	}
	return int32(cost), nil
}

func (serv *OperationService) CalculateTotalBaseCost(id_repair uint32, typeOfMotor string) (int32, error) {
	var totalCost int32 = 0
	cache := make(map[string]int32)

	operations, err := serv.Repository.GetAllOperationsByRepair(id_repair)
	if err != nil {
		return 0, err
	}

	for _, op := range operations {
		if cost, keyExists := cache[op.Type]; keyExists {
			totalCost += cost
		} else {
			cost, err := serv.TarService.GetOperationTariffValue(typeOfMotor, op.Type)
			if err != nil {
				return 0, err
			}
			cache[op.Type] = cost
			totalCost += cost
		}
	}
	return totalCost, nil
}
