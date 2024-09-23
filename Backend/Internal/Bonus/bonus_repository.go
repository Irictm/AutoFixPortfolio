package bonus

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type BonusRepository struct {
	DB *pgx.Conn
}

func (repo *BonusRepository) SaveBonus(b Bonus) (*Bonus, error) {
	var bonus Bonus
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO bonuses "+
		"(brand, remaining, amount) "+
		"VALUES ($1, $2, $3) RETURNING *",
		b.Brand, b.Remaining, b.Amount).Scan(
		&bonus.Id, &bonus.Brand, &bonus.Remaining, &bonus.Amount)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not save bonus: - %w", err)
		return nil, err
	}
	return &bonus, nil
}

func (repo *BonusRepository) GetBonusById(id uint32) (*Bonus, error) {
	var bonus Bonus
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM bonuses WHERE id = $1", id).Scan(
		&bonus.Id, &bonus.Brand, &bonus.Remaining, &bonus.Amount)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get bonus with Id %d: - %w", id, err)
		return nil, err
	}
	return &bonus, nil
}

func (repo *BonusRepository) GetBonusByBrand(brand string) (*Bonus, error) {
	var bonus Bonus
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM bonuses WHERE brand = $1", brand).Scan(
		&bonus.Id, &bonus.Brand, &bonus.Remaining, &bonus.Amount)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get bonus with brand %v: - %w", brand, err)
		return nil, err
	}
	return &bonus, nil
}

func (repo *BonusRepository) GetAllBonuses() ([]Bonus, error) {
	rows, err := repo.DB.Query(context.Background(),
		"SELECT * FROM bonuses")
	if err != nil {
		err = fmt.Errorf("failed QUERY, could not get all Bonuses: - %w", err)
		return nil, err
	}

	bonuses, err := pgx.CollectRows(rows, pgx.RowToStructByName[Bonus])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, could not get rows or parse them: - %w", err)
		return nil, err
	}

	return bonuses, nil
}

func (repo *BonusRepository) UpdateBonus(b Bonus) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE bonuses "+
		"SET brand = $2, remaining = $3, amount = $4 "+
		"WHERE id = $1",
		b.Id, b.Brand, b.Remaining, b.Amount)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not update bonus: - %w", err)
		return err
	}
	return nil
}

func (repo *BonusRepository) DeleteBonusById(id uint32) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM bonuses "+
		"WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed QUERY, could not delete bonus: - %w", err)
		return err
	}
	return nil
}
