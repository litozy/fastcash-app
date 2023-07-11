package manager

import (
	"peminjaman/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUsecase
	GetLoginUsecase() usecase.LoginUsecase
	GetTransactionAppUsecase() usecase.TransactionApplyUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	usrUsecase usecase.UserUsecase
	lgUsecase usecase.LoginUsecase
	taUsecase usecase.TransactionApplyUsecase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadTrxApplyUsecase sync.Once

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

func (um *usecaseManager) GetTransactionAppUsecase() usecase.TransactionApplyUsecase {
	onceLoadTrxApplyUsecase.Do(func() {
		um.taUsecase = usecase.NewTransactionApplyUsecase(um.repoManager.GetTransactionAppRepo())

	})
	return um.taUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}