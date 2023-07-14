package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"peminjaman/utils"
	"time"
)

type TransactionPaymentRepo interface {
	InsertPayment(*model.TransactionPayment) error
	GetPaymentViewById(int) (*model.TransactionPaymentView, error)
	GetPaymentValidateById(id int) (*model.TransactionPaymentView, error)
	UpdateStatusPayment(*model.TransactionPayment) error
}

type transactionPaymentImpl struct {
	db *sql.DB
}

func (taRepo *transactionPaymentImpl) InsertPayment(trp *model.TransactionPayment) error {
	tx, err := taRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("InsertTransaction() Begin : %w", err)
	}

	qry := utils.INSERT_TRANSACTION_PAYMENT

	trp.Status = "checking"
	trp.CreatedAt = time.Now()
	_, err = tx.Exec(qry, &trp.ApplicationId, &trp.Payment, trp.CreatedAt, &trp.ApplicationId, trp.Status)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("InsertTransaction() Detail : %w", err)
	}
	tx.Commit()

	return nil
}

func (taRepo *transactionPaymentImpl) GetPaymentViewById(id int) (*model.TransactionPaymentView, error) {
	qry := utils.GET_TRANSACTION_PAYMENT_BY_ID
	trp := &model.TransactionPaymentView{}
	trp.CompanyBankAccount = 7040594095
	var date time.Time
	err := taRepo.db.QueryRow(qry, id).Scan(&trp.CustomerId, &trp.CustomerName, &trp.Product, &trp.Amount, &trp.Paid, &trp.RemainingPayment, &trp.NeedToPayThisMonth, &date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on transactionPaymentImpl.GetPaymentById() : %w", err)
	}
	trp.PaymentDeadline = date.Format("2006-01-02")
	return trp, nil
}

func (taRepo *transactionPaymentImpl) GetPaymentValidateById(id int) (*model.TransactionPaymentView, error) {
	qry := utils.GET_TRANSACTION_PAYMENT_BY_ID_VALIDATE
	trp := &model.TransactionPaymentView{}
	err := taRepo.db.QueryRow(qry, id).Scan(&trp.CustomerId, &trp.CustomerName, &trp.Product, &trp.Amount, &trp.Paid, &trp.RemainingPayment, &trp.NeedToPayThisMonth)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on transactionPaymentImpl.GetPaymentById() : %w", err)
	}

	return trp, nil
}

func (taRepo *transactionPaymentImpl) UpdateStatusPayment(trp *model.TransactionPayment) error {
	tx, err := taRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("UpdateStatusPayment() Begin : %w", err)
	}

	qry := utils.UPDATE_TRANSACTION_PAYMENT_STATUS

	_, err = tx.Exec(qry, &trp.Status, &trp.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("UpdateStatusPayment() Detail : %w", err)
	}
	tx.Commit()

	return nil
}

func NewTransactionPaymentRepo(db *sql.DB) TransactionPaymentRepo {
	return &transactionPaymentImpl{
		db: db,
	}
}
