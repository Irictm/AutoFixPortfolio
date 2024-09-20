package tariffs

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type TariffRepository struct {
	DB *pgx.Conn
}

func (repo *TariffRepository) SaveTariff(t Tariff, table string) (*Tariff, error) {
	var row pgx.Row
	switch table {
	case "tariff_antiquety":
		row = repo.DB.QueryRow(context.Background(), "INSERT INTO tariff_antiquety "+
			"(vehicle_type, antiquety_interval, value) "+
			"VALUES ($1, $2, $3) RETURNING *",
			t.Attribute, t.Criteria, t.Value)
	case "tariff_mileage":
		row = repo.DB.QueryRow(context.Background(), "INSERT INTO tariff_mileage "+
			"(vehicle_type, mileage_interval, value) "+
			"VALUES ($1, $2, $3) RETURNING *",
			t.Attribute, t.Criteria, t.Value)
	case "tariff_operations":
		row = repo.DB.QueryRow(context.Background(), "INSERT INTO tariff_operations "+
			"(motor_type, operation_type, value) "+
			"VALUES ($1, $2, $3) RETURNING *",
			t.Attribute, t.Criteria, t.Value)
	case "tariff_repair_number":
		row = repo.DB.QueryRow(context.Background(), "INSERT INTO tariff_repair_number "+
			"(motor_type, repair_number_interval, value) "+
			"VALUES ($1, $2, $3) RETURNING *",
			t.Attribute, t.Criteria, t.Value)
	default:
		err := fmt.Errorf("ERROR: Incorrect tariff table name")
		log.Printf("Failed QUERY, could not save tariff - [%v]", err)
		return nil, err
	}

	err := row.Scan(&t.Id, &t.Attribute, &t.Criteria, &t.Value)
	if err != nil {
		log.Printf("Failed Scan, tariff saved but not returned - [%v]", err)
		return nil, err
	}
	return &t, nil
}

func (repo *TariffRepository) GetTariffById(id uint32, table string) (*Tariff, error) {
	var row pgx.Row
	var t Tariff
	switch table {
	case "tariff_antiquety":
		row = repo.DB.QueryRow(context.Background(),
			"SELECT * FROM tariff_antiquety WHERE id = $1", id)

	case "tariff_mileage":
		row = repo.DB.QueryRow(context.Background(),
			"SELECT * FROM tariff_mileage WHERE id = $1", id)

	case "tariff_operations":
		row = repo.DB.QueryRow(context.Background(),
			"SELECT * FROM tariff_operations WHERE id = $1", id)

	case "tariff_repair_number":
		row = repo.DB.QueryRow(context.Background(),
			"SELECT * FROM tariff_repair_number WHERE id = $1", id)

	default:
		err := fmt.Errorf("ERROR: Incorrect tariff table name")
		log.Printf("Failed QUERY, could not get tariff - [%v]", err)
		return nil, err
	}
	err := row.Scan(&t.Id, &t.Attribute, &t.Criteria, &t.Value)
	if err != nil {
		log.Printf("Failed QUERY, could not get tariff with Id %d - [%v]", id, err)
		return nil, err
	}
	return &t, nil
}

func (repo *TariffRepository) GetOperationTariff(motorType string, opType string) (*Tariff, error) {
	var t Tariff
	row := repo.DB.QueryRow(context.Background(),
		"SELECT * FROM tariff_operations WHERE motor_type = $1, operation_type = $2", motorType, opType)
	err := row.Scan(&t.Id, &t.Attribute, &t.Criteria, &t.Value)

	if err != nil {
		log.Printf("Failed QUERY, could not get tariff - [%v]", err)
		return nil, err
	}
	return &t, nil
}

func (repo *TariffRepository) GetAllTariffs(table string) ([]Tariff, error) {
	var rows pgx.Rows
	var err error
	switch table {
	case "tariff_antiquety":
		rows, err = repo.DB.Query(context.Background(),
			"SELECT * FROM tariff_antiquety")

	case "tariff_mileage":
		rows, err = repo.DB.Query(context.Background(),
			"SELECT * FROM tariff_mileage")

	case "tariff_operations":
		rows, err = repo.DB.Query(context.Background(),
			"SELECT * FROM tariff_operations")

	case "tariff_repair_number":
		rows, err = repo.DB.Query(context.Background(),
			"SELECT * FROM tariff_repair_number")

	default:
		err := fmt.Errorf("ERROR: Incorrect tariff table name")
		log.Printf("Failed QUERY, could not get all tariffs - [%v]", err)
		return nil, err
	}

	if err != nil {
		log.Printf("Failed QUERY, could not get all Tariffs - [%v]", err)
		return nil, err
	}

	tariffs, err := pgx.CollectRows(rows, pgx.RowToStructByName[Tariff])
	if err != nil {
		log.Printf("Failed Row Collection, could not get rows or parse them - [%v]", err)
		return nil, err
	}

	return tariffs, nil
}

func (repo *TariffRepository) UpdateTariff(t Tariff, table string) error {
	var err error
	switch table {
	case "tariff_antiquety":
		_, err = repo.DB.Exec(context.Background(), "UPDATE tariff_antiquety"+
			"SET vehicle_type = $2, antiquety_interval = $3, value = $4 "+
			"WHERE id = $1",
			t.Id, t.Attribute, t.Criteria, t.Value)

	case "tariff_mileage":
		_, err = repo.DB.Exec(context.Background(), "UPDATE tariff_mileage"+
			"SET vehicle_type = $2, mileage_interval = $3, value = $4 "+
			"WHERE id = $1",
			t.Id, t.Attribute, t.Criteria, t.Value)

	case "tariff_operations":
		_, err = repo.DB.Exec(context.Background(), "UPDATE tariff_operations"+
			"SET motor_type = $2, operation_type = $3, value = $4 "+
			"WHERE id = $1",
			t.Id, t.Attribute, t.Criteria, t.Value)

	case "tariff_repair_number":
		_, err = repo.DB.Exec(context.Background(), "UPDATE tariff_repair_number"+
			"SET motor_type = $2, repair_number_interval = $3, value = $4 "+
			"WHERE id = $1",
			t.Id, t.Attribute, t.Criteria, t.Value)

	default:
		err = fmt.Errorf("ERROR: Incorrect tariff table name")
		log.Printf("Failed QUERY, could not get all tariffs - [%v]", err)
		return err
	}

	if err != nil {
		log.Printf("Failed QUERY, could not update tariff - [%v]", err)
		return err
	}
	return nil
}

func (repo *TariffRepository) DeleteTariffById(t Tariff, table string) error {
	var err error
	switch table {
	case "tariff_antiquety":
		_, err = repo.DB.Exec(context.Background(), "DELETE tariff_antiquety "+
			"WHERE id = $1", t.Id)

	case "tariff_mileage":
		_, err = repo.DB.Exec(context.Background(), "DELETE tariff_mileage "+
			"WHERE id = $1", t.Id)

	case "tariff_operations":
		_, err = repo.DB.Exec(context.Background(), "DELETE tariff_operations "+
			"WHERE id = $1", t.Id)

	case "tariff_repair_number":
		_, err = repo.DB.Exec(context.Background(), "DELETE tariff_repair_number "+
			"WHERE id = $1", t.Id)

	default:
		err = fmt.Errorf("ERROR: Incorrect tariff table name")
		log.Printf("Failed QUERY, could not delete tariff - [%v]", err)
		return err
	}

	if err != nil {
		log.Printf("Failed QUERY, could not delete tariff - [%v]", err)
		return err
	}
	return nil
}
