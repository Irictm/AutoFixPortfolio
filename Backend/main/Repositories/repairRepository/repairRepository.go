package repairRepository

import (
	"context"
	"log"

	. "github.com/Irictm/AutoFixPortfolio/Backend/main/Entities/repair"
	"github.com/jackc/pgx/v5"
)

type RepairRepository struct {
	DB *pgx.Conn
}

func (vr *RepairRepository) SaveRepair(r Repair) error {
	_, err := vr.DB.Exec(context.Background(), "INSERT INTO repairs "+
		"(date_of_admission, date_of_release, date_of_pick_up, "+
		"operations_amount, recharge_amount, discount_amount, total_amount, id_vehicle) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		r.DateOfAdmission, r.DateOfRelease, r.DateOfPickUp,
		r.OperationsAmount, r.RechargeAmount, r.DiscountAmount, r.TotalAmount, r.Id_vehicle)

	if err != nil {
		log.Printf("Failed QUERY, could not save repair: %v", err)
		return err
	}
	return nil
}

func (vr *RepairRepository) GetRepairById(id uint32) (*Repair, error) {
	var repair Repair
	err := vr.DB.QueryRow(context.Background(), "SELECT * FROM repairs WHERE id = $1", id).Scan(
		&repair.DateOfAdmission, &repair.DateOfRelease, &repair.DateOfPickUp,
		&repair.OperationsAmount, &repair.RechargeAmount, &repair.DiscountAmount,
		&repair.TotalAmount, &repair.Id_vehicle)
	if err != nil {
		log.Printf("Failed QUERY, could not get repair with Id %d: %v", id, err)
		return nil, err
	}
	return &repair, nil
}

func (vr *RepairRepository) GetAllRepairs() ([]Repair, error) {
	rows, err := vr.DB.Query(context.Background(),
		"SELECT * FROM repairs")
	if err != nil {
		log.Printf("Failed QUERY, could not get all Repairs: %v", err)
		return nil, err
	}

	repairs, err := pgx.CollectRows(rows, pgx.RowToStructByName[Repair])
	if err != nil {
		log.Printf("Failed Row Collection, could not get rows or parse them: %v", err)
		return nil, err
	}

	return repairs, nil
}

func (vr *RepairRepository) UpdateRepair(r Repair) error {
	_, err := vr.DB.Exec(context.Background(), "UPDATE repairs "+
		"SET date_of_admission = $2, date_of_release = $3, date_of_pick_up = $4, "+
		"operations_amount = $5, recharge_amount = $6, discount_amount = $7, total_amount = $8, id_vehicle = $9 "+
		"WHERE id = $1",
		r.DateOfAdmission, r.DateOfRelease, r.DateOfPickUp,
		r.OperationsAmount, r.RechargeAmount, r.DiscountAmount, r.TotalAmount, r.Id_vehicle)

	if err != nil {
		log.Printf("Failed QUERY, could not update repair: %v", err)
		return err
	}
	return nil
}

func (vr *RepairRepository) DeleteRepairById(id uint32) error {
	_, err := vr.DB.Exec(context.Background(), "DELETE FROM repairs "+
		"WHERE id = $1", id)

	if err != nil {
		log.Printf("Failed QUERY, could not update repair: %v", err)
		return err
	}
	return nil
}
