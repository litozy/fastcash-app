package model

import "time"

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
	CustomerBankAccount int
	Nik          int
	Product      int
	Amount       float64
	DateApproval string `json:"DateApproval,omitempty"`
	StatusOjk    int
	StatusDetail string
}

type TransactionPayment struct {
	Id            int
	ApplicationId int
	Payment       float64
	CreatedBy     string
	CreatedAt     time.Time
	Status string
}

type TransactionPaymentView struct {
	CustomerId int
	CustomerName string
	Product string
	MustToPay float64
	Paid float64
	RemainingPayment float64
	OneMonthPayment float64
	CompanyBankAccount int
	PaymentDeadline string
}