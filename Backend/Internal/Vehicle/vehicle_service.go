package vehicle

import data "github.com/Irictm/AutoFixPortfolio/Backend/Data"

type Vehicle = data.Vehicle

type IVehicleRepository interface {
	SaveVehicle(Vehicle) (*Vehicle, error)
	GetVehicleById(uint32) (*Vehicle, error)
	GetAllVehicles() ([]Vehicle, error)
	UpdateVehicle(Vehicle) error
	DeleteVehicleById(uint32) error
}

type Service struct {
	Repository IVehicleRepository
}

func (serv *Service) SaveVehicle(v Vehicle) (*Vehicle, error) {
	return serv.Repository.SaveVehicle(v)
}

func (serv *Service) GetVehicleById(id uint32) (*Vehicle, error) {
	return serv.Repository.GetVehicleById(id)
}

func (serv *Service) GetAllVehicles() ([]Vehicle, error) {
	return serv.Repository.GetAllVehicles()
}

func (serv *Service) UpdateVehicle(v Vehicle) error {
	return serv.Repository.UpdateVehicle(v)
}

func (serv *Service) DeleteVehicleById(id uint32) error {
	return serv.Repository.DeleteVehicleById(id)
}
