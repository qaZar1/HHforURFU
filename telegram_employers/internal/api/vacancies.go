package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/telegram_employers/internal/models"
)

type APIVacancies struct {
	client *resty.Client
}

func NewApiVacancies(url string) *APIVacancies {
	return &APIVacancies{
		client: resty.New().SetBaseURL(url).SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APIVacancies) GetAllVacancies() ([]models.Vacancies, error) {
	const endpoint = "/vacancies/get"

	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	resp.Body()
	allVacancies := []models.Vacancies{}
	if err := jsoniter.Unmarshal(resp.Body(), &allVacancies); err != nil {
		return nil, err
	}

	return allVacancies, nil
}

func (api *APIVacancies) AddVacancy(vacancy models.Vacancies) (bool, int64, error) {
	const endpoint = "/vacancies/add"

	resp, err := api.client.R().SetBody(vacancy).Post(endpoint)
	if err != nil {
		return false, 0, err
	}

	if resp.IsError() {
		return false, 0, fmt.Errorf("failed to add seeker: %s", resp.Status())
	}

	result_vacancy := models.Vacancies{}
	if err := jsoniter.Unmarshal(resp.Body(), &result_vacancy); err != nil {
		return false, 0, err
	}

	return true, result_vacancy.Vacancy_ID, nil
}

func (api *APIVacancies) GetVacancyByVacancyID(vacancyId int64) (models.Vacancies, error) {
	endpoint := "/vacancies/%d/get"

	resp, err := api.client.R().SetBasicAuth("dev", "test").Get(fmt.Sprintf(endpoint, vacancyId))
	if err != nil {
		return models.Vacancies{}, err
	}

	vacancy := models.Vacancies{}
	if err := jsoniter.Unmarshal(resp.Body(), &vacancy); err != nil {
		return models.Vacancies{}, err
	}

	return vacancy, nil
}
