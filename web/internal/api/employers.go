package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type APIEmployers struct {
	client *resty.Client
}

func NewApiEmployers() *APIEmployers {
	return &APIEmployers{
		client: resty.New().SetBaseURL("http://localhost:8001/api").SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APIEmployers) RegisterEmployer(employer models.Employer) (bool, error) {
	const endpoint = "/registerEmployer"

	resp, err := api.client.R().SetBody(employer).Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, fmt.Errorf("failed to add employer: %s", resp.Status())
	}

	return true, nil
}

func (api *APIEmployers) CheckEmployer(employerID int64) (models.Employer, error) {
	endpoint := "/checkEmployer/%d"

	endpoint = fmt.Sprintf(endpoint, employerID)
	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return models.Employer{}, err
	}

	employer := models.Employer{}
	if err := jsoniter.Unmarshal(resp.Body(), &employer); err != nil {
		return models.Employer{}, err
	}

	return employer, nil
}

func (api *APIEmployers) LoginEmployer() (models.TokenResponse, error) {
	const endpoint = "/loginEmployer"

	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return models.TokenResponse{}, err
	}

	token := models.TokenResponse{}
	if err := jsoniter.Unmarshal(resp.Body(), &token); err != nil {
		return models.TokenResponse{}, err
	}

	return token, nil
}

func (api *APIEmployers) CheckEmployerByUsername(username string) (models.Employer, error) {
	const endpoint = "/checkEmployerByUsername/%s"

	endpointWithUsername := fmt.Sprintf(endpoint, username)
	resp, err := api.client.R().Get(endpointWithUsername)
	if err != nil {
		return models.Employer{}, err
	}

	employer := models.Employer{}
	if err := jsoniter.Unmarshal(resp.Body(), &employer); err != nil {
		return models.Employer{}, err
	}

	return employer, nil
}
