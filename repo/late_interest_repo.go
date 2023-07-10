package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"peminjaman/utils"
)

type LateInterestRepo interface {
	GetLateInterestById(int) (*model.InterestModel, error)
	GetAllLateInterest() ([]*model.InterestModel, error)
	InsertLateInterest(*model.InterestModel) error
	DeleteLateInterest(int) error
	UpdateLateInterest(*model.InterestModel) error
}

type lateInterestRepoImpl struct {
	db *sql.DB
}

func (lintrsRepo *lateInterestRepoImpl) GetLateInterestById(id int) (*model.InterestModel, error) {
	qry := utils.GET_LATE_INTEREST_BY_ID
	lintrs := &model.InterestModel{}
	err := lintrsRepo.db.QueryRow(qry, id).Scan(&lintrs.Id, &lintrs.InterestRate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on lateInterestRepoImpl.getInterestById() : %v", &err)
	}
	defer lintrsRepo.db.Close()
	return lintrs, nil
}

func (lintrsRepo *lateInterestRepoImpl) GetAllLateInterest() ([]*model.InterestModel, error) {
	qry := utils.GET_ALL_LATE_INTEREST
	var arrLInterest []*model.InterestModel
	rows, err := lintrsRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllLateInterest error : %v", &err)
	}

	defer rows.Close()

	for rows.Next() {
		lintrs := &model.InterestModel{}
		rows.Scan(&lintrs.Id, &lintrs.InterestRate)
		arrLInterest = append(arrLInterest, lintrs)
	}
	return arrLInterest, nil

}

func (lintrsRepo *lateInterestRepoImpl) InsertLateInterest(lintrs *model.InterestModel) error {
	qry := utils.INSERT_LATE_INTEREST
	_, err := lintrsRepo.db.Exec(qry, lintrs.InterestRate)
	if err != nil {
		return fmt.Errorf("error on lateInterestRepoImpl.InsertInterest() : %w", err)
	}
	return nil
}

func (lintrsRepo *lateInterestRepoImpl) DeleteLateInterest(id int) error {
	qry := utils.DELETE__LATE_INTEREST
	_, err := lintrsRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("delete late interest error : %v", &err)
	}
	return nil
}

func (lintrsRepo *lateInterestRepoImpl) UpdateLateInterest(lintrs *model.InterestModel) error {
	qry := utils.UPDATE_LATE_INTEREST
	_, err := lintrsRepo.db.Exec(qry, lintrs.Id, lintrs.InterestRate)
	if err != nil {
		return fmt.Errorf("update interest error : %v", &err)
	}
	return nil
}

func NewLateInterestRepo(db *sql.DB) LateInterestRepo {
	return &lateInterestRepoImpl{
		db: db,
	}
}
