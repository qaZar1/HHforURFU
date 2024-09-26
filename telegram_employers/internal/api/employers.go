package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/telegram_employers/internal/models"
)

type APIEmployers struct {
	client *resty.Client
}

func NewApiEmployers(url string) *APIEmployers {
	return &APIEmployers{
		client: resty.New().SetBaseURL(url).SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APIEmployers) GetAllEmployers() ([]models.Employers, error) {
	const endpoint = "/employers/get"

	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	resp.Body()
	allEmployers := []models.Employers{}
	if err := jsoniter.Unmarshal(resp.Body(), &allEmployers); err != nil {
		return nil, err
	}

	return allEmployers, nil
}

func (api *APIEmployers) AddEmployer(employer models.Employers) (bool, error) {
	const endpoint = "/employers/add"

	resp, err := api.client.R().SetBody(employer).Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, fmt.Errorf("failed to add seeker: %s", resp.Status())
	}

	return true, nil
}

func (api *APIEmployers) GetEmployerByChatID(chatId int64) (models.Employers, bool, error) {
	endpoint := "/employers/%d/get"

	resp, err := api.client.R().SetBasicAuth("dev", "test").Get(fmt.Sprintf(endpoint, chatId))
	if err != nil {
		return models.Employers{}, false, err
	}

	employers := models.Employers{}
	if err := jsoniter.Unmarshal(resp.Body(), &employers); err != nil {
		return models.Employers{}, false, err
	}

	return employers, true, nil
}
