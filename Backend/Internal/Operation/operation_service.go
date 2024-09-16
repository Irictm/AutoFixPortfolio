package operation

type IOperationRepository interface {
	SaveOperation(Operation) (*Operation, error)
	GetOperationById(uint32) (*Operation, error)
	GetAllOperations() ([]Operation, error)
	UpdateOperation(Operation) error
	DeleteOperationById(uint32) error
}

type OperationService struct {
	Repository IOperationRepository
}

func (serv *OperationService) SaveOperation(op Operation) (*Operation, error) {
	return serv.Repository.SaveOperation(op)
}

func (serv *OperationService) GetOperationById(id uint32) (*Operation, error) {
	return serv.Repository.GetOperationById(id)
}

func (serv *OperationService) GetAllOperations() ([]Operation, error) {
	return serv.Repository.GetAllOperations()
}

func (serv *OperationService) UpdateOperation(op Operation) error {
	return serv.Repository.UpdateOperation(op)
}

func (serv *OperationService) DeleteOperationById(id uint32) error {
	return serv.Repository.DeleteOperationById(id)
}
