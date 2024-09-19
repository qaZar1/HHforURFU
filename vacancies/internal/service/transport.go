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
	"github.com/qaZar1/HHforURFU/vacancies/autogen"
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
func (transport *Transport) PostApiVacanciesAdd(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var vacancy autogen.Vacancy
	if err := jsoniter.Unmarshal(data, &vacancy); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if err := transport.srv.db.AddVacancy(vacancy); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить вакансию")
		return
	}

	utils.WriteNoContent(w)
}

// Удаление вакансии из БД
// (DELETE /api/seekers/{chat_id}/remove)
func (transport *Transport) DeleteApiVacanciesVacancyIdRemove(w http.ResponseWriter, r *http.Request, chatId int64) {
	ok, err := transport.srv.RemoveVacancy(chatId)
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
func (transport *Transport) GetApiVacanciesVacancyIdGet(w http.ResponseWriter, r *http.Request, chatId int64) {
	user, err := transport.srv.GetVacancyByVacancyID(chatId)
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
func (transport *Transport) GetApiVacanciesGet(w http.ResponseWriter, r *http.Request) {
	users, err := transport.srv.GetAllVacancies()
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
func (transport *Transport) PutApiVacanciesVacancyIdUpdate(w http.ResponseWriter, r *http.Request, vacancyId int64) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var updateVacancy autogen.UpdateVacancy
	if err := jsoniter.Unmarshal(data, &updateVacancy); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	ok, err := transport.srv.UpdateVacancy(vacancyId, updateVacancy)
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
