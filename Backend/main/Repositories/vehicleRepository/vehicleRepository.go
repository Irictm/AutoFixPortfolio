package vehicleRepository

import (
	"context"
	"log"

	. "github.com/Irictm/AutoFixPortfolio/Backend/main/Entities/vehicle"
	"github.com/jackc/pgx/v5"
)

type VehicleRepository struct {
	DB *pgx.Conn
}

func (vr *VehicleRepository) SaveVehicle(v Vehicle) error {
	_, err := vr.DB.Exec(context.Background(), "INSERT INTO vehicles "+
		"(patent, brand, model, vehicle_type, fabrication_date, motor_type, seats, mileage) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		v.Patent, v.Brand, v.Model, v.VehicleType, v.FabricationDate, v.MotorType, v.Seats, v.Mileage)

	if err != nil {
		log.Printf("Failed QUERY, could not save vehicle: %v", err)
		return err
	}
	return nil
}

func (vr *VehicleRepository) GetVehicleById(id uint32) (*Vehicle, error) {
	var vehicle Vehicle
	err := vr.DB.QueryRow(context.Background(), "SELECT * FROM vehicles WHERE id = $1", id).Scan(
		&vehicle.Id, &vehicle.Patent, &vehicle.Brand, &vehicle.Model, &vehicle.VehicleType,
		&vehicle.FabricationDate, &vehicle.MotorType, &vehicle.Seats, &vehicle.Mileage)
	if err != nil {
		log.Printf("Failed QUERY, could not get vehicle with Id %d: %v", id, err)
		return nil, err
	}
	return &vehicle, nil
}

func (vr *VehicleRepository) GetAllVehicles() ([]Vehicle, error) {
	rows, err := vr.DB.Query(context.Background(),
		"SELECT * FROM vehicles")
	if err != nil {
		log.Printf("Failed QUERY, could not get all Vehicles: %v", err)
		return nil, err
	}

	vehicles, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vehicle])
	if err != nil {
		log.Printf("Failed Row Collection, could not get rows or parse them: %v", err)
		return nil, err
	}

	return vehicles, nil
}

func (vr *VehicleRepository) UpdateVehicle(v Vehicle) error {
	_, err := vr.DB.Exec(context.Background(), "UPDATE vehicles "+
		"SET patent = $2, brand = $3, model = $4, vehicle_type = $5, fabrication_date = $6, motor_type = $7, seats = $8, mileage = $9 "+
		"WHERE id = $1",
		v.Id, v.Patent, v.Brand, v.Model, v.VehicleType, v.FabricationDate, v.MotorType, v.Seats, v.Mileage)

	if err != nil {
		log.Printf("Failed QUERY, could not update vehicle: %v", err)
		return err
	}
	return nil
}

func (vr *VehicleRepository) DeleteVehicleById(id uint32) error {
	_, err := vr.DB.Exec(context.Background(), "DELETE FROM vehicles "+
		"WHERE id = $1", id)

	if err != nil {
		log.Printf("Failed QUERY, could not update vehicle: %v", err)
		return err
	}
	return nil
}
