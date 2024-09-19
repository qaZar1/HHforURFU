package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/vacancies/autogen"
	"github.com/qaZar1/HHforURFU/vacancies/internal/database"
)

type Service struct {
	db database.Database
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

func (srv *Service) GetAllVacancies() ([]autogen.Info, error) {
	return srv.db.GetAllVacancies()
}

func (srv *Service) GetVacancyByVacancyID(vacancyId int64) (*autogen.Info, error) {
	return srv.db.GetVacancyByVacancyID(vacancyId)
}

func (srv *Service) AddVacancy(vacancy autogen.Vacancy) error {
	return srv.db.AddVacancy(vacancy)
}

func (srv *Service) RemoveVacancy(vacancyId int64) (bool, error) {
	return srv.db.RemoveVacancy(vacancyId)
}

func (srv *Service) UpdateVacancy(vacancyId int64, updateVacancy autogen.UpdateVacancy) (bool, error) {
	return srv.db.UpdateVacancy(vacancyId, updateVacancy)
}
