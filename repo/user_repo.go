package repo

import "database/sql"

type UserRepo interface {
}

type userRepoInpl struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoInpl{
		db: db,
	}
}