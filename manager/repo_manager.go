package manager

import (
	"peminjaman/repo"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repo.UserRepo
	GetLateInterestRepo() repo.LateInterestRepo
	GetInterestRepo() repo.InterestRepo
}

type repoManager struct {
	infraManager InfraManager

	usrRepo    repo.UserRepo
	intrsRepo  repo.InterestRepo
	lintrsRepo repo.LateInterestRepo
}

var onceLoadUserRepo sync.Once
var onceLoadInterestRepo sync.Once
var onceLoadLInterestRepo sync.Once

func (rm *repoManager) GetUserRepo() repo.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repo.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}
func (rm *repoManager) GetInterestRepo() repo.InterestRepo {
	onceLoadInterestRepo.Do(func() {
		rm.intrsRepo = repo.NewInterestRepo(rm.infraManager.GetDB())
	})
	return rm.intrsRepo
}
func (rm *repoManager) GetLateInterestRepo() repo.LateInterestRepo {
	onceLoadLInterestRepo.Do(func() {
		rm.lintrsRepo = repo.NewLateInterestRepo(rm.infraManager.GetDB())
	})
	return rm.lintrsRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
