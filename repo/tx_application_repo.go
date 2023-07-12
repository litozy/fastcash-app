package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"peminjaman/utils"
)

type TransactionApplyRepo interface {
	InsertApplication(*model.TransactionApply) error
	GetAllApp() ([]model.TransactionApplyView, error)
	GetAppById(int) (*model.TransactionApplyView, error)
	UpdateStatusOjk(*model.TransactionApply) error
}

type transactionApplyImpl struct {
	db *sql.DB
}

func (taRepo *transactionApplyImpl) GetAllApp() ([]model.TransactionApplyView, error) {
	qry := utils.GET_ALL_TRANSACTION_APPLICATION
	rows, err := taRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllApp() : %w", err)
	}
	defer rows.Close()

	var arrayTr []model.TransactionApplyView
	for rows.Next() {
		tra := model.TransactionApplyView{}
		err := rows.Scan(
			&tra.Id, &tra.CustomerId, &tra.CustomerName, &tra.Nik, &tra.Product, &tra.Amount, &tra.DateApproval, &tra.StatusOjk,
		)
		if err != nil {
			return nil, fmt.Errorf("getAllTransaction(): %w", err)
		}
		arrayTr = append(arrayTr, tra)
	}
	

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllTransaction(): %w", err)
	}
	return arrayTr, nil
}

func (taRepo *transactionApplyImpl) GetAppById(id int) (*model.TransactionApplyView, error) {
	qry := utils.GET_TRANSACTION_APPLICATION_BY_ID
tra := &model.TransactionApplyView{}
err := taRepo.db.QueryRow(qry, id).Scan(&tra.Id, &tra.CustomerId, &tra.CustomerName, &tra.Nik, &tra.Product, &tra.Amount, &tra.DateApproval, &tra.StatusOjk)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceById() : %w", err)
	}
	return tra, nil
}

func (taRepo *transactionApplyImpl) InsertApplication(tra *model.TransactionApply) error {
	tx, err := taRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("InsertTransaction() Begin : %w", err)
	}

	tra.OjkStatusId = 1
	tra.DateApproval = "0001-01-01"
	qry := utils.INSERT_TRANSACTION_APPLICATION

	_, err = tx.Exec(qry, &tra.CustomerId, &tra.ProductId, &tra.Amount, tra.OjkStatusId, tra.DateApproval)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("InsertTransaction() Detail : %w", err)
	}
	tx.Commit()

	return nil
}

func (taRepo *transactionApplyImpl) UpdateStatusOjk(tra *model.TransactionApply) error {
	tx, err := taRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("UpdateStatusOjk() Begin : %w", err)
	}

	qry := utils.UPDATE_OJK_STATUS_TRANSACTION_APPLICATION

	_, err = tx.Exec(qry, &tra.OjkStatusId, &tra.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdateStatusOjk() Detail : %w", err)
	}
	tx.Commit()

	return nil
}


func NewTransactionApplyRepo(db *sql.DB) TransactionApplyRepo {
	return &transactionApplyImpl{
		db: db,
	}
}