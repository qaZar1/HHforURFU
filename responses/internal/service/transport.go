package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/responses/autogen/server"
	"github.com/qaZar1/HHforURFU/responses/internal/models"

	jsoniter "github.com/json-iterator/go"
)

type Transport struct {
	srv ServiceInterface
}

func NewTransport(db *sqlx.DB) server.ServerInterface {
	return &Transport{
		srv: NewService(db),
	}
}

// Set godoc
//
// @Router /api/responses/getResponsesByUsername/{username} [get]
// @Summary Получение ответов по username
// @Description При обращении, возвращаются все ответы, которые есть у username
//
// @Tags APIs
// @Produce      application/json
// @Param 	username	path	string	true	"ID пользователя"
//
// @Success 200 {array} response "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiResponsesGetResponsesByUsernameUsername(w http.ResponseWriter, r *http.Request, username string) {
	responses, err := transport.srv.GetResponsesByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteString(w, http.StatusNotFound, err, "В базе отсутствуют отклики")
			return
		}

		utils.WriteString(w, http.StatusNoContent, err, "Не удалось получить отклики")
		return
	}

	utils.WriteObject(w, responses)
}

// Set godoc
//
// @Router /api/responses/getResponseByResponseID/{response_id} [get]
// @Summary Получение ответов по id
// @Description При обращении, возвращаются все ответы, которые есть у id
//
// @Tags APIs
// @Produce      application/json
// @Param 	response_id	path	int	true	"ID отклика"
//
// @Success 200 {object} response "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiResponsesGetResponseByResponseIDResponseId(w http.ResponseWriter, r *http.Request, id int) {
	response, err := transport.srv.GetResponseByID(int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteString(w, http.StatusNotFound, err, "В базе отсутствуют отклики")
			return
		}

		utils.WriteString(w, http.StatusNoContent, err, "Не удалось получить отклики")
		return
	}

	utils.WriteObject(w, response)
}

// Set godoc
//
// @Router /api/responses/getResponsesByEmployersID/{id} [get]
// @Summary Получение ответов по employers id
// @Description При обращении, возвращаются все ответы, которые есть у employers id
//
// @Tags APIs
// @Produce      application/json
// @Param 	id	path	int	true	"ID работодателя"
//
// @Success 200 {array} response "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiResponsesGetResponsesByEmployersIDId(w http.ResponseWriter, r *http.Request, id int) {
	response, err := transport.srv.GetResponsesByEmployersID(int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteString(w, http.StatusNotFound, err, "В базе отсутствуют отклики")
			return
		}

		utils.WriteString(w, http.StatusNoContent, err, "Не удалось получить отклики")
		return
	}

	utils.WriteObject(w, response)
}

// Set godoc
//
// @Router /api/responses/addResponse [post]
// @Summary Добавление отклика в БД
// @Description При обращении, добавляется отклик в БД по телу запрсоа
//
// @Tags APIs
// @Produce      application/json
// @Param 	request	body	response	true	"Тело запроса"
//
// @Success 200 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PostApiResponsesAddResponse(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var response models.Response
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

// Set godoc
//
// @Router /api/responses/edit/{response_id} [put]
// @Summary Изменение отклика
// @Description При обращении, возвращается обновленный отклик
//
// @Tags APIs
// @Produce      application/json
// @Param 	response_id	path	int	true	"ID отклика"
//
// @Success 204 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PutApiResponsesEditResponseId(w http.ResponseWriter, r *http.Request, response_id int) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var updateResponse models.Response
	if err := jsoniter.Unmarshal(data, &updateResponse); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	ok, err := transport.srv.UpdateResponse(response_id, updateResponse)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось обновить данные пользователя")
		return
	}

	if !ok {
		utils.WriteNoContent(w)
		return
	} else {
		utils.WriteString(w, http.StatusNotFound, err, "Отклик не найден")
		return
	}
}
