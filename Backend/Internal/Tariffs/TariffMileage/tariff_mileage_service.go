package tariffMileage

import (
	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type TariffMileage = data.TariffMileage

type ITariffMileageRepository interface {
	SaveTariffMileage(TariffMileage) (*TariffMileage, error)
	GetTariffMileageById(uint32) (*TariffMileage, error)
	GetTariffMileageCell(string, int32) (*TariffMileage, error)
	GetAllTariffMileage() ([]TariffMileage, error)
	UpdateTariffMileage(TariffMileage) error
	DeleteTariffMileageById(uint32) error
}

type Service struct {
	Repository ITariffMileageRepository
}

func (serv *Service) SaveTariffMileage(t TariffMileage) (*TariffMileage, error) {
	return serv.Repository.SaveTariffMileage(t)
}

func (serv *Service) GetTariffMileageById(id uint32) (*TariffMileage, error) {
	return serv.Repository.GetTariffMileageById(id)
}

func (serv *Service) GetTariffMileageCell(vehicleType string, mileage int32) (float64, error) {
	t, err := serv.Repository.GetTariffMileageCell(vehicleType, mileage)
	if err != nil {
		return 0, err
	}
	return t.Value, err
}

func (serv *Service) GetAllTariffMileage() ([]TariffMileage, error) {
	return serv.Repository.GetAllTariffMileage()
}

func (serv *Service) UpdateTariffMileage(t TariffMileage) error {
	return serv.Repository.UpdateTariffMileage(t)
}

func (serv *Service) DeleteTariffMileageById(id uint32) error {
	return serv.Repository.DeleteTariffMileageById(id)
}
