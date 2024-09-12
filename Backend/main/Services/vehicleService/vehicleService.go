package vehicleService

import (
	. "github.com/Irictm/AutoFixPortfolio/Backend/main/Entities/vehicle"
	"github.com/Irictm/AutoFixPortfolio/Backend/main/Repositories/vehicleRepository"
)

type VehicleService struct {
	Repository vehicleRepository.VehicleRepository
}

func (serv *VehicleService) SaveVehicle(v Vehicle) error {
	return serv.Repository.SaveVehicle(v)
}

func (serv *VehicleService) GetAllVehicles() ([]Vehicle, error) {
	return serv.Repository.GetAllVehicles()
}
