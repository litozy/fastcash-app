package model

import "time"

type TxApplicationModel struct {
	Id            int
	CustomerId    int
	LoanProductId int
	Amount        float64
	OjkStatusId   int
	DateApproval  time.Time
	CreatedBy     string
	UpdatedBy     string
}
