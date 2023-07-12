package manager

import (
	"peminjaman/repo"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repo.UserRepo
	GetTransactionAppRepo() repo.TransactionApplyRepo
	GetLoanProductRepo() repo.LoanProductRepo
	GetOjkStatusRepo() repo.OjkStatusRepo
	GetCustomerRepo() repo.CustomerRepo
}

type repoManager struct {
	infraManager InfraManager

	usrRepo     repo.UserRepo
	lprdctRepo  repo.LoanProductRepo
	ojkstatRepo repo.OjkStatusRepo
	taRepo      repo.TransactionApplyRepo
	cstmRepo    repo.CustomerRepo
}

var onceLoadTrxApplyRepo sync.Once
var onceLoadUserRepo sync.Once
var onceLoadLoanProductRepo sync.Once
var onceLoadOjkStatusRepo sync.Once
var onceLoadCustomerRepo sync.Once

func (rm *repoManager) GetUserRepo() repo.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repo.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}
func (rm *repoManager) GetLoanProductRepo() repo.LoanProductRepo {
	onceLoadLoanProductRepo.Do(func() {
		rm.lprdctRepo = repo.NewLoanProductRepo(rm.infraManager.GetDB())
	})
	return rm.lprdctRepo
}
func (rm *repoManager) GetOjkStatusRepo() repo.OjkStatusRepo {
	onceLoadOjkStatusRepo.Do(func() {
		rm.ojkstatRepo = repo.NewOjkStatusRepo(rm.infraManager.GetDB())
	})
	return rm.ojkstatRepo
}

func (rm *repoManager) GetTransactionAppRepo() repo.TransactionApplyRepo {
	onceLoadTrxApplyRepo.Do(func() {
		rm.taRepo = repo.NewTransactionApplyRepo(rm.infraManager.GetDB())
	})
	return rm.taRepo
}
func (rm *repoManager) GetCustomerRepo() repo.CustomerRepo {
	onceLoadCustomerRepo.Do(func() {
		rm.cstmRepo = repo.NewCustomerRepo(rm.infraManager.GetDB())
	})
	return rm.cstmRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
