// internal/service/service.go
package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/employers/autogen"
	"github.com/qaZar1/HHforURFU/employers/internal/database"
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
func (srv *Service) GetAllUsers() ([]autogen.Info, error) {
	return srv.db.GetAllUsers()
}

func (srv *Service) GetUserByChatID(chatId int64) (*autogen.Info, error) {
	return srv.db.GetUserByChatID(chatId)
}

func (srv *Service) AddUser(user autogen.User) (bool, error) {
	return srv.db.AddUser(user)
}

func (srv *Service) RemoveUser(chatId int64) (bool, error) {
	return srv.db.RemoveUser(chatId)
}

func (srv *Service) UpdateUser(chatId int64, updateUser autogen.UpdateUser) (bool, error) {
	return srv.db.UpdateUser(chatId, updateUser)
}
