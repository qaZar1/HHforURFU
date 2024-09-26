package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/telegram_employers/internal/models"
)

type APISeekers struct {
	client *resty.Client
}

func NewApiSeekers(url string) *APISeekers {
	return &APISeekers{
		client: resty.New().SetBaseURL(url).SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APISeekers) GetAllSeekers() ([]models.Seekers, error) {
	const endpoint = "/seekers/get"

	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	resp.Body()
	allUsers := []models.Seekers{}
	if err := jsoniter.Unmarshal(resp.Body(), &allUsers); err != nil {
		return nil, err
	}

	return allUsers, nil
}

func (api *APISeekers) AddSeeker(user models.Seekers) (bool, error) {
	const endpoint = "/seekers/add"

	resp, err := api.client.R().SetBody(user).Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, fmt.Errorf("failed to add seeker: %s", resp.Status())
	}

	return true, nil
}

func (api *APISeekers) GetSeekerByChatID(chatId int64) (models.Seekers, error) {
	endpoint := "/seekers/%d/get"

	resp, err := api.client.R().SetBasicAuth("dev", "test").Get(fmt.Sprintf(endpoint, chatId))
	if err != nil {
		return models.Seekers{}, err
	}

	user := models.Seekers{}
	if err := jsoniter.Unmarshal(resp.Body(), &user); err != nil {
		return models.Seekers{}, err
	}

	return user, nil
}
