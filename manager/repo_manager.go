package manager

import (
	"peminjaman/repo"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repo.UserRepo
	GetTransactionAppRepo() repo.TransactionApplyRepo
}

type repoManager struct {
	infraManager InfraManager

	usrRepo repo.UserRepo
	taRepo repo.TransactionApplyRepo
}

var onceLoadUserRepo sync.Once
var onceLoadTrxApplyRepo sync.Once

func (rm *repoManager) GetUserRepo() repo.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repo.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func (rm *repoManager) GetTransactionAppRepo() repo.TransactionApplyRepo {
	onceLoadTrxApplyRepo.Do(func() {
		rm.taRepo = repo.NewTransactionApplyRepo(rm.infraManager.GetDB())
	})
	return rm.taRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}