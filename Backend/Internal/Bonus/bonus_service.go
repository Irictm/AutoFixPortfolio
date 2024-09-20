package bonus

import (
	"fmt"
	"log"
)

type IBonusRepository interface {
	SaveBonus(Bonus) (*Bonus, error)
	GetBonusById(uint32) (*Bonus, error)
	GetBonusByBrand(string) (*Bonus, error)
	GetAllBonuses() ([]Bonus, error)
	UpdateBonus(Bonus) error
	DeleteBonusById(uint32) error
}

type BonusService struct {
	Repository IBonusRepository
}

func (serv *BonusService) SaveBonus(b Bonus) (*Bonus, error) {
	return serv.Repository.SaveBonus(b)
}

func (serv *BonusService) GetBonusById(id uint32) (*Bonus, error) {
	return serv.Repository.GetBonusById(id)
}

func (serv *BonusService) ConsumeBonus(brand string) (int32, error) {
	b, err := serv.Repository.GetBonusByBrand(brand)
	if err != nil {
		log.Printf("Failed obtaining bonus - [%v]", err)
		return 0, err
	}
	if b.Remaining <= 0 {
		err := fmt.Errorf("No bonuses remaining for brand")
		return 0, err
	}
	value := b.Amount
	b.Remaining -= 1
	serv.Repository.UpdateBonus(*b)
	return value, nil
}

func (serv *BonusService) CheckBonus(brand string) (int32, error) {
	b, err := serv.Repository.GetBonusByBrand(brand)
	if err != nil {
		log.Printf("Failed obtaining bonus - [%v]", err)
		return 0, err
	}
	return b.Amount, nil
}

func (serv *BonusService) GetAllBonuses() ([]Bonus, error) {
	return serv.Repository.GetAllBonuses()
}

func (serv *BonusService) UpdateBonus(b Bonus) error {
	return serv.Repository.UpdateBonus(b)
}

func (serv *BonusService) DeleteBonusById(id uint32) error {
	return serv.Repository.DeleteBonusById(id)
}
