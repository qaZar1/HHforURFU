package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/filters/autogen/server"
	"github.com/qaZar1/HHforURFU/filters/internal/models"

	jsoniter "github.com/json-iterator/go"
)

type Transport struct {
	srv *Service
}

func NewTransport(db *sqlx.DB) server.ServerInterface {
	return &Transport{
		srv: NewService(db),
	}
}

// Set godoc
//
// @Router /api/filters/getTagsByVacancyID/{vacancy_id} [get]
// @Summary Получение тегов по vacancy_id
// @Description При обращении, возвращаются все теги, которые есть у вакансии
//
// @Tags APIs
// @Produce      application/json
// @Param 	vacancy_id	path	int	true	"ID вакансии"
//
// @Success 200 {array} tag "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiFiltersGetTagsByVacancyIDVacancyId(w http.ResponseWriter, r *http.Request, vacancyID int) {
	tags, err := transport.srv.GetFiltersByVacancyID(int64(vacancyID))
	if err != nil {
		utils.WriteString(w, http.StatusBadRequest, err, "Invalid get tags by vacancyID")
		return
	}

	if tags == nil {
		utils.WriteString(w, http.StatusBadRequest, err, "Tags not found")
		return
	}

	utils.WriteObject(w, tags)
}

// Set godoc
//
// @Router /api/filters/addFilter [post]
// @Summary Получение ответов по id
// @Description При обращении, создается тег
//
// @Tags APIs
// @Produce      application/json
// @Param	request	body	tag	true	"Данные о фильтре"
//
// @Success 204 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PostApiFiltersAddFilter(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var vacancy models.Tag
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
