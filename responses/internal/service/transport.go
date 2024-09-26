package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/responses/autogen"
)

type Transport struct {
	srv ServiceInterface // Изменено на интерфейс
}

func NewTransport(srv ServiceInterface) autogen.ServerInterface {
	return &Transport{
		srv: srv,
	}
}

func (transport *Transport) PostApiResponsesAdd(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var response autogen.Response
	if err := jsoniter.Unmarshal(data, &response); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if err := transport.srv.AddResponses(response); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить вакансию")
		return
	}

	utils.WriteNoContent(w)
}

// Удаление вакансии из БД
// (DELETE /api/seekers/{chat_id}/remove)
func (transport *Transport) DeleteApiResponsesVacancyIdRemove(w http.ResponseWriter, r *http.Request, vacancyId int64) {
	ok, err := transport.srv.RemoveResponses(vacancyId)
	if err != nil {
		utils.WriteString(w, http.StatusNotFound, err, "Пользователя не существует")
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
func (transport *Transport) GetApiResponsesVacancyIdVacancyIdGet(w http.ResponseWriter, r *http.Request, vacancyId int64) {
	user, err := transport.srv.GetResponsesByVacancyID(vacancyId)
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

// Set godoc
//
// @Router /api/seekers/get [get]
// @Summary Получение всех откликов из БД
// @Description Хуй
//
// @Tags APIs
// @Accept       application/json
// @Produce      application/json
//
// @Success 200 {object} response "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
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
func (transport *Transport) PutApiResponsesVacancyIdUpdate(w http.ResponseWriter, r *http.Request, vacancyId int64) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var updateRespons autogen.Response
	if err := jsoniter.Unmarshal(data, &updateRespons); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	ok, err := transport.srv.UpdateRespons(vacancyId, updateRespons)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось обновить данные пользователя")
		return
	}

	if !ok {
		utils.WriteNoContent(w)
		return
	} else {
		utils.WriteString(w, http.StatusNotFound, err, "Пользователь не найден")
		return
	}
}

func (transport *Transport) GetApiResponsesChatIdEmployerChatIdEmployerGet(w http.ResponseWriter, r *http.Request, chatIdEmployer int64) {
	users, err := transport.srv.GetResponsesByChatIDEmployer(chatIdEmployer)
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

func (transport *Transport) GetApiResponsesChatIdChatIdGet(w http.ResponseWriter, r *http.Request, chatId int64) {
	users, err := transport.srv.GetResponsesByChatID(chatId)
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

func (transport *Transport) GetApiResponsesVacancyIdAndChatIdEmployerGet(w http.ResponseWriter, r *http.Request, vacancyId int64, chatIdEmployer int64) {
	users, err := transport.srv.GetResponsesByVacancyIDAndChatID(vacancyId, chatIdEmployer)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить вакансии")
		return
	}
	if users == nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "В базе нет вакансий")
		return
	}

	utils.WriteObject(w, users)
}
