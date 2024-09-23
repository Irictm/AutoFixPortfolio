package repair

import data "github.com/Irictm/AutoFixPortfolio/Backend/Data"

type Repair = data.Repair

type IRepairRepository interface {
	SaveRepair(Repair) (*Repair, error)
	GetRepairById(uint32) (*Repair, error)
	GetRepairByIdReceipt(uint32) (*Repair, error)
	GetAllRepairs() ([]Repair, error)
	UpdateRepair(Repair) error
	DeleteRepairById(uint32) error
}

type RepairService struct {
	Repository IRepairRepository
}

func (serv *RepairService) SaveRepair(r Repair) (*Repair, error) {
	return serv.Repository.SaveRepair(r)
}

func (serv *RepairService) GetRepairById(id uint32) (*Repair, error) {
	return serv.Repository.GetRepairById(id)
}

func (serv *RepairService) GetRepairByIdReceipt(id_receipt uint32) (*Repair, error) {
	return serv.Repository.GetRepairByIdReceipt(id_receipt)
}

func (serv *RepairService) GetAllRepairs() ([]Repair, error) {
	return serv.Repository.GetAllRepairs()
}

func (serv *RepairService) UpdateRepair(r Repair) error {
	return serv.Repository.UpdateRepair(r)
}

func (serv *RepairService) DeleteRepairById(id uint32) error {
	return serv.Repository.DeleteRepairById(id)
}
