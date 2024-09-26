package receipt

import (
	"fmt"
	"math"
	"time"

	data "github.com/Irictm/AutoFixPortfolio/Backend/Data"
)

type Receipt = data.Receipt
type Vehicle = data.Vehicle
type Repair = data.Repair

const RECHARGE_PER_DELAY_DAY float64 = 0.05
const DISCOUNT_FOR_ATTENTION_DAY float64 = 0.1
const IVA float64 = 0.19

type IReceiptRepository interface {
	SaveReceipt(Receipt) (*Receipt, error)
	GetReceiptById(uint32) (*Receipt, error)
	GetVehicleRepairNumberLastYear(uint32) (int32, error)
	GetAllReceipts() ([]Receipt, error)
	UpdateReceipt(Receipt) error
	DeleteReceiptById(uint32) error
}

type ITariffService interface {
	GetValueInAntiquityInterval(string, int32) (float64, error)
	GetValueInMileageInterval(string, int32) (float64, error)
	GetValueInRepairNumberInterval(string, int32) (float64, error)
}

type IBonusService interface {
	ConsumeBonus(string) (int32, error)
	CheckBonus(string) (int32, error)
}

type IVehicleService interface {
	GetVehicleById(uint32) (*Vehicle, error)
}

type IOperationService interface {
	CalculateTotalBaseCost(uint32, string) (int32, error)
}

type IRepairService interface {
	GetRepairByIdReceipt(uint32) (*Repair, error)
}

type Service struct {
	Repository       IReceiptRepository
	TariffService    ITariffService
	BonusService     IBonusService
	VehicleService   IVehicleService
	OperationService IOperationService
	RepairService    IRepairService
}

func (serv *Service) SaveReceipt(r Receipt) (*Receipt, error) {
	return serv.Repository.SaveReceipt(r)
}

func (serv *Service) GetReceiptById(id uint32) (*Receipt, error) {
	return serv.Repository.GetReceiptById(id)
}

func (serv *Service) CalcTotalAmount(id uint32) (*Receipt, error) {
	var err error
	var baseAmount, bonusAmount int32
	var antiqRecharge, mileageRecharge, delayRecharge float64
	var repNumberDiscount, attentionDayDiscount, ivaAmount float64
	r, err := serv.Repository.GetReceiptById(id)
	if err != nil {
		err = fmt.Errorf("failed calcTotalAmount, could not get receipt with id %v: - %w", id, err)
		return nil, err
	}
	repair, err := serv.RepairService.GetRepairByIdReceipt(r.Id)
	if err != nil {
		err = fmt.Errorf("failed calcTotalAmount, could not get repair of receipt: - %w", err)
		return nil, err
	}
	vehicle, err := serv.VehicleService.GetVehicleById(repair.Id_vehicle)
	if err != nil {
		err = fmt.Errorf("failed calcTotalAmount, could not get vehicle with id %v: - %w", id, err)
		return nil, err
	}
	baseAmount, err = serv.OperationService.CalculateTotalBaseCost(repair.Id, vehicle.MotorType)
	if err != nil {
		err = fmt.Errorf("failed calcTotalAmount, could not get total base cost: - %w", err)
		return nil, err
	}
	antiqRecharge, err = serv.CalcAntiquityRecharge(vehicle.Type, vehicle.FabricationDate)
	if err != nil {
		err = fmt.Errorf("failed calcTotalAmount, could not get antiquity recharge: - %w", err)
		return nil, err
	}
	mileageRecharge, err = serv.CalcMileageRecharge(vehicle.Type, vehicle.Mileage)
	if err != nil {
		err = fmt.Errorf("failed calcTotalAmount, could not get mileage recharge: - %w", err)
		return nil, err
	}
	delayRecharge = serv.CalcDelayRecharge(repair.DateOfPickUp, repair.DateOfRelease)
	repNumberDiscount, err = serv.CalcRepairNumberDiscount(vehicle.MotorType, vehicle.Id)
	if err != nil {
		err = fmt.Errorf("failed calcTotalAmount, could not get repair number discount: - %w", err)
		return nil, err
	}
	attentionDayDiscount = serv.CalcAttentionDayDiscount(repair.DateOfAdmission)

	if r.BonusConsumed {
		bonusAmount, err = serv.BonusService.CheckBonus(vehicle.Brand)
	} else {
		bonusAmount, err = serv.BonusService.ConsumeBonus(vehicle.Brand)
		r.BonusConsumed = true
	}
	if err != nil {
		fmt.Println("[WARNING] Could not get bonus value, setting bonus discount to 0...")
		bonusAmount = 0
	}

	antiqRecharge = math.Round(antiqRecharge * float64(baseAmount))
	mileageRecharge = math.Round(mileageRecharge * float64(baseAmount))
	delayRecharge = math.Round(delayRecharge * float64(baseAmount))
	repNumberDiscount = math.Round(repNumberDiscount * float64(baseAmount))
	attentionDayDiscount = math.Round(attentionDayDiscount * float64(baseAmount))

	// Calculo final y asignacion de valores
	r.OperationsAmount = baseAmount
	r.RechargeAmount = int32(antiqRecharge) + int32(mileageRecharge) + int32(delayRecharge)
	r.DiscountAmount = int32(repNumberDiscount) + int32(attentionDayDiscount) + bonusAmount
	ivaAmount = math.Round(float64(r.OperationsAmount-r.DiscountAmount+r.RechargeAmount) * IVA)
	r.IvaAmount = int32(ivaAmount)
	r.TotalAmount = r.OperationsAmount - r.DiscountAmount + r.RechargeAmount + r.IvaAmount
	err = serv.UpdateReceipt(*r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (serv *Service) CalcAntiquityRecharge(vehicleType string, fabricationDate time.Time) (float64, error) {
	currentYear := time.Now().UTC().Year()
	fabricationYear := fabricationDate.UTC().Year()
	antiquity := int32(currentYear - fabricationYear)
	return serv.TariffService.GetValueInAntiquityInterval(vehicleType, antiquity)
}

func (serv *Service) CalcMileageRecharge(vehicleType string, mileage int32) (float64, error) {
	return serv.TariffService.GetValueInMileageInterval(vehicleType, mileage)
}

func (serv *Service) CalcDelayRecharge(dateOfPickUp time.Time, dateOfRelease time.Time) float64 {
	dateOfPickUp = dateOfPickUp.UTC()
	dateOfRelease = dateOfRelease.UTC()
	hoursBetween := dateOfPickUp.Sub(dateOfRelease).Hours()
	return float64(int(hoursBetween/24)) * RECHARGE_PER_DELAY_DAY
}

func (serv *Service) CalcRepairNumberDiscount(motorType string, id_vehicle uint32) (float64, error) {
	repairNumberCount, err := serv.Repository.GetVehicleRepairNumberLastYear(id_vehicle)
	if err != nil {
		return 0, nil
	}
	return serv.TariffService.GetValueInRepairNumberInterval(motorType, repairNumberCount)
}

func (serv *Service) CalcAttentionDayDiscount(dateOfAdmission time.Time) float64 {
	weekday := dateOfAdmission.Weekday().String()
	if weekday == "Monday" || weekday == "Thursday" {
		hour := dateOfAdmission.Hour()
		if 9 <= hour && hour <= 12 {
			return DISCOUNT_FOR_ATTENTION_DAY
		}
	}
	return 0
}

func (serv *Service) GetAllReceipts() ([]Receipt, error) {
	return serv.Repository.GetAllReceipts()
}

func (serv *Service) UpdateReceipt(r Receipt) error {
	return serv.Repository.UpdateReceipt(r)
}

func (serv *Service) DeleteReceiptById(id uint32) error {
	return serv.Repository.DeleteReceiptById(id)
}
