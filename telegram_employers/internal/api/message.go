package api

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/qaZar1/HHforURFU/telegram_employers/internal/models"
)

type APIMessage struct {
	client *resty.Client
}

func NewApiMessage(url string) *APIMessage {
	return &APIMessage{
		client: resty.New().SetBaseURL(url).SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APIMessage) SendMsg(body models.Notify) (bool, error) {
	endpoint := "/notify"

	resp, err := api.client.R().SetBody(body).SetBasicAuth("dev", "test").Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, err
	}

	return true, nil
}
