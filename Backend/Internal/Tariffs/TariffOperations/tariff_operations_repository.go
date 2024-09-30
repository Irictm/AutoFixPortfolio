package tariffOperations

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) SaveTariffOperations(t TariffOperations) (*TariffOperations, error) {
	row := repo.DB.QueryRow(context.Background(), "INSERT INTO tariff_operations "+
		"(motor_type, id_operation_type, value) "+
		"VALUES ($1, $2, $3) RETURNING *",
		t.MotorType, t.Id_operation_type, t.Value)

	err := row.Scan(&t.Id, &t.MotorType, &t.Id_operation_type, &t.Value)
	if err != nil {
		err = fmt.Errorf("failed Scan, tariff saved but not returned: - %w", err)
		return nil, err
	}
	return &t, nil
}

func (repo *Repository) GetTariffOperationsById(id int64) (*TariffOperations, error) {
	var t TariffOperations
	row := repo.DB.QueryRow(context.Background(),
		"SELECT * FROM tariff_operations WHERE id = $1", id)
	err := row.Scan(&t.Id, &t.MotorType, &t.Id_operation_type, &t.Value)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get tariff with Id %d: - %w", id, err)
		return nil, err
	}
	return &t, nil
}

func (repo *Repository) GetTariffOperationsCell(motorType string, id_operation_type int64) (*TariffOperations, error) {
	var t TariffOperations
	row := repo.DB.QueryRow(context.Background(),
		"SELECT * FROM tariff_operations WHERE motor_type = $1 AND id_operation_type = $2", motorType, id_operation_type)
	err := row.Scan(&t.Id, &t.MotorType, &t.Id_operation_type, &t.Value)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get operations tariff: - %w", err)
		return nil, err
	}
	return &t, nil
}

func (repo *Repository) GetAllTariffOperations() ([]TariffOperations, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM tariff_operations")

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all Tariffs: - %w", err)
		return nil, err
	}

	tariffs, err := pgx.CollectRows(rows, pgx.RowToStructByName[TariffOperations])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return tariffs, nil
}

func (repo *Repository) UpdateTariffOperations(t TariffOperations) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE tariff_operations "+
		"SET motor_type = $2, id_operation_type = $3, value = $4 "+
		"WHERE id = $1",
		t.Id, t.MotorType, t.Id_operation_type, t.Value)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not update tariff: - %w", err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteTariffOperationsById(id int64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE tariff_operations "+
		"WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not delete tariff: - %w", err)
		return err
	}
	return nil
}
