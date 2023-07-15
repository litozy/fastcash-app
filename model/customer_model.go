package model

import "time"

type CustomerModel struct {
	Id            int
	UserId        int
	Name          string
	Address       string
	Gender 		  string
	NIK           string
	Phone 		  string
	BankAccount   string
	Birthdate     time.Time
	FamilyMember  string
	FamilyPhone   string
	FamilyAddress string
	Status        string
}
