package service

import (
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	validator "github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/employers/autogen/server"
	"github.com/qaZar1/HHforURFU/employers/internal/infrastructure"
	"github.com/qaZar1/HHforURFU/employers/internal/models"
)

type Transport struct {
	srv      *Service
	infra    *infrastructure.Infrastructure
	validate *validator.Validate
}

func NewTransport(db *sqlx.DB, expiresIn int) server.ServerInterface {
	return &Transport{
		srv:      NewService(db),
		infra:    infrastructure.New(expiresIn),
		validate: validator.New(),
	}
}

// Set godoc
//
// @Router /api/registerEmployer [post]
// @Summary Регистрация работодателя
// @Description При обращении, регистрируется работодатель
//
// @Tags APIs
// @Accept       application/json
// @Produce      application/json
//
// @Param	request	body	employer	true	"Данные о работодателе"
//
// @Success 204 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PostApiRegisterEmployer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить данные")
		return
	}

	var employer models.Employer
	if err := jsoniter.Unmarshal(body, &employer); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Can not unmarshal body")
		return
	}

	ok, err := transport.srv.RegisterEmployer(employer)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить данные")
		return
	}

	if !ok {
		utils.WriteString(w, http.StatusBadRequest, err, "Данный пользователь уже существует!")
		return
	}

	utils.WriteNoContent(w)
}

// Set godoc
//
// @Router /api/checkEmployer/{employer_id} [get]
// @Summary Проверка, существует ли работодатель
// @Description При обращении, проверятся, сущестувет ли работодатель
//
// @Tags APIs
// @Accept       application/json
// @Produce      application/json
//
// @Param	employer_id	path	int	true	"Username работодателя"
//
// @Success 200 {object} employer "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiCheckEmployerEmployerId(w http.ResponseWriter, r *http.Request, employerID int) {
	employer, err := transport.srv.CheckEmployer(int64(employerID))
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить данные")
		return
	}

	utils.WriteObject(w, employer)
}

// Set godoc
//
// @Router /api/loginEmployer [get]
// @Summary Получение токена
// @Description Генерирует JWT токен
//
// @Tags Authentication
//
// @Produce application/json
//
// @Success 200 {object} token_response "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiLoginEmployer(w http.ResponseWriter, r *http.Request) {
	username, _, _ := r.BasicAuth()
	payload := models.Payload{
		Username: username,
	}
	if err := transport.validate.Struct(payload); err != nil {
		utils.WriteString(w, http.StatusBadRequest, err, "invalid request")
		return
	}

	claims := transport.infra.GetClaims(payload.Username)
	signed, err := transport.infra.GetSignedToken(claims)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "invalid token")
		return
	}

	utils.WriteObject(w, &models.TokenResponse{
		AccessToken: signed,
		ExpiresIn:   18000,
		TokenType:   "Bearer",
	})
}

// Set godoc
//
// @Router /api/checkEmployerByUsername/{username} [get]
// @Summary Проверка, какой id у работодателя с данным username
// @Description При обращении, проверятся, сущестувет ли работодатель
//
// @Tags APIs
// @Accept       application/json
// @Produce      application/json
//
// @Param	username	path	string	true	"Username работодателя"
//
// @Success 200 {object} employer "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiCheckEmployerByUsernameUsername(w http.ResponseWriter, r *http.Request, username string) {
	employer, err := transport.srv.CheckEmployerByUsername(username)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить данные")
		return
	}

	utils.WriteObject(w, employer)
}
