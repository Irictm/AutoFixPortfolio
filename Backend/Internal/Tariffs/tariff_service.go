package tariffs

import "log"

type ITariffRepository interface {
	SaveTariff(Tariff, string) (*Tariff, error)
	GetTariffById(uint32, string) (*Tariff, error)
	GetOperationTariff(string, string) (*Tariff, error)
	GetAllTariffs(string) ([]Tariff, error)
	UpdateTariff(Tariff, string) error
	DeleteTariffById(uint32, string) error
}

type TariffService struct {
	Repository ITariffRepository
}

func (serv *TariffService) SaveTariff(t Tariff, table string) (*Tariff, error) {
	return serv.Repository.SaveTariff(t, table)
}

func (serv *TariffService) GetTariffById(id uint32, table string) (*Tariff, error) {
	return serv.Repository.GetTariffById(id, table)
}

func (serv *TariffService) GetAllTariffs(table string) ([]Tariff, error) {
	return serv.Repository.GetAllTariffs(table)
}

func (serv *TariffService) UpdateTariff(t Tariff, table string) error {
	return serv.Repository.UpdateTariff(t, table)
}

func (serv *TariffService) DeleteTariffById(id uint32, table string) error {
	return serv.Repository.DeleteTariffById(id, table)
}

func (serv *TariffService) GetOperationTariff(opType string, typeOfMotor string) (int32, error) {
	tar, err := serv.Repository.GetOperationTariff(opType, typeOfMotor)
	cost := tar.Value
	if err != nil {
		log.Printf("Failed getting operation Tariff from TariffRepo - [%v]", err)
		return 0, err
	}
	return cost, nil
}
