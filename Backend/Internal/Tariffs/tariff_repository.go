package tariffs

import (
	"context"
	"fmt"

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
		err := fmt.Errorf("error, incorrect tariff table name")
		err = fmt.Errorf("failed QUERY, could not save tariff: - %w", err)
		return nil, err
	}

	err := row.Scan(&t.Id, &t.Attribute, &t.Criteria, &t.Value)
	if err != nil {
		err = fmt.Errorf("failed Scan, tariff saved but not returned: - %w", err)
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
		err := fmt.Errorf("error, incorrect tariff table name")
		err = fmt.Errorf("failed QUERY, could not get tariff: - %w", err)
		return nil, err
	}
	err := row.Scan(&t.Id, &t.Attribute, &t.Criteria, &t.Value)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get tariff with Id %d: - %w", id, err)
		return nil, err
	}
	return &t, nil
}

func (repo *TariffRepository) GetOperationTariff(motorType string, opType string) (*Tariff, error) {
	var row pgx.Row
	var t Tariff
	row = repo.DB.QueryRow(context.Background(),
		"SELECT * FROM tariff_operations WHERE motor_type = $1 AND operation_type = $2", motorType, opType)
	err := row.Scan(&t.Id, &t.Attribute, &t.Criteria, &t.Value)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get operation tariff: - %w", err)
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
		err := fmt.Errorf("error, incorrect tariff table name")
		err = fmt.Errorf("failed QUERY, could not get all tariffs: - %w", err)
		return nil, err
	}

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all Tariffs: - %w", err)
		return nil, err
	}

	tariffs, err := pgx.CollectRows(rows, pgx.RowToStructByName[Tariff])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return tariffs, nil
}

func (repo *TariffRepository) GetAllTariffsByAttribute(attribute string, table string) ([]Tariff, error) {
	var rows pgx.Rows
	var err error
	switch table {
	case "tariff_antiquety":
		rows, err = repo.DB.Query(context.Background(),
			"SELECT * FROM tariff_antiquety WHERE vehicle_type = $1", attribute)

	case "tariff_mileage":
		rows, err = repo.DB.Query(context.Background(),
			"SELECT * FROM tariff_mileage WHERE vehicle_type = $1", attribute)

	case "tariff_operations":
		rows, err = repo.DB.Query(context.Background(),
			"SELECT * FROM tariff_operations WHERE motor_type = $1", attribute)

	case "tariff_repair_number":
		rows, err = repo.DB.Query(context.Background(),
			"SELECT * FROM tariff_repair_number WHERE motor_type = $1", attribute)

	default:
		err := fmt.Errorf("error, incorrect tariff table name")
		err = fmt.Errorf("failed QUERY, could not get all tariffs: - %w", err)
		return nil, err
	}

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all Tariffs: - %w", err)
		return nil, err
	}

	tariffs, err := pgx.CollectRows(rows, pgx.RowToStructByName[Tariff])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return tariffs, nil
}

func (repo *TariffRepository) UpdateTariff(t Tariff, table string) error {
	var err error
	switch table {
	case "tariff_antiquety":
		_, err = repo.DB.Exec(context.Background(), "UPDATE tariff_antiquety "+
			"SET vehicle_type = $2, antiquety_interval = $3, value = $4 "+
			"WHERE id = $1",
			t.Id, t.Attribute, t.Criteria, t.Value)

	case "tariff_mileage":
		_, err = repo.DB.Exec(context.Background(), "UPDATE tariff_mileage "+
			"SET vehicle_type = $2, mileage_interval = $3, value = $4 "+
			"WHERE id = $1",
			t.Id, t.Attribute, t.Criteria, t.Value)

	case "tariff_operations":
		_, err = repo.DB.Exec(context.Background(), "UPDATE tariff_operations "+
			"SET motor_type = $2, operation_type = $3, value = $4 "+
			"WHERE id = $1",
			t.Id, t.Attribute, t.Criteria, t.Value)

	case "tariff_repair_number":
		_, err = repo.DB.Exec(context.Background(), "UPDATE tariff_repair_number "+
			"SET motor_type = $2, repair_number_interval = $3, value = $4 "+
			"WHERE id = $1",
			t.Id, t.Attribute, t.Criteria, t.Value)

	default:
		err = fmt.Errorf("error, incorrect tariff table name")
		err = fmt.Errorf("failed QUERY, could not get all tariffs: - %w", err)
		return err
	}

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not update tariff: - %w", err)
		return err
	}
	return nil
}

func (repo *TariffRepository) DeleteTariffById(id uint32, table string) error {
	var err error
	switch table {
	case "tariff_antiquety":
		_, err = repo.DB.Exec(context.Background(), "DELETE tariff_antiquety "+
			"WHERE id = $1", id)

	case "tariff_mileage":
		_, err = repo.DB.Exec(context.Background(), "DELETE tariff_mileage "+
			"WHERE id = $1", id)

	case "tariff_operations":
		_, err = repo.DB.Exec(context.Background(), "DELETE tariff_operations "+
			"WHERE id = $1", id)

	case "tariff_repair_number":
		_, err = repo.DB.Exec(context.Background(), "DELETE tariff_repair_number "+
			"WHERE id = $1", id)

	default:
		err = fmt.Errorf("error, incorrect tariff table name")
		err = fmt.Errorf("failed QUERY, could not delete tariff: - %w", err)
		return err
	}

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not delete tariff: - %w", err)
		return err
	}
	return nil
}
