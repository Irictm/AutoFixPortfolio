package tariffOperations

import (
	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type TariffOperations = data.TariffOperations

type ITariffOperationsRepository interface {
	SaveTariffOperations(TariffOperations) (*TariffOperations, error)
	GetTariffOperationsById(uint32) (*TariffOperations, error)
	GetTariffOperationsCell(string, uint32) (*TariffOperations, error)
	GetAllTariffOperations() ([]TariffOperations, error)
	UpdateTariffOperations(TariffOperations) error
	DeleteTariffOperationsById(uint32) error
}

type Service struct {
	Repository ITariffOperationsRepository
}

func (serv *Service) SaveTariffOperations(t TariffOperations) (*TariffOperations, error) {
	return serv.Repository.SaveTariffOperations(t)
}

func (serv *Service) GetTariffOperationsById(id uint32) (*TariffOperations, error) {
	return serv.Repository.GetTariffOperationsById(id)
}

func (serv *Service) GetTariffOperationsCell(motorType string, id_operation_type uint32) (int32, error) {
	t, err := serv.Repository.GetTariffOperationsCell(motorType, id_operation_type)
	if err != nil {
		return 0, err
	}
	return t.Value, err
}

func (serv *Service) GetAllTariffOperations() ([]TariffOperations, error) {
	return serv.Repository.GetAllTariffOperations()
}

func (serv *Service) UpdateTariffOperations(t TariffOperations) error {
	return serv.Repository.UpdateTariffOperations(t)
}

func (serv *Service) DeleteTariffOperationsById(id uint32) error {
	return serv.Repository.DeleteTariffOperationsById(id)
}
