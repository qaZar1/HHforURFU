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

func (srv *Service) GetAllResponses() ([]autogen.Response, error) {
	return srv.db.GetAllResponses()
}

func (srv *Service) GetResponsesByVacancyIDAndChatIDEmployer(vacancyId int64, chatIdEmployer int64) (*autogen.Response, error) {
	return srv.db.GetResponsesByVacancyIDAndChatIDEmployer(vacancyId, chatIdEmployer)
}

func (srv *Service) AddResponses(respons autogen.Response) error {
	return srv.db.AddResponses(respons)
}

func (srv *Service) RemoveResponses(vacancyId int64) (bool, error) {
	return srv.db.RemoveResponses(vacancyId)
}

func (srv *Service) UpdateRespons(vacancyId int64, updateRespons autogen.Response) (bool, error) {
	return srv.db.UpdateResponses(vacancyId, updateRespons)
}

func (srv *Service) GetResponsesByChatIDEmployer(chatIdEmployer int64) ([]autogen.Response, error) {
	return srv.db.GetResponsesByChatIDEmployer(chatIdEmployer)
}

func (srv *Service) GetResponsesByChatID(chatId int64) ([]autogen.Response, error) {
	return srv.db.GetResponsesByChatID(chatId)
}

func (srv *Service) GetResponsesByVacancyIDAndChatID(vacancyId int64, chatId int64) (*autogen.Response, error) {
	return srv.db.GetResponsesByVacancyIDAndChatID(vacancyId, chatId)
}

func (srv *Service) GetResponsesByVacancyID(vacancyId int64) ([]autogen.Response, error) {
	return srv.db.GetResponsesByVacancyID(vacancyId)
}
