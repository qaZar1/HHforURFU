package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type APIFilters struct {
	client *resty.Client
}

func NewApiFilters() *APIFilters {
	return &APIFilters{
		client: resty.New().SetBaseURL("http://localhost:8004/api").SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APIFilters) GetTagsByVacancyID(vacancy_id int64) ([]models.Tag, error) {
	endpoint := "/filters/getTagsByVacancyID/%d"

	endpoint = fmt.Sprintf(endpoint, vacancy_id)
	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return []models.Tag{}, err
	}

	tags := []models.Tag{}
	if err := jsoniter.Unmarshal(resp.Body(), &tags); err != nil {
		return []models.Tag{}, err
	}

	return tags, nil
}

func (api *APIFilters) GetResponsesByID(id int64) (models.Response, error) {
	endpoint := "/responses/getResponseByResponseID/%d"

	endpoint = fmt.Sprintf(endpoint, id)
	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return models.Response{}, err
	}

	responses := models.Response{}
	if err := jsoniter.Unmarshal(resp.Body(), &responses); err != nil {
		return models.Response{}, err
	}

	return responses, nil
}

func (api *APIFilters) AddFilters(filter models.Tag) (bool, error) {
	const endpoint = "/filters/addFilter"

	resp, err := api.client.R().SetBody(filter).Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, err
	}
	return true, nil
}
