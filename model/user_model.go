package model

type UserModel struct {
	Id       int
	UserName string
	Password string `json:"-"`
	Active   bool
}