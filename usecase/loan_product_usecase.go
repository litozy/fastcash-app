package usecase

import (
	"peminjaman/model"
	"peminjaman/repo"
)

type LoanProductUsecase interface {
	GetLoanProductById(int) (*model.LoanProductModel, error)
	GetAllLoanProduct() ([]*model.LoanProductModel, error)
	InsertLoanProduct(*model.LoanProductModel) error
	DeleteLoanProduct(int) error
	UpdateLoanProduct(lprdct *model.LoanProductModel) error
}
type loanProductUsecaseImpl struct {
	lprdctRepo repo.LoanProductRepo
}

func (lprdctUsecase *loanProductUsecaseImpl) GetLoanProductById(id int) (*model.LoanProductModel, error) {
	return lprdctUsecase.lprdctRepo.GetLoanProductById(id)
}

func (lprdctUsecase *loanProductUsecaseImpl) GetAllLoanProduct() ([]*model.LoanProductModel, error) {
	return lprdctUsecase.lprdctRepo.GetAllLoanProduct()
}

func (lprdctUsecase *loanProductUsecaseImpl) InsertLoanProduct(lprdct *model.LoanProductModel) error {
	return lprdctUsecase.lprdctRepo.InsertLoanProduct(lprdct)
}

func (lprdctUsecase *loanProductUsecaseImpl) DeleteLoanProduct(id int) error {
	return lprdctUsecase.lprdctRepo.DeleteLoanProduct(id)
}

func (lprdctUsecase *loanProductUsecaseImpl) UpdateLoanProduct(lprdct *model.LoanProductModel) error {
	return lprdctUsecase.lprdctRepo.UpdateLoanProduct(lprdct)
}

func NewLoanProductUsecase(lprdctRepo repo.LoanProductRepo) LoanProductUsecase {
	return &loanProductUsecaseImpl{
		lprdctRepo: lprdctRepo,
	}
}
