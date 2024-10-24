package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/responses/internal/database"
	"github.com/qaZar1/HHforURFU/responses/internal/models"
)

type Service struct {
	db database.Database
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

func (srv *Service) GetResponsesByUsername(username string) ([]models.Response, error) {
	return srv.db.GetResponsesByUsername(username)
}

func (srv *Service) GetResponseByID(id int64) (models.Response, error) {
	return srv.db.GetResponsesById(id)
}

func (srv *Service) GetResponsesByEmployersID(id int64) ([]models.Response, error) {
	return srv.db.GetResponsesByEmployersID(id)
}

func (srv *Service) AddResponses(respons models.Response) error {
	return srv.db.AddResponses(respons)
}

func (srv *Service) UpdateResponse(response_id int, updateRespons models.Response) (bool, error) {
	return srv.db.UpdateResponses(response_id, updateRespons)
}
