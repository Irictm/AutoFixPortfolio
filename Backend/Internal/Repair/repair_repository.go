package repair

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) SaveRepair(r Repair) (*Repair, error) {
	var repair Repair
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO repairs "+
		"(date_of_admission, date_of_release, date_of_pick_up, id_receipt, id_vehicle) "+
		"VALUES ($1, $2, $3, $4, $5) RETURNING *",
		r.DateOfAdmission, r.DateOfRelease, r.DateOfPickUp, r.Id_receipt, r.Id_vehicle).Scan(
		&repair.Id, &repair.DateOfAdmission, &repair.DateOfRelease, &repair.DateOfPickUp,
		&repair.Id_receipt, &repair.Id_vehicle)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not save repair: - %w", err)
		return nil, err
	}
	return &repair, nil
}

func (repo *Repository) GetRepairById(id int64) (*Repair, error) {
	var repair Repair
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM repairs WHERE id = $1", id).Scan(
		&repair.Id, &repair.DateOfAdmission, &repair.DateOfRelease, &repair.DateOfPickUp,
		&repair.Id_receipt, &repair.Id_vehicle)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get repair with Id %d: - %w", id, err)
		return nil, err
	}
	return &repair, nil
}

func (repo *Repository) GetRepairByIdReceipt(id_receipt int64) (*Repair, error) {
	var repair Repair
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM repairs WHERE id_receipt = $1", id_receipt).Scan(
		&repair.Id, &repair.DateOfAdmission, &repair.DateOfRelease, &repair.DateOfPickUp,
		&repair.Id_receipt, &repair.Id_vehicle)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get repair with Id of receipt %d: - %w", id_receipt, err)
		return nil, err
	}
	return &repair, nil
}

func (repo *Repository) GetAllRepairs() ([]Repair, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM repairs")
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all Repairs: - %w", err)
		return nil, err
	}

	repairs, err := pgx.CollectRows(rows, pgx.RowToStructByName[Repair])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return repairs, nil
}

func (repo *Repository) UpdateRepair(r Repair) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE repairs "+
		"SET date_of_admission = $2, date_of_release = $3, date_of_pick_up = $4, "+
		"id_receipt = $5, id_vehicle = $6 "+
		"WHERE id = $1",
		r.Id, r.DateOfAdmission, r.DateOfRelease, r.DateOfPickUp, r.Id_receipt, r.Id_vehicle)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not update repair: - %w", err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteRepairById(id int64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM repairs "+
		"WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not delete repair: - %w", err)
		return err
	}
	return nil
}
