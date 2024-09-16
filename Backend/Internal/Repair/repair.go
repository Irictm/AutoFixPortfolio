package repair

import "time"

type Repair struct {
	Id              uint32
	DateOfAdmission time.Time
	DateOfRelease   time.Time
	DateOfPickUp    time.Time
	Id_receipt      uint32
	Id_vehicle      uint32
}
