package manager

import (
	"peminjaman/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUsecase
	GetLoginUsecase() usecase.LoginUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	usrUsecase usecase.UserUsecase
	lgUsecase usecase.LoginUsecase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.UserUsecase {
	onceLoadUserUsecase.Do(func() {
		um.usrUsecase = usecase.NewUserUsecase(um.repoManager.GetUserRepo())

	})
	return um.usrUsecase
}

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUsecase {
	onceLoadLoginUsecase.Do(func() {
		um.lgUsecase = usecase.NewLoginUsecase(um.repoManager.GetUserRepo())

	})
	return um.lgUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}