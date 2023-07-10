package usecase

import (
	"peminjaman/model"
	"peminjaman/repo"
)

type InterestUsecase interface {
	GetInterestById(int) (*model.InterestModel, error)
	GetAllInterest() ([]*model.InterestModel, error)
	InsertInterest(*model.InterestModel) error
	DeleteInterest(int) error
	UpdateInterest(intrs *model.InterestModel) error
}
type interestUsecaseImpl struct {
	intrsRepo repo.InterestRepo
}

func (intrsUsecase *interestUsecaseImpl) GetInterestById(id int) (*model.InterestModel, error) {
	return intrsUsecase.intrsRepo.GetInterestById(id)
}

func (intrsUsecase *interestUsecaseImpl) GetAllInterest() ([]*model.InterestModel, error) {
	return intrsUsecase.intrsRepo.GetAllInterest()
}

func (intrsUsecase *interestUsecaseImpl) InsertInterest(intrs *model.InterestModel) error {
	return intrsUsecase.intrsRepo.InsertInterest(intrs)
}

func (intrsUsecase *interestUsecaseImpl) DeleteInterest(id int) error {
	return intrsUsecase.intrsRepo.DeleteInterest(id)
}

func (intrsUsecase *interestUsecaseImpl) UpdateInterest(intrs *model.InterestModel) error {
	return intrsUsecase.intrsRepo.UpdateInterest(intrs)
}

func NewInterestUsecase(intrsRepo repo.InterestRepo) InterestUsecase {
	return &interestUsecaseImpl{
		intrsRepo: intrsRepo,
	}
}
