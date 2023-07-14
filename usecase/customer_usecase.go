package usecase

import (
	"fmt"
	"peminjaman/apperror"
	"peminjaman/model"
	"peminjaman/repo"
)

type CustomerUsecase interface {
	GetCustomerById(int) (*model.CustomerModel, error)
	GetAllCustomer() ([]*model.CustomerModel, error)
	InsertCustomer(*model.CustomerModel) error
	DeleteCustomer(int) error
	UpdateCustomer(*model.CustomerModel) error
	UpdateCustomerStatus(*model.CustomerModel) error
}
type customerUsecaseImpl struct {
	cstmRepo repo.CustomerRepo
	usrRepo repo.UserRepo
}

func (cstmUsecase *customerUsecaseImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	return cstmUsecase.cstmRepo.GetCustomerById(id)
}

func (cstmUsecase *customerUsecaseImpl) GetAllCustomer() ([]*model.CustomerModel, error) {
	return cstmUsecase.cstmRepo.GetAllCustomer()
}

func (cstmUsecase *customerUsecaseImpl) InsertCustomer(cstm *model.CustomerModel) error {
	usrDB, err := cstmUsecase.usrRepo.GetUserById(cstm.UserId)
	if err != nil {
		return fmt.Errorf("failed to get transaction id: %v", err)
	}
	if usrDB == nil {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("data user dengan id %v tidak ada", cstm.UserId),
		}
	}
	return cstmUsecase.cstmRepo.InsertCustomer(cstm)
}

func (cstmUsecase *customerUsecaseImpl) DeleteCustomer(id int) error {
	return cstmUsecase.cstmRepo.DeleteCustomer(id)
}

func (cstmUsecase *customerUsecaseImpl) UpdateCustomer(cstm *model.CustomerModel) error {
	return cstmUsecase.cstmRepo.UpdateCustomer(cstm)
}

func (cstmUsecase *customerUsecaseImpl) UpdateCustomerStatus(cstm *model.CustomerModel) error {
	return cstmUsecase.cstmRepo.UpdateStatusCustomer(cstm)
}

func NewCustomerUsecase(cstmRepo repo.CustomerRepo, usrRepo repo.UserRepo) CustomerUsecase {
	return &customerUsecaseImpl{
		cstmRepo: cstmRepo,
		usrRepo: usrRepo,
	}
}
