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
	"github.com/qaZar1/HHforURFU/employers/autogen"
)

type Transport struct {
	srv *Service
}

func NewTransport(db *sqlx.DB) autogen.ServerInterface {
	return &Transport{
		srv: NewService(db),
	}
}

// Добавление нового юзера в БД
// (POST /api/seekers/add)
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

	if err := transport.srv.db.AddUser(user); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить пользователя")
		return
	}

	utils.WriteNoContent(w)
}

// Удаление пользователя из БД
// (DELETE /api/seekers/{chat_id}/remove)
func (transport *Transport) DeleteApiEmployersChatIdRemove(w http.ResponseWriter, r *http.Request, chatId int64) {
	ok, err := transport.srv.RemoveUser(chatId)
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

// Получение данных пользователя по chat_id
// (GET /api/seekers/{chat_id}/get)
func (transport *Transport) GetApiEmployersChatIdGet(w http.ResponseWriter, r *http.Request, chatId int64) {
	user, err := transport.srv.GetUserByChatID(chatId)
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

// Получение всех пользователей из БД
// (GET /api/seekers/get)
func (transport *Transport) GetApiEmployersGet(w http.ResponseWriter, r *http.Request) {
	users, err := transport.srv.GetAllUsers()
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить пользователей")
		return
	}
	if len(users) == 0 {
		utils.WriteString(w, http.StatusInternalServerError, err, "В базе нет пользователей")
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
		utils.WriteString(w, http.StatusOK, nil, "Невозможно обновить данные о пользователе")
		return
	} else {
		utils.WriteNoContent(w)
		return
	}
}
