package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/telegram_seekers/autogen"
)

type Transport struct {
	srv      *Service
	validate *validator.Validate
}

func NewTransport(token string, seekers string, vacancies string, responses string, tags string) autogen.ServerInterface {
	return &Transport{
		srv: NewService(token,
			seekers,
			vacancies,
			responses,
			tags),
		validate: validator.New(),
	}
}

func (transport *Transport) PostApiNotify(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Неудалось прочитать тело запроса")
		return
	}

	var notification autogen.Notification
	if err := jsoniter.Unmarshal(data, &notification); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if err := transport.validate.Struct(notification); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid body: %s", err), "Невалидное тело запроса")
		return
	}

	if err := transport.srv.Send(notification); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid notify: %s", err), "Неудалось отправить оповещения")
		return
	}

	utils.WriteNoContent(w)
}
