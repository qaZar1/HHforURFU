package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/telegram_seekers/internal/models"
)

type APIResponses struct {
	client *resty.Client
}

func NewApiResponses(url string) *APIResponses {
	return &APIResponses{
		client: resty.New().SetBaseURL(url).SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APIResponses) GetAllResponses() ([]models.Responses, error) {
	const endpoint = "/responses/get"

	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	resp.Body()
	allResponses := []models.Responses{}
	if err := jsoniter.Unmarshal(resp.Body(), &allResponses); err != nil {
		return nil, err
	}

	return allResponses, nil
}

func (api *APIResponses) AddResponses(response models.Responses) (bool, error) {
	const endpoint = "/responses/add"

	resp, err := api.client.R().SetBody(response).Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, fmt.Errorf("failed to add seeker: %s", resp.Status())
	}

	return true, nil
}

func (api *APIResponses) GetResponsesByVacancyIDAndChatID(vacancyId int64, chatId int64) (bool, error) {
	endpoint := "/responses/%d-and-%d/get"

	resp, err := api.client.R().SetBasicAuth("dev", "test").Get(fmt.Sprintf(endpoint, vacancyId, chatId))
	if err != nil {
		return false, err
	}

	user := models.Responses{}
	if err := jsoniter.Unmarshal(resp.Body(), &user); err != nil {
		return false, err
	}

	return true, nil
}

func (api *APIResponses) GetResponsesByChatID(chatId int64) ([]models.Responses, error) {
	endpoint := "/responses/chat_id/%d/get"

	resp, err := api.client.R().SetBasicAuth("dev", "test").Get(fmt.Sprintf(endpoint, chatId))
	if err != nil {
		return []models.Responses{}, err
	}

	responses := []models.Responses{}
	if err := jsoniter.Unmarshal(resp.Body(), &responses); err != nil {
		return []models.Responses{}, err
	}

	return responses, nil
}
