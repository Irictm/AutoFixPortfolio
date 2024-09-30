package repair

import data "github.com/Irictm/AutoFixPortfolio/Backend/Data"

type Repair = data.Repair

type IRepairRepository interface {
	SaveRepair(Repair) (*Repair, error)
	GetRepairById(int64) (*Repair, error)
	GetRepairByIdReceipt(int64) (*Repair, error)
	GetAllRepairs() ([]Repair, error)
	UpdateRepair(Repair) error
	DeleteRepairById(int64) error
}

type Service struct {
	Repository IRepairRepository
}

func (serv *Service) SaveRepair(r Repair) (*Repair, error) {
	return serv.Repository.SaveRepair(r)
}

func (serv *Service) GetRepairById(id int64) (*Repair, error) {
	return serv.Repository.GetRepairById(id)
}

func (serv *Service) GetRepairByIdReceipt(id_receipt int64) (*Repair, error) {
	return serv.Repository.GetRepairByIdReceipt(id_receipt)
}

func (serv *Service) GetAllRepairs() ([]Repair, error) {
	return serv.Repository.GetAllRepairs()
}

func (serv *Service) UpdateRepair(r Repair) error {
	return serv.Repository.UpdateRepair(r)
}

func (serv *Service) DeleteRepairById(id int64) error {
	return serv.Repository.DeleteRepairById(id)
}
