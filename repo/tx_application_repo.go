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
}

type transactionApplyImpl struct {
	db *sql.DB
}

func (taRepo *transactionApplyImpl) GetAllApp() ([]model.TransactionApplyView, error) {
	qry := `
		SELECT tx.customer_id AS custId, c.name AS custName, c.nik AS nik, p.tenor AS tenor, tx.amount AS amount, tx.date_approval, o.status AS status 
		FROM tx_application AS tx
		JOIN customer AS c ON tx.customer_id = c.id
		JOIN ojk_status AS o ON tx.ojk_status_id = o.id 
		JOIN loan_product AS p ON tx.loan_product_id = p.id
	`
	rows, err := taRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllApp() : %w", err)
	}
	defer rows.Close()

	var arrayTr []model.TransactionApplyView
	for rows.Next() {
		tra := model.TransactionApplyView{}
		err := rows.Scan(
			&tra.CustomerId, &tra.CustomerName, &tra.Nik, &tra.Product, &tra.Amount, &tra.DateApproval, &tra.StatusOjk,
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

func (taRepo *transactionApplyImpl) InsertApplication(tra *model.TransactionApply) error {
	tx, err := taRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("InsertTransaction() Begin : %w", err)
	}

	tra.OjkStatus = 1
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