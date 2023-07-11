package usecase

import (
	"peminjaman/model"
	"peminjaman/repo"
)

type TransactionApplyUsecase interface {
	InsertApplication(*model.TransactionApply) error
}

type transactionApplyUsecaseImpl struct {
	taRepo repo.TransactionApplyRepo
}

func (taUsecase *transactionApplyUsecaseImpl) InsertApplication(tra *model.TransactionApply) error {
	return taUsecase.taRepo.InsertApplication(tra)
}

func NewTransactionApplyUsecase(taRepo repo.TransactionApplyRepo) TransactionApplyUsecase {
	return &transactionApplyUsecaseImpl{
		taRepo: taRepo,
	}
}