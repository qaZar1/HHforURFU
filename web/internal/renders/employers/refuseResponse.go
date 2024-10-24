package renders

import (
	"net/http"
	"strconv"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
	"github.com/sirupsen/logrus"
)

func RefuseResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		logrus.Errorf("Error: invalid method")
		return
	}

	apiResponse := api.NewApiResponses()

	idStr := r.FormValue("id")
	if idStr == "" {
		logrus.Errorf("Error: invalid username")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Errorf("Error: invalid id")
		return
	}

	response := models.Response{
		Status: "Отказано",
	}

	err = apiResponse.EditResponse(int64(id), response)
	if err != nil {
		logrus.Errorf("Error - invalid edit response: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
