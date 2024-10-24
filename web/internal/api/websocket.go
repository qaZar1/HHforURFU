package api

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type APIWebSocket struct {
	client *resty.Client
}

func NewApiWebSocket() *APIWebSocket {
	return &APIWebSocket{
		client: resty.New().SetBaseURL("http://localhost:8006/api").SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APIWebSocket) AddResponse(message models.Message) error {
	const endpoint = "/websocket/addMessage"

	resp, err := api.client.R().SetBody(message).Post(endpoint)
	if err != nil {
		return err
	}

	if resp.IsError() {
		return errors.New("Error adding message")
	}

	return nil
}
