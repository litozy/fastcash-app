package usecase

import (
	"fmt"
	"peminjaman/apperror"
	"peminjaman/repo"

	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase interface {
	GetUserByNameAndPassword(string, string) error
}

type loginUsecaseImpl struct {
	usrRepo repo.UserRepo
}

func (lgnUsecase *loginUsecaseImpl) GetUserByNameAndPassword(name string, pass string) error {
	usr, err := lgnUsecase.usrRepo.GetUserByName(name)
	if err != nil {
		return fmt.Errorf("usrUsecase.usrRepo.GetUserByName() : %w" , err)
	}
	if usr == nil {
		return  apperror.AppError{
			ErrorCode: 400,
			ErrorMessage: fmt.Sprintf("data user dengan nama : %s tidak ada", name),
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(pass))
	if err != nil {
		return  apperror.AppError{
			ErrorCode: 400,
			ErrorMessage: "password is incorrect",
		}
	}
	return nil
}

func NewLoginUsecase(usrRepo repo.UserRepo) LoginUsecase {
	return &loginUsecaseImpl{
		usrRepo: usrRepo,
	}
}