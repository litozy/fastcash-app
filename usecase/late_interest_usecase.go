package usecase

import (
	"peminjaman/model"
	"peminjaman/repo"
)

type LateInterestUsecase interface {
	GetLateInterestById(int) (*model.InterestModel, error)
	GetAllLateInterest() ([]*model.InterestModel, error)
	InsertLateInterest(*model.InterestModel) error
	DeleteLateInterest(int) error
	UpdateLateInterest(intrs *model.InterestModel) error
}
type lateInterestUsecaseImpl struct {
	intrsRepo repo.LateInterestRepo
}

func (intrsUsecase *lateInterestUsecaseImpl) GetLateInterestById(id int) (*model.InterestModel, error) {
	return intrsUsecase.intrsRepo.GetLateInterestById(id)
}

func (intrsUsecase *lateInterestUsecaseImpl) GetAllLateInterest() ([]*model.InterestModel, error) {
	return intrsUsecase.intrsRepo.GetAllLateInterest()
}

func (intrsUsecase *lateInterestUsecaseImpl) InsertLateInterest(intrs *model.InterestModel) error {
	return intrsUsecase.intrsRepo.InsertLateInterest(intrs)
}

func (intrsUsecase *lateInterestUsecaseImpl) DeleteLateInterest(id int) error {
	return intrsUsecase.intrsRepo.DeleteLateInterest(id)
}

func (intrsUsecase *lateInterestUsecaseImpl) UpdateLateInterest(intrs *model.InterestModel) error {
	return intrsUsecase.intrsRepo.UpdateLateInterest(intrs)
}

func NewLateInterestUsecase(intrsRepo repo.LateInterestRepo) LateInterestUsecase {
	return &lateInterestUsecaseImpl{
		intrsRepo: intrsRepo,
	}
}
