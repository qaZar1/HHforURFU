package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type APISeekers struct {
	client *resty.Client
}

func NewApiSeekers() *APISeekers {
	return &APISeekers{
		client: resty.New().SetBaseURL("http://localhost:8005/api").SetTimeout(1*time.Minute).SetBasicAuth("dev", "test"),
	}
}

func (api *APISeekers) RegisterSeeker(seeker models.SeekerWithoutID) (bool, error) {
	const endpoint = "/registerSeeker"

	resp, err := api.client.R().SetBody(seeker).Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, fmt.Errorf("failed to add seeker: %s", resp.Status())
	}

	return true, nil
}

func (api *APISeekers) CheckSeeker(username string) (models.Seeker, error) {
	endpoint := "/checkSeeker/%s"

	endpoint = fmt.Sprintf(endpoint, username)
	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return models.Seeker{}, err
	}

	seeker := models.Seeker{}
	if err := jsoniter.Unmarshal(resp.Body(), &seeker); err != nil {
		return models.Seeker{}, err
	}

	return seeker, nil
}

func (api *APISeekers) Login() (models.TokenResponse, error) {
	const endpoint = "/login"

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
