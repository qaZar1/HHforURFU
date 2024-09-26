package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/employers/autogen"
)

type Transport struct {
	srv ServiceInterface // Изменено на интерфейс
}

func NewTransport(srv ServiceInterface) autogen.ServerInterface {
	return &Transport{
		srv: srv,
	}
}

func (transport *Transport) PostApiEmployersAdd(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var user autogen.User
	if err := jsoniter.Unmarshal(data, &user); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	ok, err := transport.srv.AddUser(user)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить пользователя")
		return
	}

	if !ok {
		utils.WriteString(w, http.StatusConflict, err, "Пользователь уже существует")
		return
	}

	utils.WriteNoContent(w)
}

func (transport *Transport) DeleteApiEmployersChatIdRemove(w http.ResponseWriter, r *http.Request, chatId int64) {
	ok, err := transport.srv.RemoveUser(chatId)
	if err != nil {
		utils.WriteString(w, http.StatusNotFound, err, "Пользователя не существует")
		return
	}

	if ok {
		utils.WriteNoContent(w)
	} else {
		utils.WriteString(w, http.StatusNotFound, nil, "Пользователь не найден")
	}
}

func (transport *Transport) GetApiEmployersChatIdGet(w http.ResponseWriter, r *http.Request, chatId int64) {
	user, err := transport.srv.GetUserByChatID(chatId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteNoContent(w)
			return
		}
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить пользователя")
		return
	}

	utils.WriteObject(w, user)
}

func (transport *Transport) GetApiEmployersGet(w http.ResponseWriter, r *http.Request) {
	users, err := transport.srv.GetAllUsers()
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить пользователей")
		return
	}
	if len(users) == 0 {
		utils.WriteNoContent(w)
		return
	}

	utils.WriteObject(w, users)
}

func (transport *Transport) PutApiEmployersChatIdUpdate(w http.ResponseWriter, r *http.Request, chatId int64) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var updateUser autogen.UpdateUser
	if err := jsoniter.Unmarshal(data, &updateUser); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	ok, err := transport.srv.UpdateUser(chatId, updateUser)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось обновить данные пользователя")
		return
	}

	if ok {
		utils.WriteNoContent(w)
	} else {
		utils.WriteString(w, http.StatusNotFound, nil, "Пользователь не найден")
	}
}
