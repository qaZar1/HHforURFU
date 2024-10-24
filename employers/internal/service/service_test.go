package service_test

// import (
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/qaZar1/HHforURFU/employers/autogen"
// 	"github.com/qaZar1/HHforURFU/employers/internal/mocks"
// 	"github.com/qaZar1/HHforURFU/employers/internal/service"
// )

// func TestPostApiEmployersAdd(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	mockService.EXPECT().AddUser(gomock.Any()).Return(true, nil).Times(1)

// 	transport := service.NewTransport(mockService)

// 	user := autogen.User{
// 		ChatId:   123,
// 		Company:  "Company_1",
// 		Nickname: "user_1",
// 	}
// 	body, _ := json.Marshal(user)

// 	req := httptest.NewRequest(http.MethodPost, "/api/seekers/add", bytes.NewBuffer(body))
// 	w := httptest.NewRecorder()

// 	transport.PostApiEmployersAdd(w, req)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusNoContent {
// 		t.Errorf("expected status 204, got %d", res.StatusCode)
// 	}
// }

// func TestPostApiEmployersAdd_InvalidData(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	mockService.EXPECT().AddUser(gomock.Any()).Times(0)

// 	transport := service.NewTransport(mockService)

// 	// Пример некорректного JSON
// 	invalidUserJSON := []byte(`{"ChatId": "not_an_integer", "Company": 123, "Nickname": null}`)

// 	req := httptest.NewRequest(http.MethodPost, "/api/seekers/add", bytes.NewBuffer(invalidUserJSON))
// 	w := httptest.NewRecorder()

// 	transport.PostApiEmployersAdd(w, req)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusBadRequest {
// 		t.Errorf("expected status 400, got %d", res.StatusCode)
// 	}
// }

// func TestPostApiEmployersAdd_UserAlreadyExists(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	// Имитация ситуации, когда пользователь уже существует
// 	mockService.EXPECT().AddUser(gomock.Any()).Return(false, nil).Times(1)

// 	transport := service.NewTransport(mockService)

// 	user := autogen.User{
// 		ChatId:   123,
// 		Company:  "Company_1",
// 		Nickname: "user_1",
// 	}
// 	body, _ := json.Marshal(user)

// 	req := httptest.NewRequest(http.MethodPost, "/api/seekers/add", bytes.NewBuffer(body))
// 	w := httptest.NewRecorder()

// 	transport.PostApiEmployersAdd(w, req)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusConflict { // 409 Conflict
// 		t.Errorf("expected status 409, got %d", res.StatusCode)
// 	}
// }

// func TestPostApiEmployersAdd_InternalServerError(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	// Имитация внутренней ошибки сервиса
// 	mockService.EXPECT().AddUser(gomock.Any()).Return(false, fmt.Errorf("internal server error")).Times(1)

// 	transport := service.NewTransport(mockService)

// 	user := autogen.User{
// 		ChatId:   123,
// 		Company:  "Company_1",
// 		Nickname: "user_1",
// 	}
// 	body, _ := json.Marshal(user)

// 	req := httptest.NewRequest(http.MethodPost, "/api/seekers/add", bytes.NewBuffer(body))
// 	w := httptest.NewRecorder()

// 	transport.PostApiEmployersAdd(w, req)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusInternalServerError {
// 		t.Errorf("expected status 500, got %d", res.StatusCode)
// 	}
// }

// // Тесты для получения пользователя по chat_id
// func TestGetApiEmployersChatIdGet(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	expectedUser := &autogen.Info{
// 		ChatId:   123,
// 		Nickname: "user_1",
// 		Company:  "Company_1",
// 	}
// 	mockService.EXPECT().GetUserByChatID(int64(123)).Return(expectedUser, nil).Times(1)

// 	transport := service.NewTransport(mockService)

// 	req := httptest.NewRequest(http.MethodGet, "/api/seekers/123/get", nil)
// 	w := httptest.NewRecorder()

// 	transport.GetApiEmployersChatIdGet(w, req, 123)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("expected status 200, got %d", res.StatusCode)
// 	}

// 	var userResponse autogen.Info
// 	json.NewDecoder(res.Body).Decode(&userResponse)
// 	if userResponse != *expectedUser {
// 		t.Errorf("expected user %v, got %v", expectedUser, userResponse)
// 	}
// }

// func TestGetApiEmployersChatIdGet_UserNotFound(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	mockService.EXPECT().GetUserByChatID(int64(123)).Return(nil, sql.ErrNoRows).Times(1)

// 	transport := service.NewTransport(mockService)

// 	req := httptest.NewRequest(http.MethodGet, "/api/seekers/123/get", nil)
// 	w := httptest.NewRecorder()

// 	transport.GetApiEmployersChatIdGet(w, req, 123)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusNoContent {
// 		t.Errorf("expected status 204, got %d", res.StatusCode)
// 	}
// }

// // Тесты для удаления пользователя
// func TestDeleteApiEmployersChatIdRemove(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	mockService.EXPECT().RemoveUser(int64(123)).Return(true, nil).Times(1)

// 	transport := service.NewTransport(mockService)

// 	req := httptest.NewRequest(http.MethodDelete, "/api/seekers/123/remove", nil)
// 	w := httptest.NewRecorder()

// 	transport.DeleteApiEmployersChatIdRemove(w, req, 123)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusNoContent {
// 		t.Errorf("expected status 204, got %d", res.StatusCode)
// 	}
// }

// func TestDeleteApiEmployersChatIdRemove_UserNotFound(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	mockService.EXPECT().RemoveUser(int64(123)).Return(false, sql.ErrNoRows).Times(1)

// 	transport := service.NewTransport(mockService)

// 	req := httptest.NewRequest(http.MethodDelete, "/api/seekers/123/remove", nil)
// 	w := httptest.NewRecorder()

// 	transport.DeleteApiEmployersChatIdRemove(w, req, 123)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusNotFound {
// 		t.Errorf("expected status 404, got %d", res.StatusCode)
// 	}
// }

// // Тесты для обновления пользователя
// func TestPutApiEmployersChatIdUpdate(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	mockService.EXPECT().UpdateUser(int64(123), gomock.Any()).Return(true, nil).Times(1)

// 	transport := service.NewTransport(mockService)

// 	updateUser := autogen.UpdateUser{
// 		Nickname: "updated_user",
// 		Company:  "Updated_Company",
// 	}
// 	body, _ := json.Marshal(updateUser)

// 	req := httptest.NewRequest(http.MethodPut, "/api/seekers/123/update", bytes.NewBuffer(body))
// 	w := httptest.NewRecorder()

// 	transport.PutApiEmployersChatIdUpdate(w, req, 123)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusNoContent {
// 		t.Errorf("expected status 204, got %d", res.StatusCode)
// 	}
// }

// func TestPutApiEmployersChatIdUpdate_UserNotFound(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	mockService.EXPECT().UpdateUser(int64(123), gomock.Any()).Return(false, nil).Times(1)

// 	transport := service.NewTransport(mockService)

// 	updateUser := autogen.UpdateUser{
// 		Nickname: "updated_user",
// 		Company:  "Updated_Company",
// 	}
// 	body, _ := json.Marshal(updateUser)

// 	req := httptest.NewRequest(http.MethodPut, "/api/seekers/123/update", bytes.NewBuffer(body))
// 	w := httptest.NewRecorder()

// 	transport.PutApiEmployersChatIdUpdate(w, req, 123)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusNotFound {
// 		t.Errorf("expected status 404, got %d", res.StatusCode)
// 	}
// }

// // Тесты для получения всех пользователей
// func TestGetApiEmployersGet(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	expectedUsers := []autogen.Info{
// 		{ChatId: 123, Nickname: "user_1", Company: "Company_1"},
// 		{ChatId: 456, Nickname: "user_2", Company: "Company_2"},
// 	}
// 	mockService.EXPECT().GetAllUsers().Return(expectedUsers, nil).Times(1)

// 	transport := service.NewTransport(mockService)

// 	req := httptest.NewRequest(http.MethodGet, "/api/seekers/get", nil)
// 	w := httptest.NewRecorder()

// 	transport.GetApiEmployersGet(w, req)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("expected status 200, got %d", res.StatusCode)
// 	}

// 	var usersResponse []autogen.Info
// 	json.NewDecoder(res.Body).Decode(&usersResponse)
// 	if len(usersResponse) != len(expectedUsers) {
// 		t.Errorf("expected %d users, got %d", len(expectedUsers), len(usersResponse))
// 	}
// }

// func TestGetApiEmployersGet_NoUsers(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockServiceInterface(ctrl)

// 	mockService.EXPECT().GetAllUsers().Return([]autogen.Info{}, nil).Times(1)

// 	transport := service.NewTransport(mockService)

// 	req := httptest.NewRequest(http.MethodGet, "/api/seekers/get", nil)
// 	w := httptest.NewRecorder()

// 	transport.GetApiEmployersGet(w, req)

// 	res := w.Result()
// 	if res.StatusCode != http.StatusNoContent {
// 		t.Errorf("expected status 204, got %d", res.StatusCode)
// 	}

// 	var usersResponse []autogen.Info
// 	json.NewDecoder(res.Body).Decode(&usersResponse)
// 	if len(usersResponse) != 0 {
// 		t.Errorf("expected 0 users, got %d", len(usersResponse))
// 	}
// }
