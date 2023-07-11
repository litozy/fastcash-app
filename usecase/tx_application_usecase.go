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
}

type transactionApplyUsecaseImpl struct {
	taRepo repo.TransactionApplyRepo
	lpRepo repo.LoanProductRepo
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

	// Melakukan penyisipan aplikasi transaksi menggunakan repositori transaksi
	err = taUsecase.taRepo.InsertApplication(tra)
	if err != nil {
		return fmt.Errorf("failed to insert application: %v", err)
	}

	return nil
}

func NewTransactionApplyUsecase(taRepo repo.TransactionApplyRepo, lpRepo repo.LoanProductRepo) TransactionApplyUsecase {
	return &transactionApplyUsecaseImpl{
		taRepo: taRepo,
		lpRepo: lpRepo,
	}
}