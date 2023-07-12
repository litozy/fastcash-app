package usecase

import (
	"fmt"
	"peminjaman/apperror"
	"peminjaman/model"
	"peminjaman/repo"
)

type TransactionApplyUsecase interface {
	InsertApplication(*model.TransactionApply) error
	GetAllApp() ([]model.TransactionApplyView, error)
	GetTransactionApplyById(int) (*model.TransactionApplyView, error)
	UpdateStatusOjk(*model.TransactionApply) error
}

type transactionApplyUsecaseImpl struct {
	taRepo repo.TransactionApplyRepo
	lpRepo repo.LoanProductRepo
	ojkRepo repo.OjkStatusRepo
}

func (taUsecase *transactionApplyUsecaseImpl) GetTransactionApplyById(id int) (*model.TransactionApplyView, error) {
	trxDB, err := taUsecase.taRepo.GetAppById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction id: %v", err)
	}
	if trxDB == nil {
		return nil, apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("data transaksi dengan id %v tidak ada", id),
		}
	}
	return taUsecase.taRepo.GetAppById(id)
}

func (taUsecase *transactionApplyUsecaseImpl) GetAllApp() ([]model.TransactionApplyView, error) {
	return taUsecase.taRepo.GetAllApp()
}

func (taUsecase *transactionApplyUsecaseImpl) InsertApplication(tra *model.TransactionApply) error {
	loanProduct, err := taUsecase.lpRepo.GetLoanProductById(tra.ProductId)
	if err != nil {
		return fmt.Errorf("failed to get loan product: %v", err)
	}
	if loanProduct == nil {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("data product dengan id %v tidak ada", tra.ProductId),
		}
	}
	if tra.Amount > loanProduct.MaxLoan {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("peminjaman di limit di angka %v, tidak boleh lebih", loanProduct.MaxLoan),
		}
	}
	if tra.Amount <= 0 {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: "peminjaman tidak boleh kurang dari 0",
		}
	}

	// Melakukan penyisipan aplikasi transaksi menggunakan repositori transaksi
	err = taUsecase.taRepo.InsertApplication(tra)
	if err != nil {
		return fmt.Errorf("failed to insert application: %v", err)
	}

	return nil
}

func (taUsecase *transactionApplyUsecaseImpl) UpdateStatusOjk(tra *model.TransactionApply) error {
	ojkStatus, err := taUsecase.ojkRepo.GetOjkStatusById(tra.OjkStatusId)
	if err != nil {
		return fmt.Errorf("failed to get status ojk: %v", err)
	}
	if ojkStatus == nil {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("data ojk status dengan id %v tidak ada", tra.OjkStatusId),
		}
	}
	if tra.OjkStatusId <= 0 {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: "status id tidak boleh kurang dari 0",
		}
	}
	
	
	err =  taUsecase.taRepo.UpdateStatusOjk(tra)
	if err != nil {
		return fmt.Errorf("failed to update status: %v", err)
	}

	return nil
}

func NewTransactionApplyUsecase(taRepo repo.TransactionApplyRepo, lpRepo repo.LoanProductRepo, ojkRepo repo.OjkStatusRepo) TransactionApplyUsecase {
	return &transactionApplyUsecaseImpl{
		taRepo: taRepo,
		lpRepo: lpRepo,
		ojkRepo: ojkRepo,
	}
}