package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
)

type UserRepo interface {
	GetAllUser() ([]model.UserModel, error)
	GetUserByName(string) (*model.UserModel, error)
	InsertUser(*model.UserModel) error
}

type userRepoImpl struct {
	db *sql.DB
}

func (usrRepo *userRepoImpl) GetAllUser() ([]model.UserModel, error)  {
	qry := "SELECT id, user_name, is_active FROM user_credential"

	rows, err := usrRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllUser() : %w", err)
	}
	defer rows.Close()

	var arrUser []model.UserModel
	for rows.Next() {
		usr := &model.UserModel{}
		rows.Scan(&usr.Id, &usr.UserName, &usr.Active)
		arrUser = append(arrUser, *usr)
	}
	return arrUser, nil
}

func (usrRepo *userRepoImpl) InsertUser(usr *model.UserModel) error {
	qry := "INSERT INTO user_credential(id, user_name, password, is_active) VALUES($1, $2, $3, $4)"

	_, err := usrRepo.db.Exec(qry, &usr.Id , &usr.UserName, &usr.Password, &usr.Active)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.InsertUser() : %w", err)
	}
	return nil
}

func (usrRepo *userRepoImpl) GetUserByName(name string) (*model.UserModel, error) {
	qry := "SELECT id, user_name, password, is_active FROM user_credential WHERE user_name = $1"

	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, name).Scan(&usr.Id, &usr.UserName, &usr.Password, &usr.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetUserByName() : %w", err)
	}
	return usr, nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}


