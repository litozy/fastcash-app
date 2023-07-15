package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"peminjaman/utils"
	"time"
)

type TransactionApplyRepo interface {
	InsertApplication(*model.TransactionApply) error
	GetAllApp() ([]model.TransactionApplyView, error)
	GetAppById(int) (*model.TransactionApplyView, error)
	UpdateStatusOjk(*model.TransactionApply) error
	UpdateAmountForLateInterest(*model.TransactionApply) error
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
		var date time.Time
		err := rows.Scan(
			&tra.Id, &tra.CustomerId, &tra.CustomerName, &tra.CustomerBankAccount, &tra.Nik, &tra.ProductRaw, &tra.Product, &tra.Amount, &date, &tra.StatusOjk, &tra.StatusDetail,
		)
		if err != nil {
			return nil, fmt.Errorf("getAllTransaction(): %w", err)
		}
		tra.DateApproval = date.Format("2006-01-02")
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
	var date time.Time
	err := taRepo.db.QueryRow(qry, id).Scan(&tra.Id, &tra.CustomerId, &tra.CustomerName, &tra.CustomerBankAccount, &tra.Nik, &tra.ProductRaw, &tra.Product, &tra.Amount, &date, &tra.StatusOjk, &tra.StatusDetail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on transactionApplyImpl.GetAppById() : %w", err)
	}
	tra.DateApproval = date.Format("2006-01-02")
	return tra, nil
}

func (taRepo *transactionApplyImpl) InsertApplication(tra *model.TransactionApply) error {
	tx, err := taRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("InsertTransaction() Begin : %w", err)
	}

	tra.OjkStatusId = 1
	tra.DateApproval = "0001-01-01"
	tra.UpdatedBy = ""
	qry := utils.INSERT_TRANSACTION_APPLICATION

	_, err = tx.Exec(qry, &tra.CustomerId, &tra.ProductId, &tra.Amount, tra.OjkStatusId, tra.DateApproval, tra.UpdatedBy, &tra.CustomerId)
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

	_, err = tx.Exec(qry, &tra.OjkStatusId, &tra.UpdatedBy, &tra.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdateStatusOjk() Detail : %w", err)
	}

	tra.DateApproval = time.Now().Format("2006-01-02")
	qryDate := utils.UPDATE_OJK_STATUS_DATE_APPROVAL
	if tra.OjkStatusId == 2 {
		_, err = tx.Exec(qryDate, tra.DateApproval, &tra.Id)
		if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdateStatusOjk.DateApproval() Detail : %w", err)
	}
	}
	tx.Commit()

	return nil
}

func (taRepo *transactionApplyImpl) UpdateAmountForLateInterest(tra *model.TransactionApply) error {
	tx, err := taRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("UpdateAmountForLateInterest() Begin : %w", err)
	}

	qry := "UPDATE tx_application SET amount = amount + (amount * (SELECT late_interest FROM loan_product) / 100) WHERE id = $1"

	_, err = tx.Exec(qry, &tra.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdateAmountForLateInterest() Detail : %w", err)
	}
	tx.Commit()

	return nil
}


func NewTransactionApplyRepo(db *sql.DB) TransactionApplyRepo {
	return &transactionApplyImpl{
		db: db,
	}
}