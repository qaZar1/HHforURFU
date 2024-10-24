package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/web/internal/models"
	"github.com/sirupsen/logrus"
)

type APIVacancies struct {
	client *resty.Client
}

func NewApiVacancies() *APIVacancies {
	return &APIVacancies{
		client: resty.New().SetBaseURL("http://localhost:8002/api").SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APIVacancies) GetAllVacancies() ([]models.Vacancy, error) {
	const endpoint = "/vacancies/getAllVacancy"

	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	allVacancies := []models.Vacancy{}
	if err := jsoniter.Unmarshal(resp.Body(), &allVacancies); err != nil {
		return nil, err
	}

	return allVacancies, nil
}

func (api *APIVacancies) GetVacancyByVacancyID(id int64) (models.Vacancy, error) {
	endpoint := "/vacancies/getVacancyByVacancyID/%d"

	endpoint = fmt.Sprintf(endpoint, id)
	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return models.Vacancy{}, err
	}

	vacancy := models.Vacancy{}
	if err := jsoniter.Unmarshal(resp.Body(), &vacancy); err != nil {
		return models.Vacancy{}, err
	}

	return vacancy, nil
}

func (api *APIVacancies) GetAllVacanciesByEmployerID(employerID int64) ([]models.Vacancy, error) {
	const endpoint = "/vacancies/getVacancyByEmployerID/%d"

	ep := fmt.Sprintf(endpoint, employerID)
	resp, err := api.client.R().Get(ep)
	if err != nil {
		return nil, err
	}

	allVacancies := []models.Vacancy{}
	if err := jsoniter.Unmarshal(resp.Body(), &allVacancies); err != nil {
		return nil, err
	}

	return allVacancies, nil
}

func (api *APIVacancies) AddVacancy(vacancy models.AddVacancy) (int64, error) {
	const endpoint = "/vacancies/addVacancy"

	resp, err := api.client.R().SetBody(vacancy).Post(endpoint)
	if err != nil {
		logrus.Errorf("Invalid add vacancy: %s", err)
		return 0, err
	}

	if resp.IsError() {
		logrus.Errorf("Invalid add vacancy: %s", err)
		return 0, err
	}

	var a models.AddVacancy
	if err := jsoniter.Unmarshal(resp.Body(), &a); err != nil {
		logrus.Errorf("Invalid add vacancy: %s", err)
		return 0, err
	}

	return a.ID, nil
}

func (api *APIVacancies) DeleteVacancy(vacancy_id int64) error {
	endpoint := "/vacancies/delete/%d"

	endpoint = fmt.Sprintf(endpoint, vacancy_id)
	resp, err := api.client.R().Delete(endpoint)
	if err != nil {
		logrus.Errorf("Invalid add vacancy: %s", err)
		return err
	}

	if resp.IsError() {
		logrus.Errorf("Invalid add vacancy: %s", err)
		return err
	}

	return nil
}
