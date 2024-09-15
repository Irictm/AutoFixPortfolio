package receipt

type IReceiptRepository interface {
	SaveReceipt(Receipt) error
	GetReceiptById(uint32) (*Receipt, error)
	GetAllReceipts() ([]Receipt, error)
	UpdateReceipt(Receipt) error
	DeleteReceiptById(uint32) error
}

type ReceiptService struct {
	Repository IReceiptRepository
}

func (serv *ReceiptService) SaveReceipt(r Receipt) error {
	return serv.Repository.SaveReceipt(r)
}

func (serv *ReceiptService) GetReceiptById(id uint32) (*Receipt, error) {
	return serv.Repository.GetReceiptById(id)
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
