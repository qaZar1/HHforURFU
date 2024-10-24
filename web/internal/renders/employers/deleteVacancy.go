package renders

import (
	"net/http"
	"strconv"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/sirupsen/logrus"
)

func DeleteVacancy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		logrus.Errorf("Error: invalid method")
		return
	}

	apiVacancies := api.NewApiVacancies()

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

	err = apiVacancies.DeleteVacancy(int64(id))
	if err != nil {
		logrus.Errorf("Error - invalid edit response: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
