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
	"github.com/qaZar1/HHforURFU/responses/autogen"
)

type Transport struct {
	srv *Service
}

func NewTransport(db *sqlx.DB) autogen.ServerInterface {
	return &Transport{
		srv: NewService(db),
	}
}

// Добавление новой вакансии в БД
// (POST /api/seekers/add)
func (transport *Transport) PostApiResponsesAdd(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var respons autogen.Respons
	if err := jsoniter.Unmarshal(data, &respons); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if err := transport.srv.db.AddResponses(respons); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить вакансию")
		return
	}

	utils.WriteNoContent(w)
}

// Удаление вакансии из БД
// (DELETE /api/seekers/{chat_id}/remove)
func (transport *Transport) DeleteApiResponsesVacancyIdAndChatIdRemove(w http.ResponseWriter, r *http.Request, vacancyId int64, chatId int64) {
	ok, err := transport.srv.RemoveResponses(vacancyId, chatId)
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

// Получение данных вакансии по vacancy_id
// (GET /api/seekers/{vacancy_id}/get)
func (transport *Transport) GetApiResponsesVacancyIdAndChatIdGet(w http.ResponseWriter, r *http.Request, vacancyId int64, chatId int64) {
	user, err := transport.srv.GetResponsesByVacancyIDAndChatID(vacancyId, chatId)
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

// Получение всех вакансий из БД
// (GET /api/seekers/get)
func (transport *Transport) GetApiResponsesGet(w http.ResponseWriter, r *http.Request) {
	users, err := transport.srv.GetAllResponses()
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

// Обновление информации о вакансии в БД
// (PUT /api/vacancies/{vacancy_id}/update)
func (transport *Transport) PutApiResponsesVacancyIdAndChatIdUpdate(w http.ResponseWriter, r *http.Request, vacancyId int64, chatId int64) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var updateRespons autogen.UpdateRespons
	if err := jsoniter.Unmarshal(data, &updateRespons); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	ok, err := transport.srv.UpdateRespons(vacancyId, chatId, updateRespons)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось обновить данные пользователя")
		return
	}

	if ok {
		utils.WriteString(w, http.StatusOK, nil, "Невозможно обновить данные о пользователе")
		return
	} else {
		utils.WriteNoContent(w)
		return
	}
}
