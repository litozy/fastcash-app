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

type TransactionApplyView struct {
	CustomerId   int
	CustomerName string
	Nik          int
	Product      string
	Amount       float64
	DateApproval string
	StatusOjk    int
}