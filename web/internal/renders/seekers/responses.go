package renders

import (
	"html/template"
	"net/http"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type Responses struct {
	Responses []models.ResponseWithVacancy
}

func renderResponses(w http.ResponseWriter, templateName string, data Responses) {
	tmpl, err := template.ParseFiles("internal/temp/" + templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data) // Здесь можно передать данные в шаблон, если нужно
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderResp(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		return
	}

	apiResponses := api.NewApiResponses()
	apiVacancy := api.NewApiVacancies()
	responses, err := apiResponses.GetAllResponsesByUsername(cookie.Value)
	if err != nil {
		return
	}

	responsesWithVacancy := []models.ResponseWithVacancy{}
	for _, resp := range responses {
		vacancy, err := apiVacancy.GetVacancyByVacancyID(resp.VacancyID)
		if err != nil {
			continue
		}

		response := models.ResponseWithVacancy{
			RespID:      resp.RespID,
			VacancyID:   resp.VacancyID,
			EmployerID:  resp.EmployerID,
			Username:    resp.Username,
			Status:      resp.Status,
			Company:     vacancy.Company,
			Title:       vacancy.Title,
			Description: vacancy.Description,
			Tags:        vacancy.Tags,
		}

		responsesWithVacancy = append(responsesWithVacancy, response)
	}

	data := Responses{
		Responses: responsesWithVacancy,
	}
	renderResponses(w, "seekers/responses.html", data)
}
