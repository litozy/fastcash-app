package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"peminjaman/utils"
)

type CustomerRepo interface {
	GetAllCustomer() ([]*model.CustomerModel, error)
	GetCustomerById(id int) (*model.CustomerModel, error)
	InsertCustomer(cstm *model.CustomerModel) error
	DeleteCustomer(id int) error
	UpdateCustomer(cstm *model.CustomerModel) error
}

type customerRepoImpl struct {
	db *sql.DB
}

func (cstmRepo *customerRepoImpl) GetAllCustomer() ([]*model.CustomerModel, error) {
	qry := utils.GET_ALL_CUSTOMER
	var arrCustomer []*model.CustomerModel
	rows, err := cstmRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllCustomer error : %v", &err)
	}

	for rows.Next() {
		cstm := &model.CustomerModel{}
		rows.Scan(&cstm.Id, &cstm.UserId, &cstm.Name, &cstm.Address, &cstm.NIK, &cstm.Birthdate, &cstm.FamilyMember, &cstm.FamilyPhone, &cstm.FamilyAddress, &cstm.Status)
		arrCustomer = append(arrCustomer, cstm)
	}
	return arrCustomer, nil
}

func (cstmRepo *customerRepoImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	qry := utils.GET_CUSTOMER_BY_ID
	cstm := &model.CustomerModel{}
	err := cstmRepo.db.QueryRow(qry, id).Scan(&cstm.Id, &cstm.UserId, &cstm.Name, &cstm.Address, &cstm.NIK, &cstm.Birthdate, &cstm.FamilyMember, &cstm.FamilyPhone, &cstm.FamilyAddress, &cstm.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("getAllCustomer error : %v", err)
	}
	return cstm, nil
}

func (cstmRepo *customerRepoImpl) InsertCustomer(cstm *model.CustomerModel) error {
	qry := utils.INSERT_CUSTOMER
	_, err := cstmRepo.db.Exec(qry, cstm.UserId, cstm.Name, cstm.Address, cstm.NIK, cstm.Birthdate, cstm.FamilyMember, cstm.FamilyPhone, cstm.FamilyAddress, cstm.Status)
	if err != nil {
		return fmt.Errorf("insertCustomer error : %v", err)
	}
	return nil
}

func (cstmRepo *customerRepoImpl) DeleteCustomer(id int) error {
	qry := utils.DELETE_CUSTOMER
	_, err := cstmRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("DeleteCustomer error : %v", err)
	}
	return nil
}

func (cstmRepo *customerRepoImpl) UpdateCustomer(cstm *model.CustomerModel) error {
	qry := utils.UPDATE_CUSTOMER
	_, err := cstmRepo.db.Exec(qry, cstm.Id, cstm.UserId, cstm.Name, cstm.Address, cstm.NIK, cstm.Birthdate, cstm.FamilyMember, cstm.FamilyPhone, cstm.FamilyAddress, cstm.Status)
	if err != nil {
		return fmt.Errorf("UpdateCustomer error : %v", err)
	}
	return nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepoImpl{
		db: db,
	}
}
