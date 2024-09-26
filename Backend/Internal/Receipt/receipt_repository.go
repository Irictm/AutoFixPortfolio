package receipt

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) SaveReceipt(r Receipt) (*Receipt, error) {
	var receipt Receipt
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO receipts "+
		"(operations_amount, recharge_amount, discount_amount, iva_amount, total_amount, bonus_consumed) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING *",
		r.OperationsAmount, r.RechargeAmount, r.DiscountAmount, r.IvaAmount, r.TotalAmount, r.BonusConsumed).Scan(
		&receipt.Id, &receipt.OperationsAmount, &receipt.RechargeAmount, &receipt.DiscountAmount,
		&receipt.IvaAmount, &receipt.TotalAmount, &receipt.BonusConsumed)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not save receipt: - %w", err)
		return nil, err
	}
	return &receipt, nil
}

func (repo *Repository) GetReceiptById(id uint32) (*Receipt, error) {
	var receipt Receipt
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM receipts WHERE id = $1", id).Scan(
		&receipt.Id, &receipt.OperationsAmount, &receipt.RechargeAmount, &receipt.DiscountAmount,
		&receipt.IvaAmount, &receipt.TotalAmount, &receipt.BonusConsumed)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get receipt with Id %d: - %w", id, err)
		return nil, err
	}
	return &receipt, nil
}

func (repo *Repository) GetVehicleRepairNumberLastYear(id_vehicle uint32) (int32, error) {
	var repairNumberCount int32
	err := repo.DB.QueryRow(context.Background(), "SELECT COUNT(id) FROM repairs WHERE id_vehicle = $1 AND date_of_admission BETWEEN (NOW() - INTERVAL '1 year') AND NOW()",
		id_vehicle).Scan(&repairNumberCount)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get repair count last year of vehiclle with id %d: - %w", id_vehicle, err)
		return 0, err
	}
	return repairNumberCount, nil
}

func (repo *Repository) GetAllReceipts() ([]Receipt, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM receipts")
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all Receipts: - %w", err)
		return nil, err
	}

	receipts, err := pgx.CollectRows(rows, pgx.RowToStructByName[Receipt])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return receipts, nil
}

func (repo *Repository) UpdateReceipt(r Receipt) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE receipts "+
		"SET operations_amount = $2, recharge_amount = $3, discount_amount = $4, "+
		"iva_amount = $5, total_amount = $6, bonus_consumed = $7 "+
		"WHERE id = $1",
		r.Id, r.OperationsAmount, r.RechargeAmount, r.DiscountAmount, r.IvaAmount, r.TotalAmount, r.BonusConsumed)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not update receipt: - %w", err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteReceiptById(id uint32) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM receipts "+
		"WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not delete receipt: - %w", err)
		return err
	}
	return nil
}
