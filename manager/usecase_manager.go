package manager

import (
	"peminjaman/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUsecase
	GetLoginUsecase() usecase.LoginUsecase
	GetLoanProductUsecase() usecase.LoanProductUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	usrUsecase    usecase.UserUsecase
	lgUsecase     usecase.LoginUsecase
	lprdctUsecase usecase.LoanProductUsecase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadLoanProductUsecase sync.Once

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
func (um *usecaseManager) GetLoanProductUsecase() usecase.LoanProductUsecase {
	onceLoadLoanProductUsecase.Do(func() {
		um.lprdctUsecase = usecase.NewLoanProductUsecase(um.repoManager.GetLoanProductRepo())

	})
	return um.lprdctUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
