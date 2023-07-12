package model

type TransactionApply struct {
	Id           int
	CustomerId   int
	ProductId    int
	Amount       float64
	OjkStatusId  int
	DateApproval string
	CreatedBy    string
	UpdatedBy    string
}

type TransactionApplyView struct {
	Id           int
	CustomerId   int
	CustomerName string
	Nik          int
	Product      string
	Amount       float64
	DateApproval string
	StatusOjk    int
}