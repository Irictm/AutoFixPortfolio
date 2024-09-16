package bonus

type IBonusRepository interface {
	SaveBonus(Bonus) (*Bonus, error)
	GetBonusById(uint32) (*Bonus, error)
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

func (serv *BonusService) GetAllBonuses() ([]Bonus, error) {
	return serv.Repository.GetAllBonuses()
}

func (serv *BonusService) UpdateBonus(b Bonus) error {
	return serv.Repository.UpdateBonus(b)
}

func (serv *BonusService) DeleteBonusById(id uint32) error {
	return serv.Repository.DeleteBonusById(id)
}
