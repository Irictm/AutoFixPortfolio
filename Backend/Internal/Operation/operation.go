package operation

import "time"

type Operation struct {
	Id        uint32
	Patent    string
	Type      string
	Date      time.Time
	Cost      int32
	Id_repair uint32
}
