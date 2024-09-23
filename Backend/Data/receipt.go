package data

type Receipt struct {
	Id               uint32
	OperationsAmount int32
	RechargeAmount   int32
	DiscountAmount   int32
	IvaAmount        int32
	TotalAmount      int32
	BonusConsumed    bool
}
