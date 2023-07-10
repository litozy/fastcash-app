package usecase

import (
	"fmt"
	"peminjaman/apperror"
	"peminjaman/model"
	"peminjaman/repo"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	InsertUser(*model.UserModel) error
	GetAllUser() ([]model.UserModel, error)
	GetUserByName(string) (*model.UserModel, error)
}

type userUsecaseImpl struct {
	usrRepo repo.UserRepo
}

func (usrUsecase *userUsecaseImpl) GetUserByName(name string) (*model.UserModel, error) {
	usr, err := usrUsecase.usrRepo.GetUserByName(name)
	if err != nil {
		return nil, fmt.Errorf("usrUsecase.usrRepo.GetUserByName() : %w" , err)
	}
	if usr == nil {
		return nil, apperror.AppError{
			ErrorCode: 400,
			ErrorMessage: fmt.Sprintf("data user dengan nama : %s tidak ada", name),
		}
	}

	return usr, nil
}

func (usrUsecase *userUsecaseImpl) InsertUser(usr *model.UserModel) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password),bcrypt.DefaultCost)  //hashing
	if err != nil {
		return nil
	} 
	usr.Password = string(hashedPassword)

	return usrUsecase.usrRepo.InsertUser(usr)
}

func (usrUsecase *userUsecaseImpl) GetAllUser() ([]model.UserModel, error) {
	return usrUsecase.usrRepo.GetAllUser()
}

func NewUserUsecase(usrRepo repo.UserRepo) UserUsecase {
	return &userUsecaseImpl{
		usrRepo: usrRepo,
	}
}