package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"peminjaman/utils"
)

type TransactionApplyRepo interface {
	InsertApplication(*model.TransactionApply) error
}

type transactionApplyImpl struct {
	db *sql.DB
}

func (taRepo transactionApplyImpl) InsertApplication(tra *model.TransactionApply) error {
	tx, err := taRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("InsertTransaction() Begin : %w", err)
	}

	tra.OjkStatus = 0
	tra.DateApproval = "0001-01-01"
	qry := utils.INSERT_TRANSACTION_APPLICATION

	_, err = tx.Exec(qry, &tra.CustomerId, &tra.ProductId, &tra.Amount, tra.OjkStatus, tra.DateApproval)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("InsertTransaction() Detail : %w", err)
	}
	tx.Commit()

	return nil
}



func NewTransactionApplyRepo(db *sql.DB) TransactionApplyRepo {
	return &transactionApplyImpl{
		db: db,
	}
}