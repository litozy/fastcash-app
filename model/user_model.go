package model

import "time"

type UserModel struct {
	Id        int
	UserName  string
	Password  string 
	CreatedAt time.Time
	UpdatedAt time.Time
}