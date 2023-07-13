package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"time"
)

type TransactionPaymentRepo interface {
	InsertPayment(*model.TransactionPayment) error
}

type transactionPaymentImpl struct {
	db *sql.DB
}

func (taRepo *transactionPaymentImpl) InsertPayment(trp *model.TransactionPayment) error {
	tx, err := taRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("InsertTransaction() Begin : %w", err)
	}

	qry := `
		INSERT INTO tx_payment(aplication_id, payment, created_by, created_at)
		SELECT $1, $2, create_by, $3
		FROM tx_application
		WHERE id = $4
	`

	trp.CreatedAt = time.Now()
	_, err = tx.Exec(qry, &trp.ApplicationId, &trp.Payment, trp.CreatedAt, &trp.ApplicationId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("InsertTransaction() Detail : %w", err)
	}
	tx.Commit()

	return nil
}

func NewTransactionPaymentRepo(db *sql.DB) TransactionPaymentRepo {
	return &transactionPaymentImpl{
		db: db,
	}
}
