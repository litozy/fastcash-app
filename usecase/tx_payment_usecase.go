package usecase

import (
	"fmt"
	"peminjaman/apperror"
	"peminjaman/model"
	"peminjaman/repo"
)

type TransactionPaymentUsecase interface {
	InsertPayment(*model.TransactionPayment) error
	GetTransactionPaymentViewById(int) (*model.TransactionPaymentView, error)
}

type transactionPaymentUsecaseImpl struct {
	trpRepo repo.TransactionPaymentRepo
	taRepo repo.TransactionApplyRepo
}

func (trpUsecase *transactionPaymentUsecaseImpl) GetTransactionPaymentViewById(id int) (*model.TransactionPaymentView, error) {
	trxDB, err := trpUsecase.trpRepo.GetPaymentViewById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction id: %v", err)
	}
	if trxDB == nil {
		return nil, apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("data transaksi dengan id %v tidak ada", id),
		}
	}
	return trpUsecase.trpRepo.GetPaymentViewById(id)
}

func (trpUsecase *transactionPaymentUsecaseImpl) InsertPayment(trp *model.TransactionPayment) error {
	apply, err := trpUsecase.taRepo.GetAppById(trp.ApplicationId)
	if err != nil {
		return fmt.Errorf("failed to get transaction apply id: %v", err)
	}
	if apply == nil {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("data transaksi application dengan id %v tidak ada", trp.ApplicationId),
		}
	}
	if apply.StatusOjk == 0 {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("status anda masih %v, yang berarti masih pending, harap tunggu", apply.StatusOjk),
		}
	}
	if apply.StatusOjk == 2 {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("status anda %v, yang berarti ditolak mohon hubungi admin", apply.StatusOjk),
		}
	}

	valid, err := trpUsecase.trpRepo.GetPaymentValidateById(trp.ApplicationId)
	if err != nil {
		return fmt.Errorf("failed to get transaction id: %v", err)
	}
	if trp.Payment != valid.NeedToPayThisMonth {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: fmt.Sprintf("pembayaran harus sesuai dengan yang tertera, yaitu: %v", valid.NeedToPayThisMonth),
		}
	}
	if valid.Paid == valid.Amount {
		return apperror.AppError{
			ErrorCode: 1,
			ErrorMessage: "pembayaran sudah selesai",
		}
	}

	err = trpUsecase.trpRepo.InsertPayment(trp)
	if err != nil {
		return fmt.Errorf("failed to insert payment: %v", err)
	}

	return nil
}

func NewTransactionPaymentUsecase(trpRepo repo.TransactionPaymentRepo, taRepo repo.TransactionApplyRepo) TransactionPaymentUsecase {
	return &transactionPaymentUsecaseImpl{
		trpRepo: trpRepo,
		taRepo: taRepo,
	}
} 