package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/telegram_employers/internal/models"
)

type APITags struct {
	client *resty.Client
}

func NewApiTags(url string) *APITags {
	return &APITags{
		client: resty.New().SetBaseURL(url).SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APITags) GetAllFilters() ([]models.Filters, error) {
	const endpoint = "/filters/get"

	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	resp.Body()
	allFilters := []models.Filters{}
	if err := jsoniter.Unmarshal(resp.Body(), &allFilters); err != nil {
		return nil, err
	}

	return allFilters, nil
}

func (api *APITags) AddFilters(filters models.Filters) (bool, error) {
	const endpoint = "/filters/add"

	resp, err := api.client.R().SetBody(filters).Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, fmt.Errorf("failed to add seeker: %s", resp.Status())
	}

	return true, nil
}

func (api *APITags) GetFiltersByVacancyID(vacancyId int64) (bool, error) {
	endpoint := "/filters/%d/get"

	resp, err := api.client.R().SetBasicAuth("dev", "test").Get(fmt.Sprintf(endpoint, vacancyId))
	if err != nil {
		return false, err
	}

	user := models.Seekers{}
	if err := jsoniter.Unmarshal(resp.Body(), &user); err != nil {
		return false, err
	}

	return true, nil
}
