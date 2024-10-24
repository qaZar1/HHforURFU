package renders

import (
	"net/http"
	"strconv"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
	"github.com/sirupsen/logrus"
)

func Accept(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logrus.Errorf("Error: invalid method")
		return
	}

	apiResponses := api.NewApiResponses()

	username := r.FormValue("username")
	if username == "undefined" {
		logrus.Errorf("Error: invalid username")
		return
	}

	vacancy_id := r.FormValue("vacancyID")
	if vacancy_id == "" {
		logrus.Errorf("Error: invalid vacancyID")
		return
	}

	vacancyID, err := strconv.Atoi(vacancy_id)
	if err != nil {
		logrus.Errorf("Error: invalid vacancyID")
	}

	employer_id := r.FormValue("employerID")
	if employer_id == "" {
		logrus.Errorf("Error: invalid employerID")
		return
	}

	employerID, err := strconv.Atoi(employer_id)
	if err != nil {
		logrus.Errorf("Error: invalid employerID")
	}

	response := models.Response{
		Username:   username,
		VacancyID:  int64(vacancyID),
		EmployerID: int64(employerID),
		Status:     "Ожидает ответа",
	}

	if err := apiResponses.AddResponse(response); err != nil {
		logrus.Errorf("Invalid add response: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
