package usecase

import (
	"peminjaman/model"
	"peminjaman/repo"
)

type OjkStatusUsecase interface {
	GetOjkStatusById(int) (*model.OjkStatusModel, error)
	GetAllOjkStatus() ([]*model.OjkStatusModel, error)
	InsertOjkStatus(*model.OjkStatusModel) error
	DeleteOjkStatus(id int) error
	UpdateOjkStatus(ojkstat *model.OjkStatusModel) error
}
type ojkStatusUsecaseImpl struct {
	ojkstatRepo repo.OjkStatusRepo
}

func (ojkstatUsecase *ojkStatusUsecaseImpl) GetOjkStatusById(id int) (*model.OjkStatusModel, error) {
	return ojkstatUsecase.ojkstatRepo.GetOjkStatusById(id)
}

func (ojkstatUsecase *ojkStatusUsecaseImpl) GetAllOjkStatus() ([]*model.OjkStatusModel, error) {
	return ojkstatUsecase.ojkstatRepo.GetAllOjkStatus()
}

func (ojkstatUsecase *ojkStatusUsecaseImpl) InsertOjkStatus(ojkstat *model.OjkStatusModel) error {
	return ojkstatUsecase.ojkstatRepo.InsertOjkStatus(ojkstat)
}
func (ojkstatUsecase *ojkStatusUsecaseImpl) DeleteOjkStatus(id int) error {
	return ojkstatUsecase.ojkstatRepo.DeleteOjkStatus(id)
}
func (ojkstatUsecase *ojkStatusUsecaseImpl) UpdateOjkStatus(ojkstat *model.OjkStatusModel) error {
	return ojkstatUsecase.ojkstatRepo.UpdateOjkStatus(ojkstat)
}

func NewOjkStatusUsecase(ojkstatRepo repo.OjkStatusRepo) OjkStatusUsecase {
	return &ojkStatusUsecaseImpl{
		ojkstatRepo: ojkstatRepo,
	}
}
