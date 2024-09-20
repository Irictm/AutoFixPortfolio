package operation

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type OperationRepository struct {
	DB *pgx.Conn
}

func (repo *OperationRepository) SaveOperation(op Operation) (*Operation, error) {
	var operation Operation
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO operations "+
		"(id, patent, type, date, cost, id_repair) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING *",
		op.Id, op.Patent, op.Type, op.Date, op.Cost, op.Id_repair).Scan(
		&operation.Id, &operation.Patent, &operation.Type,
		&operation.Date, &operation.Cost, &operation.Id_repair)

	if err != nil {
		log.Printf("Failed QUERY, could not save operation - [%v]", err)
		return nil, err
	}
	return &operation, nil
}

func (repo *OperationRepository) GetOperationById(id uint32) (*Operation, error) {
	var operation Operation
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM operations WHERE id = $1", id).Scan(
		&operation.Id, &operation.Patent, &operation.Type, &operation.Date, &operation.Cost, &operation.Id_repair)
	if err != nil {
		log.Printf("Failed QUERY, could not get operation with Id %d - [%v]", id, err)
		return nil, err
	}
	return &operation, nil
}

func (repo *OperationRepository) GetOperationVehicleMotorType(op Operation) (string, error) {
	var motorType string
	err := repo.DB.QueryRow(context.Background(), "SELECT motor_type FROM vehicles WHERE patent = $1", op.Patent).Scan(
		&motorType)
	if err != nil {
		log.Printf("Failed QUERY, could not get operation motor type - [%v]", err)
		return "", err
	}
	return motorType, nil
}

func (repo *OperationRepository) GetAllOperations() ([]Operation, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM operations")
	if err != nil {
		log.Printf("Failed QUERY, could not get all Operations - [%v]", err)
		return nil, err
	}

	operations, err := pgx.CollectRows(rows, pgx.RowToStructByName[Operation])
	if err != nil {
		log.Printf("Failed Row Collection, could not get rows or parse them - [%v]", err)
		return nil, err
	}

	return operations, nil
}

func (repo *OperationRepository) GetAllOperationsByRepair(id_repair uint32) ([]Operation, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM operations WHERE id_repair = $1", id_repair)
	if err != nil {
		log.Printf("Failed QUERY, could not get all Operations by repair - [%v]", err)
		return nil, err
	}

	operations, err := pgx.CollectRows(rows, pgx.RowToStructByName[Operation])
	if err != nil {
		log.Printf("Failed Row Collection, could not get rows or parse them - [%v]", err)
		return nil, err
	}

	return operations, nil
}

func (repo *OperationRepository) UpdateOperation(op Operation) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE operations "+
		"SET brand = $2, type = $3, date = $4, cost = $5, id_repair = $6"+
		"WHERE id = $1",
		op.Id, op.Patent, op.Type, op.Date, op.Cost, op.Id_repair)

	if err != nil {
		log.Printf("Failed QUERY, could not update operation - [%v]", err)
		return err
	}
	return nil
}

func (repo *OperationRepository) DeleteOperationById(id uint32) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM operations "+
		"WHERE id = $1", id)

	if err != nil {
		log.Printf("Failed QUERY, could not delete operation - [%v]", err)
		return err
	}
	return nil
}
