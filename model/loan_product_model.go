package model

type LoanProductModel struct {
	Id           int
	ProductName  string
	Tenor        string
	MaxLoan      float64
	Interest     float64
	LateInterest float64
}
