package model

import "time"

type CustomerModel struct {
	Id            int
	UserId        int
	Name          string
	Address       string
	NIK           int
	Birthdate     time.Time
	FamilyMember  string
	FamilyPhone   string
	FamilyAddress string
	Status        string
}
