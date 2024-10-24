package service

import "github.com/qaZar1/HHforURFU/responses/internal/models"

type ServiceInterface interface {
	GetResponsesByUsername(username string) ([]models.Response, error)
	GetResponseByID(id int64) (models.Response, error)
	GetResponsesByEmployersID(id int64) ([]models.Response, error)
	AddResponses(respons models.Response) error
	UpdateResponse(response_id int, updateRespons models.Response) (bool, error)
}
