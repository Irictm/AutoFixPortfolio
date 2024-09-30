package tariffs

type ITariffAntiquityService interface {
	GetTariffAntiquityCell(string, int32) (float64, error)
}

type ITariffMileageService interface {
	GetTariffMileageCell(string, int32) (float64, error)
}

type ITariffOperationsService interface {
	GetTariffOperationsCell(string, int64) (int32, error)
}

type ITariffRepairNumberService interface {
	GetTariffRepairNumberCell(string, int32) (float64, error)
}

type TariffService struct {
	TariffAntiquity    ITariffAntiquityService
	TariffMileage      ITariffMileageService
	TariffOperations   ITariffOperationsService
	TariffRepairNumber ITariffRepairNumberService
}

func (serv *TariffService) GetOperationTariffValue(motorType string, id_operation_type int64) (int32, error) {
	return serv.TariffOperations.GetTariffOperationsCell(motorType, id_operation_type)
}

func (serv *TariffService) GetValueInAntiquityInterval(vehicleType string, antiquity int32) (float64, error) {
	return serv.TariffAntiquity.GetTariffAntiquityCell(vehicleType, antiquity)
}

func (serv *TariffService) GetValueInMileageInterval(vehicleType string, mileage int32) (float64, error) {
	return serv.TariffMileage.GetTariffMileageCell(vehicleType, mileage)
}

func (serv *TariffService) GetValueInRepairNumberInterval(motorType string, repairNumber int32) (float64, error) {
	return serv.TariffRepairNumber.GetTariffRepairNumberCell(motorType, repairNumber)
}
