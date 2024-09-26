package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/qaZar1/HHforURFU/telegram_employers/internal/models"
)

type APItelegram struct {
	client *resty.Client
}

func NewApiTelegram(url string) *APItelegram {
	return &APItelegram{
		client: resty.New().SetBaseURL(url).SetTimeout(1 * time.Minute),
	}
}

func (api *APItelegram) AddEmployer(body models.SendMessageRequest) (bool, error) {
	const endpoint = "/sendMessage"

	resp, err := api.client.R().SetBody(body).Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, fmt.Errorf("failed to add seeker: %s", resp.Status())
	}

	return true, nil
}
