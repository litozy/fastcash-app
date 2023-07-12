package model

type LoanProductModel struct {
	Id           int
	ProductName  string
	Tenor        int
	MaxLoan      float64
	Interest     float64
	LateInterest float64
}
