package bonus

import (
	"fmt"

	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type Bonus = data.Bonus

type IBonusRepository interface {
	SaveBonus(Bonus) (*Bonus, error)
	GetBonusById(uint32) (*Bonus, error)
	GetBonusByBrand(string) (*Bonus, error)
	GetAllBonuses() ([]Bonus, error)
	UpdateBonus(Bonus) error
	DeleteBonusById(uint32) error
}

type Service struct {
	Repository IBonusRepository
}

func (serv *Service) SaveBonus(b Bonus) (*Bonus, error) {
	return serv.Repository.SaveBonus(b)
}

func (serv *Service) GetBonusById(id uint32) (*Bonus, error) {
	return serv.Repository.GetBonusById(id)
}

func (serv *Service) ConsumeBonus(brand string) (int32, error) {
	b, err := serv.Repository.GetBonusByBrand(brand)
	if err != nil {
		return 0, err
	}
	if b.Remaining <= 0 {
		err := fmt.Errorf("no bonuses remaining for brand")
		return 0, err
	}
	value := b.Amount
	b.Remaining -= 1
	serv.Repository.UpdateBonus(*b)
	return value, nil
}

func (serv *Service) CheckBonus(brand string) (int32, error) {
	b, err := serv.Repository.GetBonusByBrand(brand)
	if err != nil {
		return 0, err
	}
	return b.Amount, nil
}

func (serv *Service) GetAllBonuses() ([]Bonus, error) {
	return serv.Repository.GetAllBonuses()
}

func (serv *Service) UpdateBonus(b Bonus) error {
	return serv.Repository.UpdateBonus(b)
}

func (serv *Service) DeleteBonusById(id uint32) error {
	return serv.Repository.DeleteBonusById(id)
}
