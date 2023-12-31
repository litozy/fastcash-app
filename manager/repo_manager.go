package manager

import (
	"peminjaman/repo"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repo.UserRepo
	GetLoanProductRepo() repo.LoanProductRepo
	GetOjkStatusRepo() repo.OjkStatusRepo
}

type repoManager struct {
	infraManager InfraManager

	usrRepo     repo.UserRepo
	lprdctRepo  repo.LoanProductRepo
	ojkstatRepo repo.OjkStatusRepo
}

var onceLoadUserRepo sync.Once
var onceLoadLoanProductRepo sync.Once
var onceLoadOjkStatusRepo sync.Once

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

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
