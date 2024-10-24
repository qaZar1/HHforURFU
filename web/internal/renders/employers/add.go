package renders

import (
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
	"github.com/sirupsen/logrus"
)

func AddVacancyInList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logrus.Errorf("Error: invalid method")
		return
	}

	apiVacancies := api.NewApiVacancies()
	apiFilters := api.NewApiFilters()
	apiEmployers := api.NewApiEmployers()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Error: invalid read body")
		return
	}

	var temp models.Temp
	if err := jsoniter.Unmarshal(data, &temp); err != nil {
		logrus.Errorf("Error: invalid unmarshal body")
		return
	}

	employer, err := apiEmployers.CheckEmployerByUsername(temp.Username)
	if err != nil {
		logrus.Errorf("Error: invalid check employer")
		return
	}

	vacancy := models.AddVacancy{
		Company:     employer.Company,
		Title:       temp.Title,
		Description: temp.Description,
		Status:      "Открыта",
		EmployerID:  employer.EmployerID,
	}

	id, err := apiVacancies.AddVacancy(vacancy)
	if err != nil {
		logrus.Errorf("Error: invalid add vacancy")
		return
	}

	for _, tag := range temp.Tags {
		filter := models.Tag{
			VacancyID: id,
			Tag:       tag,
		}
		_, err := apiFilters.AddFilters(filter)
		if err != nil {
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
