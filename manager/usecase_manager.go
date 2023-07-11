package manager

import (
	"peminjaman/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUsecase
	GetLoginUsecase() usecase.LoginUsecase
	GetTransactionAppUsecase() usecase.TransactionApplyUsecase
	GetLoanProductUsecase() usecase.LoanProductUsecase
	GetOjkStatusUsecase() usecase.OjkStatusUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	lgUsecase usecase.LoginUsecase
	taUsecase usecase.TransactionApplyUsecase
	usrUsecase     usecase.UserUsecase
	lprdctUsecase  usecase.LoanProductUsecase
	ojkstatUsecase usecase.OjkStatusUsecase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadTrxApplyUsecase sync.Once
var onceLoadLoanProductUsecase sync.Once
var onceLoadOjkStatusUsecase sync.Once

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
func (um *usecaseManager) GetOjkStatusUsecase() usecase.OjkStatusUsecase {
	onceLoadOjkStatusUsecase.Do(func() {
		um.ojkstatUsecase = usecase.NewOjkStatusUsecase(um.repoManager.GetOjkStatusRepo())

	})
	return um.ojkstatUsecase
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
