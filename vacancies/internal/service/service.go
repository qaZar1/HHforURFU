package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/vacancies/internal/database"
	"github.com/qaZar1/HHforURFU/vacancies/internal/models"
)

type Service struct {
	db database.Database
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

func (srv *Service) GetAllVacancies() ([]models.Vacancy, error) {
	return srv.db.GetAllVacancies()
}

func (srv *Service) GetVacancyByVacancyID(vacancyId int64) (models.Vacancy, error) {
	return srv.db.GetVacancyByVacancyID(vacancyId)
}

func (srv *Service) GetVacancyByEmployerID(employer_id int64) ([]models.Vacancy, error) {
	return srv.db.GetAllVacanciesByEmployerID(employer_id)
}

func (srv *Service) AddVacancy(vacancy models.AddVacancy) (int64, error) {
	return srv.db.AddVacancy(vacancy)
}

func (srv *Service) RemoveVacancy(vacancyId int64) (bool, error) {
	return srv.db.RemoveVacancy(vacancyId)
}

// func (srv *Service) UpdateVacancy(vacancyId int64, updateVacancy autogen.UpdateVacancy) (bool, error) {
// 	return srv.db.UpdateVacancy(vacancyId, updateVacancy)
// }
