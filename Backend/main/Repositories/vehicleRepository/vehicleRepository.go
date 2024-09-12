package vehicleRepository

import (
	"database/sql"

	. "github.com/Irictm/AutoFixPortfolio/Backend/main/Entities/vehicle"
	_ "github.com/lib/pq"
)

type VehicleRepository struct {
	DB *sql.DB
}

func (vr *VehicleRepository) SaveVehicle(v Vehicle) error {
	_, err := vr.DB.Exec("INSERT INTO vehicles "+
		"(patent, brand, model, vehicle_type, fabrication_date, motor_type, seats, mileage) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		v.Patent, v.Brand, v.Model, v.VehicleType, v.FabricationDate, v.MotorType, v.Seats, v.Mileage)
	return err
}

func (vr *VehicleRepository) GetAllVehicles() ([]Vehicle, error) {
	rows, err := vr.DB.Query(
		"SELECT id, patent, brand, model, vehicle_type, fabrication_date, motor_type, seats, mileage FROM vehicles")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var vehicles []Vehicle

	for rows.Next() {
		var vehicle Vehicle

		if err := rows.Scan(&vehicle.Id, &vehicle.Patent, &vehicle.Brand, &vehicle.Model, &vehicle.VehicleType,
			&vehicle.FabricationDate, &vehicle.MotorType, &vehicle.Seats, &vehicle.Mileage); err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}
