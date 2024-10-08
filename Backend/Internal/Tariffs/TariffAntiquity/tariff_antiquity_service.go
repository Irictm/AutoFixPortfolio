package tariffAntiquity

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type TariffAntiquity = data.TariffAntiquity

type ITariffAntiquityRepository interface {
	SaveTariffAntiquity(TariffAntiquity) (*TariffAntiquity, error)
	GetTariffAntiquityById(int64) (*TariffAntiquity, error)
	GetTariffAntiquityCell(string, int32) (*TariffAntiquity, error)
	GetAllTariffAntiquity() ([]TariffAntiquity, error)
	UpdateTariffAntiquity(TariffAntiquity) error
	DeleteTariffAntiquityById(int64) error
}

type Service struct {
	Repository ITariffAntiquityRepository
}

func (serv *Service) SaveTariffAntiquity(t TariffAntiquity) (*TariffAntiquity, error) {
	return serv.Repository.SaveTariffAntiquity(t)
}

func (serv *Service) SaveAndParseTariffAntiquity(data map[string]interface{}) (map[string]interface{}, error) {
	tariff, err := serv.parseTariffAntiquityFromJSON(data)
	if err != nil {
		return nil, err
	}
	tariff, err = serv.Repository.SaveTariffAntiquity(*tariff)
	if err != nil {
		return nil, err
	}
	return serv.parseJSONFromTariffAntiquity(*tariff)
}

func (serv *Service) ReceiveTariffAntiquityCSV(table [][]string) error {
	// Handle Infinity
	// Test if sending one big query is better than sending a lot of small ones (yes, it is)
	var tariff TariffAntiquity
	var interval []string
	var bottom, top float64
	var value float64
	var err error
	for row := 1; row < len(table); row++ {
		interval = strings.Split(table[row][0], " - ")
		if len(interval) != 2 {
			err = fmt.Errorf("failed processing csv, invalid interval in row %d", row)
			return err
		}
		bottom, err = strconv.ParseFloat(interval[0], 64)
		if err != nil {
			err = fmt.Errorf("failed processing csv, could not parse bottom of interval in row %d: - %w", row, err)
			return err
		}
		top, err = strconv.ParseFloat(interval[1], 64)
		if err != nil {
			err = fmt.Errorf("failed processing csv, could not parse top of interval in row %d: - %w", row, err)
			return err
		}
		for col := 1; col < len(table[row]); col++ {
			value, err = strconv.ParseFloat(table[row][col], 64)
			if err != nil {
				err = fmt.Errorf("failed processing csv, could not parse value to float64 in row %d, col %d: - %w", row, col, err)
				return err
			}
			tariff.VehicleType = table[0][col]
			tariff.Bottom = bottom
			tariff.Top = top
			tariff.Value = value
			_, err = serv.SaveTariffAntiquity(tariff)
			if err != nil {
				err = fmt.Errorf("failed processing csv, could not save tariff in row %d, col %d: - %w", row, col, err)
				return err
			}
		}
	}
	return nil
}

func (serv *Service) GetTariffAntiquityById(id int64) (*TariffAntiquity, error) {
	return serv.Repository.GetTariffAntiquityById(id)
}

func (serv *Service) GetTariffAntiquityCell(vehicleType string, antiquity int32) (float64, error) {
	t, err := serv.Repository.GetTariffAntiquityCell(vehicleType, antiquity)
	if err != nil {
		return 0, err
	}
	return t.Value, err
}

func (serv *Service) GetAllTariffAntiquity() ([]TariffAntiquity, error) {
	return serv.Repository.GetAllTariffAntiquity()
}

func (serv *Service) UpdateTariffAntiquity(t TariffAntiquity) error {
	return serv.Repository.UpdateTariffAntiquity(t)
}

func (serv *Service) DeleteTariffAntiquityById(id int64) error {
	return serv.Repository.DeleteTariffAntiquityById(id)
}

func (serv *Service) parseTariffAntiquityFromJSON(data map[string]interface{}) (*TariffAntiquity, error) {
	var tariffAntiquity TariffAntiquity
	var err error

	vehicleType := data["vehicleType"].(string)
	tariffAntiquity.VehicleType = vehicleType

	value, ok := data["value"].(float64)
	if ok {
		tariffAntiquity.Value = value
	} else {
		err = fmt.Errorf("failed parsing tariff antiquity, invalid value")
		return nil, err
	}

	bottom, ok := data["bottom"].(float64)
	if !ok {
		bottomString, ok := data["bottom"].(string)
		if ok && bottomString == "Infinity" {
			tariffAntiquity.Bottom = math.Inf(1)
		} else {
			err = fmt.Errorf("failed parsing tariff antiquity, invalid bottom")
			return nil, err
		}
	} else {
		tariffAntiquity.Bottom = bottom
	}

	top, ok := data["top"].(float64)
	if !ok {
		topString, ok := data["top"].(string)
		if ok && topString == "Infinity" {
			tariffAntiquity.Top = math.Inf(1)
		} else {
			err = fmt.Errorf("failed parsing tariff antiquity, invalid top")
			return nil, err
		}
	} else {
		tariffAntiquity.Top = top
	}

	return &tariffAntiquity, nil
}

func (serv *Service) parseJSONFromTariffAntiquity(tariff TariffAntiquity) (map[string]interface{}, error) {
	jsonBody := make(map[string]interface{})
	jsonBody["id"] = tariff.Id
	jsonBody["vehicleType"] = tariff.VehicleType
	if tariff.Bottom == math.Inf(1) {
		jsonBody["bottom"] = "Infinity"
	} else {
		jsonBody["bottom"] = tariff.Bottom
	}
	if tariff.Top == math.Inf(1) {
		jsonBody["top"] = "Infinity"
	} else {
		jsonBody["top"] = tariff.Top
	}
	jsonBody["value"] = tariff.Value
	return jsonBody, nil
}
