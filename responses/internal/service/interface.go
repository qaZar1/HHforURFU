package service

import (
	"github.com/qaZar1/HHforURFU/responses/autogen"
)

type ServiceInterface interface {
	GetAllResponses() ([]autogen.Response, error)
	GetResponsesByVacancyIDAndChatIDEmployer(vacancyId int64, chatIdEmployer int64) (*autogen.Response, error)
	AddResponses(response autogen.Response) error
	RemoveResponses(vacancyId int64) (bool, error)
	UpdateRespons(vacancyId int64, updateResponse autogen.Response) (bool, error)
	GetResponsesByChatIDEmployer(chatIdEmployer int64) ([]autogen.Response, error)
	GetResponsesByChatID(chatId int64) ([]autogen.Response, error)
	GetResponsesByVacancyIDAndChatID(vacancyId int64, chatId int64) (*autogen.Response, error)
	GetResponsesByVacancyID(vacancyId int64) ([]autogen.Response, error)
}
