package manager

import (
	"peminjaman/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUsecase
	GetLoginUsecase() usecase.LoginUsecase
	GetInterestUsecase() usecase.InterestUsecase
	GetLateInterestUsecase() usecase.LateInterestUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	usrUsecase    usecase.UserUsecase
	lgUsecase     usecase.LoginUsecase
	intrsUsecase  usecase.InterestUsecase
	lintrsUsecase usecase.LateInterestUsecase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadInterestUsecase sync.Once
var onceLoadLateInterestUsecase sync.Once

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
func (um *usecaseManager) GetInterestUsecase() usecase.InterestUsecase {
	onceLoadInterestUsecase.Do(func() {
		um.intrsUsecase = usecase.NewInterestUsecase(um.repoManager.GetInterestRepo())

	})
	return um.intrsUsecase
}
func (um *usecaseManager) GetLateInterestUsecase() usecase.LateInterestUsecase {
	onceLoadLateInterestUsecase.Do(func() {
		um.lintrsUsecase = usecase.NewLateInterestUsecase(um.repoManager.GetLateInterestRepo())

	})
	return um.lintrsUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
