package vehicle

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) SaveVehicle(v Vehicle) (*Vehicle, error) {
	var vehicle Vehicle
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO vehicles "+
		"(patent, brand, model, type, fabrication_date, motor_type, seats, mileage) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *",
		v.Patent, v.Brand, v.Model, v.Type, v.FabricationDate, v.MotorType, v.Seats, v.Mileage).Scan(
		&vehicle.Id, &vehicle.Patent, &vehicle.Brand, &vehicle.Model, &vehicle.Type,
		&vehicle.FabricationDate, &vehicle.MotorType, &vehicle.Seats, &vehicle.Mileage)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not save vehicle: - %w", err)
		return nil, err
	}
	return &vehicle, nil
}

func (repo *Repository) GetVehicleById(id int64) (*Vehicle, error) {
	var vehicle Vehicle
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM vehicles WHERE id = $1", id).Scan(
		&vehicle.Id, &vehicle.Patent, &vehicle.Brand, &vehicle.Model, &vehicle.Type,
		&vehicle.FabricationDate, &vehicle.MotorType, &vehicle.Seats, &vehicle.Mileage)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get vehicle with Id %d: - %w", id, err)
		return nil, err
	}
	return &vehicle, nil
}

func (repo *Repository) GetAllVehicles() ([]Vehicle, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM vehicles")
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all Vehicles: - %w", err)
		return nil, err
	}

	vehicles, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vehicle])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return vehicles, nil
}

func (repo *Repository) UpdateVehicle(v Vehicle) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE vehicles "+
		"SET patent = $2, brand = $3, model = $4, type = $5, fabrication_date = $6, motor_type = $7, seats = $8, mileage = $9 "+
		"WHERE id = $1",
		v.Id, v.Patent, v.Brand, v.Model, v.Type, v.FabricationDate, v.MotorType, v.Seats, v.Mileage)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not update vehicle: - %w", err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteVehicleById(id int64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM vehicles "+
		"WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not delete vehicle: - %w", err)
		return err
	}
	return nil
}
