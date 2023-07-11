package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"peminjaman/utils"
)

type LoanProductRepo interface {
	GetLoanProductById(int) (*model.LoanProductModel, error)
	GetAllLoanProduct() ([]*model.LoanProductModel, error)
	InsertLoanProduct(*model.LoanProductModel) error
	DeleteLoanProduct(int) error
	UpdateLoanProduct(*model.LoanProductModel) error
}

type loanProductRepoImpl struct {
	db *sql.DB
}

func (lprdctRepo *loanProductRepoImpl) GetLoanProductById(id int) (*model.LoanProductModel, error) {
	qry := utils.GET_LOAN_PRODUCT_BY_ID
	lprdct := &model.LoanProductModel{}
	err := lprdctRepo.db.QueryRow(qry, id).Scan(&lprdct.Id, &lprdct.ProductName, &lprdct.Tenor, &lprdct.MaxLoan, &lprdct.Interest, &lprdct.LateInterest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on loanProductRepoImpl.getLoanProductById() : %v", &err)
	}
	return lprdct, nil
}

func (lprdctRepo *loanProductRepoImpl) GetAllLoanProduct() ([]*model.LoanProductModel, error) {
	qry := utils.GET_ALL_LOAN_PRODUCT
	var arrLoanProduct []*model.LoanProductModel
	rows, err := lprdctRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllLoanProduct error : %v", &err)
	}

	for rows.Next() {
		lprdct := &model.LoanProductModel{}
		rows.Scan(&lprdct.Id, &lprdct.ProductName, &lprdct.Tenor, &lprdct.MaxLoan, &lprdct.Interest, &lprdct.LateInterest)
		arrLoanProduct = append(arrLoanProduct, lprdct)
	}
	return arrLoanProduct, nil

}

func (lprdctRepo *loanProductRepoImpl) InsertLoanProduct(lprdct *model.LoanProductModel) error {
	qry := utils.INSERT_LOAN_PRODUCT
	_, err := lprdctRepo.db.Exec(qry, lprdct.ProductName, lprdct.Tenor, lprdct.MaxLoan, lprdct.Interest, lprdct.LateInterest)
	if err != nil {
		return fmt.Errorf("error on loanProductRepoImpl.InsertLoanProduct() : %w", err)
	}
	return nil
}

func (lprdctRepo *loanProductRepoImpl) DeleteLoanProduct(id int) error {
	qry := utils.DELETE_LOAN_PRODUCT
	_, err := lprdctRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on loanProductRepoImpl.DeleteLoanProduct : %v", &err)
	}
	return nil
}

func (lprdctRepo *loanProductRepoImpl) UpdateLoanProduct(lprdct *model.LoanProductModel) error {
	qry := utils.UPDATE_LOAN_PRODUCT
	_, err := lprdctRepo.db.Exec(qry, lprdct.Id, lprdct.ProductName, lprdct.Tenor, lprdct.MaxLoan, lprdct.Interest, lprdct.LateInterest)
	if err != nil {
		return fmt.Errorf("error on loanProductRepoImpl.UpdateLoanProduct : %v", &err)
	}
	return nil
}

// func (lprdctRepo *loanProductRepoImpl) GetLoanProductByName(name string) (*model.LoanProductModel, error) {
// 	qry := utils.GET_SERVICE_BY_NAME

// 	lprdct := &model.LoanProductModel{}
// 	err := lprdctRepo.db.QueryRow(qry, name).Scan(&lprdct.Id, &lprdct.Name, &lprdct.Uom, &lprdct.Price)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("error on loanProductRepoImpl.GetLoanProductByName() : %w", err)
// 	}
// 	return lprdct, nil
// }

func NewLoanProductRepo(db *sql.DB) LoanProductRepo {
	return &loanProductRepoImpl{
		db: db,
	}
}
