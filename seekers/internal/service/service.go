package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/seekers/internal/database"
	"github.com/qaZar1/HHforURFU/seekers/internal/models"
)

type Service struct {
	db database.Database
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

func (srv *Service) RegisterSeeker(seeker models.Seeker) (bool, error) {
	return srv.db.AddSeeker(seeker)
}

func (srv *Service) CheckSeekerUsername(username string) (models.Seeker, error) {
	return srv.db.CheckSeekerUsername(username)
}

// func (srv *Service) AddVacancy(vacancy autogen.Vacancy) (int64, error) {
// 	return srv.db.AddVacancy(vacancy)
// }

// func (srv *Service) RemoveVacancy(vacancyId int64) (bool, error) {
// 	return srv.db.RemoveVacancy(vacancyId)
// }

// func (srv *Service) UpdateVacancy(vacancyId int64, updateVacancy autogen.UpdateVacancy) (bool, error) {
// 	return srv.db.UpdateVacancy(vacancyId, updateVacancy)
// }
