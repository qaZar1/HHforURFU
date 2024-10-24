package service

import (
	"database/sql"
	"errors"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/vacancies/autogen/server"
	"github.com/qaZar1/HHforURFU/vacancies/internal/models"
	_ "github.com/swaggo/swag"
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
// @Router /api/vacancies/getAllVacancy [get]
// @Summary Получение всех вакансий
// @Description При обращении, возвращается массив из вакансий
//
// @Tags APIs
// @Accept       application/json
// @Produce      application/json
//
// @Success 200 {array} vacancy "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiVacanciesGetAllVacancy(w http.ResponseWriter, r *http.Request) {
	allVacancies, err := transport.srv.GetAllVacancies()
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить данные")
		return
	}

	if allVacancies == nil {
		utils.WriteString(w, http.StatusOK, nil, "Вакансий нет")
		return
	}

	utils.WriteObject(w, allVacancies)
}

// Set godoc
//
// @Router /api/vacancies/getVacancyByVacancyID/{id} [get]
// @Summary Получение вакансии по VacancyID
// @Description При обращении, возвращается вакансия
//
// @Tags APIs
// @Produce      application/json
// @Param 	id	path	int	true	"ID вакансии"
//
// @Success 200 {object} vacancy "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiVacanciesGetVacancyByVacancyIDId(w http.ResponseWriter, r *http.Request, id int) {
	user, err := transport.srv.GetVacancyByVacancyID(int64(id))
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
// @Router /api/vacancies/getVacancyByEmployerID/{employer_id} [get]
// @Summary Получение вакансии по VacancyID
// @Description При обращении, возвращается вакансия
//
// @Tags APIs
// @Produce      application/json
// @Param 	employer_id	path	int	true	"ID работодателя"
//
// @Success 200 {object} vacancy "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiVacanciesGetVacancyByEmployerIDEmployerId(w http.ResponseWriter, r *http.Request, employer_id int) {
	allVacancies, err := transport.srv.GetVacancyByEmployerID(int64(employer_id))
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить данные")
		return
	}

	if allVacancies == nil {
		utils.WriteString(w, http.StatusOK, nil, "Вакансий нет")
		return
	}

	utils.WriteObject(w, allVacancies)
}

// Set godoc
//
// @Router /api/vacancies/addVacancy [post]
// @Summary Добавление вакансии
// @Description При обращении, добавляется вакансия и возвращается код
//
// @Tags APIs
// @Produce      application/json
// @Param 	request	body	add_vacancy	true	"Данные о вакансии"
//
// @Success 204 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PostApiVacanciesAddVacancy(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusBadRequest, err, "Ошибка при чтении данных")
		return
	}

	var vacancy models.AddVacancy
	if err := jsoniter.Unmarshal(data, &vacancy); err != nil {
		utils.WriteString(w, http.StatusBadRequest, err, "Ошибка при парсинге данных")
		return
	}

	vacancy_id, err := transport.srv.AddVacancy(vacancy)
	if err != nil {
		utils.WriteString(w, http.StatusBadRequest, err, "Данные не могут быть добавлены")
		return
	}

	if vacancy_id == 0 {
		utils.WriteString(w, http.StatusBadRequest, err, "Данные не могут быть добавлены")
		return
	}

	vac := models.Num{
		VacancyID: vacancy_id,
	}
	utils.WriteObject(w, vac)
}

// Set godoc
//
// @Router /api/vacancies/delete/{vacancy-id} [delete]
// @Summary Удаляется вакансия
// @Description При обращении, добавляется вакансия и возвращается код
//
// @Tags APIs
// @Produce      application/json
// @Param 	vacancy-id	path	int	true	"ID вакансии"
//
// @Success 204 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) DeleteApiVacanciesDeleteVacancyId(w http.ResponseWriter, r *http.Request, vacancyID int) {
	ok, err := transport.srv.RemoveVacancy(int64(vacancyID))
	if err != nil {
		utils.WriteString(w, http.StatusBadRequest, err, "Данные не могут быть удалены")
		return
	}

	if ok {
		utils.WriteString(w, http.StatusBadRequest, err, "Данные не могут быть удалены")
		return
	}

	utils.WriteNoContent(w)
}
