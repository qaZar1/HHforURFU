package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/filters/internal/database"
	"github.com/qaZar1/HHforURFU/filters/internal/models"
)

type Service struct {
	db database.Database
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

func (srv *Service) GetAllFilters() ([]models.Tag, error) {
	return srv.db.GetAllFilters()
}

func (srv *Service) GetFiltersByVacancyID(vacancyId int64) ([]models.Tag, error) {
	return srv.db.GetFiltersByVacancyID(vacancyId)
}

func (srv *Service) AddFilters(vacancy models.Tag) error {
	return srv.db.AddFilters(vacancy)
}

func (srv *Service) RemoveFilters(vacancyId int64) (bool, error) {
	return srv.db.RemoveFilters(vacancyId)
}
