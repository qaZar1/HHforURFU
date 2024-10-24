package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type APIResponses struct {
	client *resty.Client
}

func NewApiResponses() *APIResponses {
	return &APIResponses{
		client: resty.New().SetBaseURL("http://localhost:8003/api").SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APIResponses) GetAllResponsesByUsername(username string) ([]models.Response, error) {
	endpoint := "/responses/getResponsesByUsername/%s"

	endpoint = fmt.Sprintf(endpoint, username)
	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return []models.Response{}, err
	}

	responses := []models.Response{}
	if err := jsoniter.Unmarshal(resp.Body(), &responses); err != nil {
		return []models.Response{}, err
	}

	return responses, nil
}

func (api *APIResponses) GetResponsesByID(id int64) (models.Response, error) {
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

func (api *APIResponses) GetAllResponsesByEmployerID(id int64) ([]models.Response, error) {
	endpoint := "/responses/getResponsesByEmployersID/%d"

	endpoint = fmt.Sprintf(endpoint, id)
	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return []models.Response{}, err
	}

	responses := []models.Response{}
	if err := jsoniter.Unmarshal(resp.Body(), &responses); err != nil {
		return []models.Response{}, err
	}

	return responses, nil
}

func (api *APIResponses) AddResponse(response models.Response) error {
	const endpoint = "/responses/addResponse"

	resp, err := api.client.R().SetBody(response).Post(endpoint)
	if err != nil {
		return err
	}

	if resp.IsError() {
		return errors.New("Error adding response")
	}

	return nil
}

func (api *APIResponses) EditResponse(response_id int64, response models.Response) error {
	const endpointWithoutResponseID = "/responses/edit/%d"

	endpoint := fmt.Sprintf(endpointWithoutResponseID, response_id)
	resp, err := api.client.R().SetBody(response).Put(endpoint)
	if err != nil {
		return err
	}

	if resp.IsError() {
		return errors.New("Error put response")
	}

	return nil
}
