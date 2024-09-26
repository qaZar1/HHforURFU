package service_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qaZar1/HHforURFU/responses/autogen"
	"github.com/qaZar1/HHforURFU/responses/internal/mocks"
	"github.com/qaZar1/HHforURFU/responses/internal/service"
)

func TestGetApiResponsesGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	expectedResponses := []autogen.Response{{ /* инициализация полей */ }}
	mockService.EXPECT().GetAllResponses().Return(expectedResponses, nil).Times(1)

	transport := service.NewTransport(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/responses/get", nil)
	w := httptest.NewRecorder()

	transport.GetApiResponsesGet(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

	var responses []autogen.Response
	json.NewDecoder(res.Body).Decode(&responses)
	if len(responses) != len(expectedResponses) {
		t.Errorf("expected %d responses, got %d", len(expectedResponses), len(responses))
	}
}

func TestGetApiResponsesGet_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	mockService.EXPECT().GetAllResponses().Return(nil, fmt.Errorf("error")).Times(1)

	transport := service.NewTransport(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/responses/get", nil)
	w := httptest.NewRecorder()

	transport.GetApiResponsesGet(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", res.StatusCode)
	}
}

func TestPostApiResponsesAdd_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	newResponse := autogen.Response{ /* инициализация полей */ }
	mockService.EXPECT().AddResponses(newResponse).Return(nil).Times(1)

	transport := service.NewTransport(mockService)

	body, _ := json.Marshal(newResponse)
	req := httptest.NewRequest(http.MethodPost, "/api/responses/add", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	transport.PostApiResponsesAdd(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", res.StatusCode)
	}
}

func TestPostApiResponsesAdd_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	newResponse := autogen.Response{ /* инициализация полей */ }
	mockService.EXPECT().AddResponses(newResponse).Return(fmt.Errorf("error")).Times(1)

	transport := service.NewTransport(mockService)

	body, _ := json.Marshal(newResponse)
	req := httptest.NewRequest(http.MethodPost, "/api/responses/add", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	transport.PostApiResponsesAdd(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", res.StatusCode)
	}
}

func TestDeleteApiResponsesRemove_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	mockService.EXPECT().RemoveResponses(int64(1)).Return(false, nil).Times(1)

	transport := service.NewTransport(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/api/responses/1/remove", nil)
	w := httptest.NewRecorder()

	transport.DeleteApiResponsesVacancyIdRemove(w, req, int64(1))

	res := w.Result()
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", res.StatusCode)
	}
}

func TestDeleteApiResponsesRemove_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	mockService.EXPECT().RemoveResponses(int64(1)).Return(false, sql.ErrNoRows).Times(1)

	transport := service.NewTransport(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/api/responses/1/remove", nil)
	w := httptest.NewRecorder()

	transport.DeleteApiResponsesVacancyIdRemove(w, req, 1)

	res := w.Result()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", res.StatusCode)
	}
}

func TestPutApiResponsesUpdate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	updatedResponse := autogen.Response{
		ChatId:         1,
		ChatIdEmployer: 2,
		Status:         "Успешно",
		VacancyId:      3,
	}
	mockService.EXPECT().UpdateRespons(int64(1), updatedResponse).Return(false, nil).Times(1)

	transport := service.NewTransport(mockService)

	body, _ := json.Marshal(updatedResponse)
	req := httptest.NewRequest(http.MethodPut, "/api/responses/1/update", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	transport.PutApiResponsesVacancyIdUpdate(w, req, int64(1))

	res := w.Result()
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", res.StatusCode)
	}
}

func TestPutApiResponsesUpdate_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	updatedResponse := autogen.Response{
		ChatId:         1,
		ChatIdEmployer: 2,
		Status:         "Успешно",
		VacancyId:      3,
	}
	mockService.EXPECT().UpdateRespons(int64(1), updatedResponse).Return(true, nil).Times(1)

	transport := service.NewTransport(mockService)

	body, _ := json.Marshal(updatedResponse)
	req := httptest.NewRequest(http.MethodPut, "/api/responses/1/update", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	transport.PutApiResponsesVacancyIdUpdate(w, req, int64(1))

	res := w.Result()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", res.StatusCode)
	}
}
