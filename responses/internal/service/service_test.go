package service_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qaZar1/HHforURFU/responses/internal/mocks"
	"github.com/qaZar1/HHforURFU/responses/internal/models"
	assert "github.com/stretchr/testify/assert"
)

func TestGetResponsesByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)
	username := "test_user"

	expectedResponses := []models.Response{
		{
			RespID:     1,
			VacancyID:  2,
			EmployerID: 3,
			Username:   "test_user",
			Status:     "status",
		},
	}

	t.Run("Success", func(t *testing.T) {
		// Определяем поведение мока
		mockService.EXPECT().
			GetResponsesByUsername(username).
			Return(expectedResponses, nil)

		// Вызываем метод
		resp, err := mockService.GetResponsesByUsername(username)

		// Проверяем результат
		assert.NoError(t, err)
		assert.Equal(t, expectedResponses, resp)
	})

	t.Run("No Responses", func(t *testing.T) {
		mockService.EXPECT().
			GetResponsesByUsername(username).
			Return([]models.Response{}, nil)

		resp, err := mockService.GetResponsesByUsername(username)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(resp))
	})

	t.Run("Error", func(t *testing.T) {
		mockService.EXPECT().
			GetResponsesByUsername(username).
			Return(nil, errors.New("some error"))

		resp, err := mockService.GetResponsesByUsername(username)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestAddResponses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)
	newResponse := models.Response{
		RespID:     1,
		VacancyID:  2,
		EmployerID: 3,
		Username:   "test_user",
		Status:     "status",
	}

	t.Run("Success", func(t *testing.T) {
		mockService.EXPECT().
			AddResponses(newResponse).
			Return(nil)

		err := mockService.AddResponses(newResponse)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		mockService.EXPECT().
			AddResponses(newResponse).
			Return(errors.New("insert error"))

		err := mockService.AddResponses(newResponse)
		assert.Error(t, err)
	})
}

func TestGetResponseByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)
	responseID := int64(1)

	expectedResponse := models.Response{
		RespID:     1,
		VacancyID:  2,
		EmployerID: 3,
		Username:   "test_user",
		Status:     "status",
	}

	t.Run("Success", func(t *testing.T) {
		mockService.EXPECT().
			GetResponseByID(responseID).
			Return(expectedResponse, nil)

		resp, err := mockService.GetResponseByID(responseID)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, resp)
	})

	t.Run("Not Found", func(t *testing.T) {
		mockService.EXPECT().
			GetResponseByID(responseID).
			Return(models.Response{}, sql.ErrNoRows)

		resp, err := mockService.GetResponseByID(responseID)
		assert.Error(t, err)
		assert.Equal(t, models.Response{}, resp)
	})

	t.Run("Error", func(t *testing.T) {
		mockService.EXPECT().
			GetResponseByID(responseID).
			Return(models.Response{}, errors.New("some error"))

		resp, err := mockService.GetResponseByID(responseID)
		assert.Error(t, err)
		assert.Equal(t, models.Response{}, resp)
	})
}

func TestGetResponsesByEmployersID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)
	employerID := int64(10)

	expectedResponses := []models.Response{
		{
			RespID:     1,
			VacancyID:  2,
			EmployerID: 3,
			Username:   "test_user",
			Status:     "status",
		},
		{
			RespID:     2,
			VacancyID:  3,
			EmployerID: 4,
			Username:   "test_user",
			Status:     "status",
		},
	}

	t.Run("Success", func(t *testing.T) {
		mockService.EXPECT().
			GetResponsesByEmployersID(employerID).
			Return(expectedResponses, nil)

		resp, err := mockService.GetResponsesByEmployersID(employerID)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponses, resp)
	})

	t.Run("Not Found", func(t *testing.T) {
		mockService.EXPECT().
			GetResponsesByEmployersID(employerID).
			Return([]models.Response{}, sql.ErrNoRows)

		resp, err := mockService.GetResponsesByEmployersID(employerID)
		assert.Error(t, err)
		assert.Equal(t, 0, len(resp))
	})

	t.Run("Error", func(t *testing.T) {
		mockService.EXPECT().
			GetResponsesByEmployersID(employerID).
			Return(nil, errors.New("some error"))

		resp, err := mockService.GetResponsesByEmployersID(employerID)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestUpdateResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	responseID := 1
	updateResponse := models.Response{
		RespID:     2,
		VacancyID:  3,
		EmployerID: 4,
		Username:   "test_user",
		Status:     "status",
	}

	mockService.EXPECT().UpdateResponse(responseID, updateResponse).Return(true, nil)

	ok, err := mockService.UpdateResponse(responseID, updateResponse)
	assert.NoError(t, err)
	assert.True(t, ok)

	mockService.EXPECT().UpdateResponse(responseID, updateResponse).Return(false, nil)

	ok, err = mockService.UpdateResponse(responseID, updateResponse)
	assert.NoError(t, err)
	assert.False(t, ok)

	mockService.EXPECT().UpdateResponse(responseID, updateResponse).Return(true, errors.New("db error"))

	ok, err = mockService.UpdateResponse(responseID, updateResponse)
	assert.Error(t, err)
	assert.False(t, ok)
}
