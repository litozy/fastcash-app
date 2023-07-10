package manager

import (
	"peminjaman/repo"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repo.UserRepo
}

type repoManager struct {
	infraManager InfraManager

	usrRepo repo.UserRepo
}

var onceLoadUserRepo sync.Once

func (rm *repoManager) GetUserRepo() repo.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repo.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}