package receipt

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type ReceiptRepository struct {
	DB *pgx.Conn
}

func (repo *ReceiptRepository) SaveReceipt(r Receipt) (*Receipt, error) {
	var receipt Receipt
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO receipts "+
		"(operations_amount, recharge_amount, discount_amount, iva_amount, total_amount) "+
		"VALUES ($1, $2, $3, $4, $5) RETURNING *",
		r.OperationsAmount, r.RechargeAmount, r.DiscountAmount, r.IvaAmount, r.TotalAmount).Scan(
		&receipt.Id, &receipt.OperationsAmount, &receipt.RechargeAmount, &receipt.DiscountAmount,
		&receipt.IvaAmount, &receipt.TotalAmount)

	if err != nil {
		log.Printf("Failed QUERY, could not save receipt - [%v]", err)
		return nil, err
	}
	return &receipt, nil
}

func (repo *ReceiptRepository) GetReceiptById(id uint32) (*Receipt, error) {
	var receipt Receipt
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM receipts WHERE id = $1", id).Scan(
		&receipt.Id, &receipt.OperationsAmount, &receipt.RechargeAmount, &receipt.DiscountAmount,
		&receipt.IvaAmount, &receipt.TotalAmount)
	if err != nil {
		log.Printf("Failed QUERY, could not get receipt with Id %d - [%v]", id, err)
		return nil, err
	}
	return &receipt, nil
}

func (repo *ReceiptRepository) GetAllReceipts() ([]Receipt, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM receipts")
	if err != nil {
		log.Printf("Failed QUERY, could not get all Receipts - [%v]", err)
		return nil, err
	}

	receipts, err := pgx.CollectRows(rows, pgx.RowToStructByName[Receipt])
	if err != nil {
		log.Printf("Failed Row Collection, could not get rows or parse them - [%v]", err)
		return nil, err
	}

	return receipts, nil
}

func (repo *ReceiptRepository) UpdateReceipt(r Receipt) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE receipts "+
		"SET operations_amount = $2, recharge_amount = $3, discount_amount = $4, "+
		"iva_amount = $5, total_amount = $6 "+
		"WHERE id = $1",
		r.Id, r.OperationsAmount, r.RechargeAmount, r.DiscountAmount, r.IvaAmount, r.TotalAmount)

	if err != nil {
		log.Printf("Failed QUERY, could not update receipt - [%v]", err)
		return err
	}
	return nil
}

func (repo *ReceiptRepository) DeleteReceiptById(id uint32) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM receipts "+
		"WHERE id = $1", id)

	if err != nil {
		log.Printf("Failed QUERY, could not update receipt - [%v]", err)
		return err
	}
	return nil
}
