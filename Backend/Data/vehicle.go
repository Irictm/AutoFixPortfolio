package data

import "time"

type Vehicle struct {
	Id              uint32
	Patent          string
	Brand           string
	Model           string
	Type            string
	FabricationDate time.Time
	MotorType       string
	Seats           uint8
	Mileage         int32
}
