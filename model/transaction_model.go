package model

type TransactionApply struct {
	Id           int
	CustomerId   int
	ProductId    int
	Amount       float64
	OjkStatus    int
	DateApproval string
	CreatedBy    string
	UpdatedBy    string
}