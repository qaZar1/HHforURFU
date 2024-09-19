package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/filters/autogen"
)

type Transport struct {
	srv *Service
}

func NewTransport(db *sqlx.DB) autogen.ServerInterface {
	return &Transport{
		srv: NewService(db),
	}
}

// Добавление нового тэга в БД
// (POST /api/seekers/add)
func (transport *Transport) PostApiFiltersAdd(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var vacancy autogen.Filters
	if err := jsoniter.Unmarshal(data, &vacancy); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if err := transport.srv.db.AddFilters(vacancy); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить вакансию")
		return
	}

	utils.WriteNoContent(w)
}

// Удаление тэга из БД
// (DELETE /api/seekers/{chat_id}/remove)
func (transport *Transport) DeleteApiFiltersVacancyIdRemove(w http.ResponseWriter, r *http.Request, chatId int64) {
	ok, err := transport.srv.RemoveFilters(chatId)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Пользователя не существует")
		return
	}

	if ok {
		utils.WriteString(w, http.StatusOK, nil, "Пользователя не существует")
		return
	} else {
		utils.WriteNoContent(w)
		return
	}
}

// Получение тэгов по vacancy_id
// (GET /api/seekers/{vacancy_id}/get)
func (transport *Transport) GetApiFiltersVacancyIdGet(w http.ResponseWriter, r *http.Request, chatId int64) {
	user, err := transport.srv.GetFiltersByVacancyID(chatId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteNoContent(w)
			return
		}

		utils.WriteString(w, http.StatusNoContent, err, "Не удалось получить пользователя")
		return
	}

	utils.WriteObject(w, user)
}

// Получение всех тэгов из БД
// (GET /api/seekers/get)
func (transport *Transport) GetApiFiltersGet(w http.ResponseWriter, r *http.Request) {
	users, err := transport.srv.GetAllFilters()
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить вакансии")
		return
	}
	if len(users) == 0 {
		utils.WriteString(w, http.StatusInternalServerError, err, "В базе нет вакансий")
		return
	}

	utils.WriteObject(w, users)
}
