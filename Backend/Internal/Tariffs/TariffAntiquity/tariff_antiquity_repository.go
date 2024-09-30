package tariffAntiquity

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) SaveTariffAntiquity(t TariffAntiquity) (*TariffAntiquity, error) {
	row := repo.DB.QueryRow(context.Background(), "INSERT INTO tariff_antiquity "+
		"(vehicle_type, bottom, top, value) "+
		"VALUES ($1, $2, $3, $4) RETURNING *",
		t.VehicleType, t.Bottom, t.Top, t.Value)

	err := row.Scan(&t.Id, &t.VehicleType, &t.Bottom, &t.Top, &t.Value)
	if err != nil {
		err = fmt.Errorf("failed Scan, tariff saved but not returned: - %w", err)
		return nil, err
	}
	return &t, nil
}

func (repo *Repository) GetTariffAntiquityById(id int64) (*TariffAntiquity, error) {
	var t TariffAntiquity
	row := repo.DB.QueryRow(context.Background(),
		"SELECT * FROM tariff_antiquity WHERE id = $1", id)
	err := row.Scan(&t.Id, &t.VehicleType, &t.Bottom, &t.Top, &t.Value)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get tariff with Id %d: - %w", id, err)
		return nil, err
	}
	return &t, nil
}

func (repo *Repository) GetTariffAntiquityCell(vehicleType string, antiquity int32) (*TariffAntiquity, error) {
	var t TariffAntiquity
	row := repo.DB.QueryRow(context.Background(),
		"SELECT * FROM tariff_antiquity WHERE vehicle_type = $1 AND bottom <= $2 AND top >= $2", vehicleType, antiquity)
	err := row.Scan(&t.Id, &t.VehicleType, &t.Bottom, &t.Top, &t.Value)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get antiquity tariff: - %w", err)
		return nil, err
	}
	return &t, nil
}

func (repo *Repository) GetAllTariffAntiquity() ([]TariffAntiquity, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM tariff_antiquity")

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all Tariffs: - %w", err)
		return nil, err
	}

	tariffs, err := pgx.CollectRows(rows, pgx.RowToStructByName[TariffAntiquity])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return tariffs, nil
}

func (repo *Repository) UpdateTariffAntiquity(t TariffAntiquity) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE tariff_antiquity "+
		"SET vehicle_type = $2, bottom = $3, top = $4, value = $5 "+
		"WHERE id = $1",
		t.Id, t.VehicleType, t.Bottom, t.Top, t.Value)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not update tariff: - %w", err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteTariffAntiquityById(id int64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE tariff_antiquity "+
		"WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not delete tariff: - %w", err)
		return err
	}
	return nil
}
