package data

import "time"

type Operation struct {
	Id                uint32
	Patent            string
	Id_operation_type uint32
	Date              time.Time
	Cost              int32
	Id_repair         uint32
}
