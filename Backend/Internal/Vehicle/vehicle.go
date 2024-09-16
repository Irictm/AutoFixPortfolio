package vehicle

import "time"

type Vehicle struct {
	Id              uint32
	Patent          string
	Brand           string
	Model           string
	VehicleType     string
	FabricationDate time.Time
	MotorType       string
	Seats           uint8
	Mileage         float32
}
