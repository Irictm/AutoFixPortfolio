package tariffs

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type Tariff = data.Tariff

type ITariffRepository interface {
	SaveTariff(Tariff, string) (*Tariff, error)
	GetTariffById(uint32, string) (*Tariff, error)
	GetOperationTariff(string, string) (*Tariff, error)
	GetAllTariffs(string) ([]Tariff, error)
	GetAllTariffsByAttribute(string, string) ([]Tariff, error)
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

func (serv *TariffService) GetOperationTariffValue(motorType string, opType string) (int32, error) {
	t, err := serv.Repository.GetOperationTariff(motorType, opType)
	if err != nil {
		return 0, err
	}
	return int32(t.Value), err
}

func (serv *TariffService) GetAllTariffs(table string) ([]Tariff, error) {
	return serv.Repository.GetAllTariffs(table)
}

func (serv *TariffService) GetValueInAntiquetyInterval(vehicleType string, antiquety int32) (float64, error) {
	return serv.getValueInCriteriaInterval(vehicleType, antiquety, "tariff_antiquety")
}

func (serv *TariffService) GetValueInMileageInterval(vehicleType string, mileage int32) (float64, error) {
	return serv.getValueInCriteriaInterval(vehicleType, mileage, "tariff_mileage")
}

func (serv *TariffService) GetValueInRepairNumberInterval(motorType string, repairNumber int32) (float64, error) {
	return serv.getValueInCriteriaInterval(motorType, repairNumber, "tariff_repair_number")
}

func (serv *TariffService) getValueInCriteriaInterval(attribute string, criteriaValue int32, table string) (float64, error) {
	var split_criteria []string
	var value int
	var bottom, top int32
	tariffs, err := serv.Repository.GetAllTariffsByAttribute(attribute, table)
	if err != nil {
		return 0, err
	}

	for _, tariff := range tariffs {
		split_criteria = strings.Split(tariff.Criteria, " - ")
		value, err = strconv.Atoi(split_criteria[0])
		if err != nil {
			bottom = math.MinInt32
		}
		bottom = int32(value)
		value, err = strconv.Atoi(split_criteria[1])
		if err != nil {
			top = math.MaxInt32
		}
		top = int32(value)
		if bottom <= criteriaValue && criteriaValue <= top {
			return tariff.Value, nil
		}
	}

	err = fmt.Errorf("error, value not contained in any interval")
	return 0, err
}

func (serv *TariffService) UpdateTariff(t Tariff, table string) error {
	return serv.Repository.UpdateTariff(t, table)
}

func (serv *TariffService) DeleteTariffById(id uint32, table string) error {
	return serv.Repository.DeleteTariffById(id, table)
}
