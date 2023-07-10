package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"peminjaman/utils"
)

type InterestRepo interface {
	GetInterestById(int) (*model.InterestModel, error)
	GetAllInterest() ([]*model.InterestModel, error)
	InsertInterest(*model.InterestModel) error
	DeleteInterest(int) error
	UpdateInterest(*model.InterestModel) error
}

type interestRepoImpl struct {
	db *sql.DB
}

func (intrsRepo *interestRepoImpl) GetInterestById(id int) (*model.InterestModel, error) {
	qry := utils.GET_INTEREST_BY_ID
	intrs := &model.InterestModel{}
	err := intrsRepo.db.QueryRow(qry, id).Scan(&intrs.Id, &intrs.InterestRate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on interestRepoImpl.getInterestById() : %v", &err)
	}
	defer intrsRepo.db.Close()
	return intrs, nil
}

func (intrsRepo *interestRepoImpl) GetAllInterest() ([]*model.InterestModel, error) {
	qry := utils.GET_ALL_INTEREST
	var arrInterest []*model.InterestModel
	rows, err := intrsRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllInterest error : %v", &err)
	}

	defer rows.Close()

	for rows.Next() {
		intrs := &model.InterestModel{}
		rows.Scan(&intrs.Id, &intrs.InterestRate)
		arrInterest = append(arrInterest, intrs)
	}
	return arrInterest, nil

}

func (intrsRepo *interestRepoImpl) InsertInterest(intrs *model.InterestModel) error {
	qry := utils.INSERT_INTEREST
	_, err := intrsRepo.db.Exec(qry, intrs.InterestRate)
	if err != nil {
		return fmt.Errorf("error on interestRepoImpl.InsertInterest() : %w", err)
	}
	return nil
}

func (intrsRepo *interestRepoImpl) DeleteInterest(id int) error {
	qry := utils.DELETE_INTEREST
	_, err := intrsRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("delete interest error : %v", &err)
	}
	return nil
}

func (intrsRepo *interestRepoImpl) UpdateInterest(intrs *model.InterestModel) error {
	qry := utils.UPDATE_INTEREST
	_, err := intrsRepo.db.Exec(qry, intrs.Id, intrs.InterestRate)
	if err != nil {
		return fmt.Errorf("update interest error : %v", &err)
	}
	return nil
}

// func (intrsRepo *interestRepoImpl) GetInterestByName(name string) (*model.InterestModel, error) {
// 	qry := utils.GET_SERVICE_BY_NAME

// 	intrs := &model.InterestModel{}
// 	err := intrsRepo.db.QueryRow(qry, name).Scan(&intrs.Id, &intrs.Name, &intrs.Uom, &intrs.Price)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("error on interestRepoImpl.GetInterestByName() : %w", err)
// 	}
// 	return intrs, nil
// }

func NewInterestRepo(db *sql.DB) InterestRepo {
	return &interestRepoImpl{
		db: db,
	}
}
