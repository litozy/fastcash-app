package model

import "time"

type UserModel struct {
	Id        int
	UserName  string
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}