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

type VehicleService struct {
	Repository IVehicleRepository
}

func (serv *VehicleService) SaveVehicle(v Vehicle) (*Vehicle, error) {
	return serv.Repository.SaveVehicle(v)
}

func (serv *VehicleService) GetVehicleById(id uint32) (*Vehicle, error) {
	return serv.Repository.GetVehicleById(id)
}

func (serv *VehicleService) GetAllVehicles() ([]Vehicle, error) {
	return serv.Repository.GetAllVehicles()
}

func (serv *VehicleService) UpdateVehicle(v Vehicle) error {
	return serv.Repository.UpdateVehicle(v)
}

func (serv *VehicleService) DeleteVehicleById(id uint32) error {
	return serv.Repository.DeleteVehicleById(id)
}
