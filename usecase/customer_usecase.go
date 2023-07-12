package usecase

import (
	"peminjaman/model"
	"peminjaman/repo"
)

type CustomerUsecase interface {
	GetCustomerById(int) (*model.CustomerModel, error)
	GetAllCustomer() ([]*model.CustomerModel, error)
	InsertCustomer(*model.CustomerModel) error
	DeleteCustomer(int) error
	UpdateCustomer(cstm *model.CustomerModel) error
}
type customerUsecaseImpl struct {
	cstmRepo repo.CustomerRepo
}

func (cstmUsecase *customerUsecaseImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	return cstmUsecase.cstmRepo.GetCustomerById(id)
}

func (cstmUsecase *customerUsecaseImpl) GetAllCustomer() ([]*model.CustomerModel, error) {
	return cstmUsecase.cstmRepo.GetAllCustomer()
}

func (cstmUsecase *customerUsecaseImpl) InsertCustomer(cstm *model.CustomerModel) error {
	return cstmUsecase.cstmRepo.InsertCustomer(cstm)
}

func (cstmUsecase *customerUsecaseImpl) DeleteCustomer(id int) error {
	return cstmUsecase.cstmRepo.DeleteCustomer(id)
}

func (cstmUsecase *customerUsecaseImpl) UpdateCustomer(cstm *model.CustomerModel) error {
	return cstmUsecase.cstmRepo.UpdateCustomer(cstm)
}

func NewCustomerUsecase(cstmRepo repo.CustomerRepo) CustomerUsecase {
	return &customerUsecaseImpl{
		cstmRepo: cstmRepo,
	}
}
