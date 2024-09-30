package operationType

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) SaveOperationType(opT OperationType) (*OperationType, error) {
	var opType OperationType
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO operation_types "+
		"(name) "+
		"VALUES ($1) RETURNING *",
		opT.Name).Scan(
		&opType.Id, &opType.Name)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not save operationType: - %w", err)
		return nil, err
	}
	return &opType, nil
}

func (repo *Repository) GetOperationTypeById(id int64) (*OperationType, error) {
	var opType OperationType
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM operation_types WHERE id = $1", id).Scan(
		&opType.Id, &opType.Name)
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get operationType with Id %d: - %w", id, err)
		return nil, err
	}
	return &opType, nil
}

func (repo *Repository) GetAllOperationTypes() ([]OperationType, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM operation_types")
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all OperationTypes: - %w", err)
		return nil, err
	}

	operationTypes, err := pgx.CollectRows(rows, pgx.RowToStructByName[OperationType])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return operationTypes, nil
}

func (repo *Repository) UpdateOperationType(opType OperationType) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE operation_types "+
		"SET name = $2"+
		"WHERE id = $1",
		opType.Id, opType.Name)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not update operationType: - %w", err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteOperationTypeById(id int64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM operation_types "+
		"WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not delete operationType: - %w", err)
		return err
	}
	return nil
}
