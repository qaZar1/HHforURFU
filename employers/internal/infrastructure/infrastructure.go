package infrastructure

import (
	"os"

	"github.com/qaZar1/HHforURFU/employers/internal/models"
)

type Infrastructure struct {
	ExpiresIn int

	secret []byte

	issuer    string
	resources map[string]models.Roles
}

func New(expiresIn int) *Infrastructure {
	return &Infrastructure{
		secret:    []byte(os.Getenv(envtSecret)),
		issuer:    issuer,
		ExpiresIn: expiresIn,
		resources: map[string]models.Roles{
			resourceHelloWorld: {
				Roles: []string{
					permGet + resourceHelloWorld,
				},
			},
			resourcePingPong: {
				Roles: []string{
					permGet + resourcePingPong,
				},
			},
		},
	}
}
