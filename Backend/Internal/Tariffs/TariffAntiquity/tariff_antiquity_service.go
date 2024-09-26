package tariffAntiquity

import (
	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type TariffAntiquity = data.TariffAntiquity

type ITariffAntiquityRepository interface {
	SaveTariffAntiquity(TariffAntiquity) (*TariffAntiquity, error)
	GetTariffAntiquityById(uint32) (*TariffAntiquity, error)
	GetTariffAntiquityCell(string, int32) (*TariffAntiquity, error)
	GetAllTariffAntiquity() ([]TariffAntiquity, error)
	UpdateTariffAntiquity(TariffAntiquity) error
	DeleteTariffAntiquityById(uint32) error
}

type Service struct {
	Repository ITariffAntiquityRepository
}

func (serv *Service) SaveTariffAntiquity(t TariffAntiquity) (*TariffAntiquity, error) {
	return serv.Repository.SaveTariffAntiquity(t)
}

func (serv *Service) GetTariffAntiquityById(id uint32) (*TariffAntiquity, error) {
	return serv.Repository.GetTariffAntiquityById(id)
}

func (serv *Service) GetTariffAntiquityCell(vehicleType string, antiquity int32) (float64, error) {
	t, err := serv.Repository.GetTariffAntiquityCell(vehicleType, antiquity)
	if err != nil {
		return 0, err
	}
	return t.Value, err
}

func (serv *Service) GetAllTariffAntiquity() ([]TariffAntiquity, error) {
	return serv.Repository.GetAllTariffAntiquity()
}

func (serv *Service) UpdateTariffAntiquity(t TariffAntiquity) error {
	return serv.Repository.UpdateTariffAntiquity(t)
}

func (serv *Service) DeleteTariffAntiquityById(id uint32) error {
	return serv.Repository.DeleteTariffAntiquityById(id)
}
