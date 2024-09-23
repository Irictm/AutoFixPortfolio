package receipt

import (
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
	GetValueInAntiquetyInterval(string, int32) (float64, error)
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

type ReceiptService struct {
	Repository       IReceiptRepository
	TariffService    ITariffService
	BonusService     IBonusService
	VehicleService   IVehicleService
	OperationService IOperationService
	RepairService    IRepairService
}

func (serv *ReceiptService) SaveReceipt(r Receipt) (*Receipt, error) {
	return serv.Repository.SaveReceipt(r)
}

func (serv *ReceiptService) GetReceiptById(id uint32) (*Receipt, error) {
	return serv.Repository.GetReceiptById(id)
}

func (serv *ReceiptService) CalcTotalAmount(id uint32) (*Receipt, error) {
	var err error
	var baseAmount, bonusAmount int32
	var antiqRecharge, mileageRecharge, delayRecharge float64
	var repNumberDiscount, attentionDayDiscount, ivaAmount float64
	r, err := serv.Repository.GetReceiptById(id)
	if err != nil {
		return nil, err
	}
	repair, err := serv.RepairService.GetRepairByIdReceipt(r.Id)
	if err != nil {
		return nil, err
	}
	vehicle, err := serv.VehicleService.GetVehicleById(repair.Id_vehicle)
	if err != nil {
		return nil, err
	}
	baseAmount, err = serv.OperationService.CalculateTotalBaseCost(repair.Id, vehicle.MotorType)
	if err != nil {
		return nil, err
	}
	antiqRecharge, err = serv.CalcAntiquetyRecharge(vehicle.Type, vehicle.FabricationDate)
	if err != nil {
		return nil, err
	}
	mileageRecharge, err = serv.CalcMileageRecharge(vehicle.Type, vehicle.Mileage)
	if err != nil {
		return nil, err
	}
	delayRecharge = serv.CalcDelayRecharge(repair.DateOfPickUp, repair.DateOfRelease)
	repNumberDiscount, err = serv.CalcRepairNumberDiscount(vehicle.MotorType, vehicle.Id)
	if err != nil {
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
		bonusAmount = 0
	}

	antiqRecharge = math.Round(antiqRecharge * float64(baseAmount))
	mileageRecharge = math.Round(mileageRecharge * float64(baseAmount))
	delayRecharge = math.Round(delayRecharge * float64(baseAmount))
	repNumberDiscount = math.Round(repNumberDiscount * float64(baseAmount))
	attentionDayDiscount = math.Round(attentionDayDiscount * float64(baseAmount))
	ivaAmount = math.Round(float64(r.OperationsAmount-r.DiscountAmount+r.RechargeAmount) * IVA)

	// Calculo final y asignacion de valores
	r.OperationsAmount = baseAmount
	r.RechargeAmount = int32(antiqRecharge) + int32(mileageRecharge) + int32(delayRecharge)
	r.DiscountAmount = int32(repNumberDiscount) + int32(attentionDayDiscount) + bonusAmount
	r.IvaAmount = int32(ivaAmount)
	r.TotalAmount = r.OperationsAmount - r.DiscountAmount + r.RechargeAmount + r.IvaAmount
	serv.UpdateReceipt(*r)
	return r, nil
}

func (serv *ReceiptService) CalcAntiquetyRecharge(vehicleType string, fabricationDate time.Time) (float64, error) {
	currentYear := time.Now().UTC().Year()
	fabricationYear := fabricationDate.UTC().Year()
	antiquety := int32(currentYear - fabricationYear)
	return serv.TariffService.GetValueInAntiquetyInterval(vehicleType, antiquety)
}

func (serv *ReceiptService) CalcMileageRecharge(vehicleType string, mileage int32) (float64, error) {
	return serv.TariffService.GetValueInMileageInterval(vehicleType, mileage)
}

func (serv *ReceiptService) CalcDelayRecharge(dateOfPickUp time.Time, dateOfRelease time.Time) float64 {
	dateOfPickUp = dateOfPickUp.UTC()
	dateOfRelease = dateOfRelease.UTC()
	hoursBetween := dateOfPickUp.Sub(dateOfRelease).Hours()
	return float64(int(hoursBetween/24)) * RECHARGE_PER_DELAY_DAY
}

func (serv *ReceiptService) CalcRepairNumberDiscount(motorType string, id_vehicle uint32) (float64, error) {
	repairNumberCount, err := serv.Repository.GetVehicleRepairNumberLastYear(id_vehicle)
	if err != nil {
		return 0, nil
	}
	return serv.TariffService.GetValueInRepairNumberInterval(motorType, repairNumberCount)
}

func (serv *ReceiptService) CalcAttentionDayDiscount(dateOfAdmission time.Time) float64 {
	weekday := dateOfAdmission.Weekday().String()
	if weekday == "Monday" || weekday == "Thursday" {
		hour := dateOfAdmission.Hour()
		if 9 <= hour && hour <= 12 {
			return DISCOUNT_FOR_ATTENTION_DAY
		}
	}
	return 0
}

func (serv *ReceiptService) GetAllReceipts() ([]Receipt, error) {
	return serv.Repository.GetAllReceipts()
}

func (serv *ReceiptService) UpdateReceipt(r Receipt) error {
	return serv.Repository.UpdateReceipt(r)
}

func (serv *ReceiptService) DeleteReceiptById(id uint32) error {
	return serv.Repository.DeleteReceiptById(id)
}
