package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/filters/autogen"
	"github.com/qaZar1/HHforURFU/filters/internal/database"
)

type Service struct {
	db database.Database
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

func (srv *Service) GetAllFilters() ([]autogen.Info, error) {
	return srv.db.GetAllFilters()
}

func (srv *Service) GetFiltersByVacancyID(vacancyId int64) (*[]autogen.Info, error) {
	return srv.db.GetFiltersByVacancyID(vacancyId)
}

func (srv *Service) AddFilters(vacancy autogen.Filters) error {
	return srv.db.AddFilters(vacancy)
}

func (srv *Service) RemoveFilters(vacancyId int64) (bool, error) {
	return srv.db.RemoveFilters(vacancyId)
}
