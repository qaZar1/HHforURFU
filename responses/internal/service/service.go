package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/responses/autogen"
	"github.com/qaZar1/HHforURFU/responses/internal/database"
)

type Service struct {
	db database.Database
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

func (srv *Service) GetAllResponses() ([]autogen.Info, error) {
	return srv.db.GetAllResponses()
}

func (srv *Service) GetResponsesByVacancyIDAndChatID(vacancyId int64, chatId int64) (*autogen.Info, error) {
	return srv.db.GetResponsesByVacancyIDAndChatID(vacancyId, chatId)
}

func (srv *Service) AddResponses(respons autogen.Respons) error {
	return srv.db.AddResponses(respons)
}

func (srv *Service) RemoveResponses(vacancyId int64, chatId int64) (bool, error) {
	return srv.db.RemoveResponses(vacancyId, chatId)
}

func (srv *Service) UpdateRespons(vacancyId int64, chatId int64, updateRespons autogen.UpdateRespons) (bool, error) {
	return srv.db.UpdateResponses(vacancyId, chatId, updateRespons)
}
