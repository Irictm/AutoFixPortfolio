package tariffRepairNumber

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) SaveTariffRepairNumber(t TariffRepairNumber) (*TariffRepairNumber, error) {
	row := repo.DB.QueryRow(context.Background(), "INSERT INTO tariff_repair_number "+
		"(motor_type, bottom, top, value) "+
		"VALUES ($1, $2, $3, $4) RETURNING *",
		t.MotorType, t.Bottom, t.Top, t.Value)

	err := row.Scan(&t.Id, &t.MotorType, &t.Bottom, &t.Top, &t.Value)
	if err != nil {
		err = fmt.Errorf("failed Scan, tariff saved but not returned: - %w", err)
		return nil, err
	}
	return &t, nil
}

func (repo *Repository) GetTariffRepairNumberById(id int64) (*TariffRepairNumber, error) {
	var t TariffRepairNumber
	row := repo.DB.QueryRow(context.Background(),
		"SELECT * FROM tariff_repair_number WHERE id = $1", id)
	err := row.Scan(&t.Id, &t.MotorType, &t.Bottom, &t.Top, &t.Value)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get tariff with Id %d: - %w", id, err)
		return nil, err
	}
	return &t, nil
}

func (repo *Repository) GetTariffRepairNumberCell(motorType string, repairNumber int32) (*TariffRepairNumber, error) {
	var t TariffRepairNumber
	row := repo.DB.QueryRow(context.Background(),
		"SELECT * FROM tariff_repair_number WHERE motor_type = $1 AND bottom <= $2 AND top >= $2", motorType, repairNumber)
	err := row.Scan(&t.Id, &t.MotorType, &t.Bottom, &t.Top, &t.Value)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get repair number tariff: - %w", err)
		return nil, err
	}
	return &t, nil
}

func (repo *Repository) GetAllTariffRepairNumber() ([]TariffRepairNumber, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM tariff_repair_number")

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all Tariffs: - %w", err)
		return nil, err
	}

	tariffs, err := pgx.CollectRows(rows, pgx.RowToStructByName[TariffRepairNumber])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return tariffs, nil
}

func (repo *Repository) UpdateTariffRepairNumber(t TariffRepairNumber) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE tariff_repair_number "+
		"SET motor_type = $2, bottom = $3, top = $4, value = $5 "+
		"WHERE id = $1",
		t.Id, t.MotorType, t.Bottom, t.Top, t.Value)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not update tariff: - %w", err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteTariffRepairNumberById(id int64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE tariff_repair_number "+
		"WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not delete tariff: - %w", err)
		return err
	}
	return nil
}
