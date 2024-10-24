package service

import (
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	validator "github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/seekers/autogen/server"
	"github.com/qaZar1/HHforURFU/seekers/internal/infrastructure"
	"github.com/qaZar1/HHforURFU/seekers/internal/models"
	_ "github.com/swaggo/swag"
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
// @Router /api/registerSeeker [post]
// @Summary Регистрация соискателя
// @Description При обращении, регистрируется соискатель
//
// @Tags APIs
// @Accept       application/json
// @Produce      application/json
//
// @Param	request	body	seeker	true	"Данные о пользователе"
//
// @Success 204 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PostApiRegisterSeeker(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить данные")
		return
	}

	var seeker models.Seeker
	if err := jsoniter.Unmarshal(body, &seeker); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Can not unmarshal body")
		return
	}

	ok, err := transport.srv.RegisterSeeker(seeker)
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
// @Router /api/checkSeeker/{username} [get]
// @Summary Проверка, существует ли пользователь
// @Description При обращении, регистрируется соискатель
//
// @Tags APIs
// @Accept       application/json
// @Produce      application/json
//
// @Param	username	path	string	true	"Username пользователя"
//
// @Success 200 {object} seeker "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiCheckSeekerUsername(w http.ResponseWriter, r *http.Request, username string) {
	seeker, err := transport.srv.CheckSeekerUsername(username)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить данные")
		return
	}

	utils.WriteObject(w, seeker)
}

// Set godoc
//
// @Router /api/login [get]
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
func (transport *Transport) GetApiLogin(w http.ResponseWriter, r *http.Request) {
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
