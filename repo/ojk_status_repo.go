package repo

import (
	"database/sql"
	"fmt"
	"peminjaman/model"
	"peminjaman/utils"
)

type OjkStatusRepo interface {
	GetOjkStatusById(id int) (*model.OjkStatusModel, error)
	GetAllOjkStatus() ([]*model.OjkStatusModel, error)
	InsertOjkStatus(ojkstat *model.OjkStatusModel) error
	DeleteOjkStatus(id int) error
	UpdateOjkStatus(ojkstat *model.OjkStatusModel) error
}

type ojkStatusRepoImpl struct {
	db *sql.DB
}

// OJK STATUS HANYA MEMPUNYAI FUNGSI VIEW DAN CREATE

func (ojkstatRepo *ojkStatusRepoImpl) GetOjkStatusById(id int) (*model.OjkStatusModel, error) {
	qry := utils.GET_LOAN_PRODUCT_BY_ID
	ojkstat := &model.OjkStatusModel{}
	err := ojkstatRepo.db.QueryRow(qry, id).Scan(&ojkstat.Id, &ojkstat.Status, &ojkstat.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on ojkStatusRepoImpl.getOjkStatusById() : %v", &err)
	}
	return ojkstat, nil
}

func (ojkstatRepo *ojkStatusRepoImpl) GetAllOjkStatus() ([]*model.OjkStatusModel, error) {
	qry := utils.GET_ALL_LOAN_PRODUCT
	var arrOjkStatus []*model.OjkStatusModel
	rows, err := ojkstatRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("getAllOjkStatus error : %v", &err)
	}

	for rows.Next() {
		ojkstat := &model.OjkStatusModel{}
		rows.Scan(&ojkstat.Id, &ojkstat.Status)
		arrOjkStatus = append(arrOjkStatus, ojkstat)
	}
	return arrOjkStatus, nil

}

func (ojkstatRepo *ojkStatusRepoImpl) InsertOjkStatus(ojkstat *model.OjkStatusModel) error {
	qry := utils.INSERT_LOAN_PRODUCT
	_, err := ojkstatRepo.db.Exec(qry, ojkstat.Status, ojkstat.Description)
	if err != nil {
		return fmt.Errorf("error on ojkStatusRepoImpl.InsertOjkStatus() : %w", err)
	}
	return nil
}

func (ojkstatRepo *ojkStatusRepoImpl) DeleteOjkStatus(id int) error {
	qry := utils.DELETE_OJK_STATUS
	_, err := ojkstatRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on ojkStatusRepo.Impl.DeleteOjkStatus() : %w", err)
	}
	return nil
}

func (ojkstatRepo *ojkStatusRepoImpl) UpdateOjkStatus(ojkstat *model.OjkStatusModel) error {
	qry := utils.UPDATE_OJK_STATUS
	_, err := ojkstatRepo.db.Exec(qry, ojkstat.Id, ojkstat.Status, ojkstat.Description)
	if err != nil {
		return fmt.Errorf("error on ojkStatusRepoImpl.UpdateOjkStatus() : %w", err)
	}
	return nil
}

func NewOjkStatusRepo(db *sql.DB) OjkStatusRepo {
	return &ojkStatusRepoImpl{
		db: db,
	}
}
