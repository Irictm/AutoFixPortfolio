package repair

type Repair struct {
	Id               uint32
	DateOfAdmission  string
	DateOfRelease    string
	DateOfPickUp     string
	OperationsAmount int32
	RechargeAmount   int32
	DiscountAmount   int32
	IvaAmount        int32
	TotalAmount      int32
	Id_vehicle       uint32
}
