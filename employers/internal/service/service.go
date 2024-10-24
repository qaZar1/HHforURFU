// internal/service/service.go
package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/employers/internal/database"
	"github.com/qaZar1/HHforURFU/employers/internal/models"
)

type Service struct {
	db database.Database
}

// NewService создает новый сервис
func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

// Реализация интерфейса
func (srv *Service) RegisterEmployer(seeker models.Employer) (bool, error) {
	return srv.db.AddEmployer(seeker)
}

func (srv *Service) CheckEmployer(employerID int64) (models.Employer, error) {
	return srv.db.CheckEmployerUsername(employerID)
}

func (srv *Service) CheckEmployerByUsername(username string) (models.Employer, error) {
	return srv.db.CheckEmployerByUsername(username)
}
