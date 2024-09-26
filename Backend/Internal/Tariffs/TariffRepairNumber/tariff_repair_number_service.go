package tariffRepairNumber

import (
	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type TariffRepairNumber = data.TariffRepairNumber

type ITariffRepairNumberRepository interface {
	SaveTariffRepairNumber(TariffRepairNumber) (*TariffRepairNumber, error)
	GetTariffRepairNumberById(uint32) (*TariffRepairNumber, error)
	GetTariffRepairNumberCell(string, int32) (*TariffRepairNumber, error)
	GetAllTariffRepairNumber() ([]TariffRepairNumber, error)
	UpdateTariffRepairNumber(TariffRepairNumber) error
	DeleteTariffRepairNumberById(uint32) error
}

type Service struct {
	Repository ITariffRepairNumberRepository
}

func (serv *Service) SaveTariffRepairNumber(t TariffRepairNumber) (*TariffRepairNumber, error) {
	return serv.Repository.SaveTariffRepairNumber(t)
}

func (serv *Service) GetTariffRepairNumberById(id uint32) (*TariffRepairNumber, error) {
	return serv.Repository.GetTariffRepairNumberById(id)
}

func (serv *Service) GetTariffRepairNumberCell(motorType string, repairNumber int32) (float64, error) {
	t, err := serv.Repository.GetTariffRepairNumberCell(motorType, repairNumber)
	if err != nil {
		return 0, err
	}
	return t.Value, err
}

func (serv *Service) GetAllTariffRepairNumber() ([]TariffRepairNumber, error) {
	return serv.Repository.GetAllTariffRepairNumber()
}

func (serv *Service) UpdateTariffRepairNumber(t TariffRepairNumber) error {
	return serv.Repository.UpdateTariffRepairNumber(t)
}

func (serv *Service) DeleteTariffRepairNumberById(id uint32) error {
	return serv.Repository.DeleteTariffRepairNumberById(id)
}
