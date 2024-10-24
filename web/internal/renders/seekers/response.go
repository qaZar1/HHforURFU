package renders

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type Response struct {
	Response models.Response
	Vacancy  models.Vacancy
	Seeker   models.Seeker
}

func renderResponse(w http.ResponseWriter, templateName string, data Response) {
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

func RenderResponse(w http.ResponseWriter, r *http.Request) {
	id_str := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	apiVacancies := api.NewApiVacancies()
	apiResponses := api.NewApiResponses()
	apiSeekers := api.NewApiSeekers()
	resp, err := apiResponses.GetResponsesByID(id)
	if err != nil {
		http.Error(w, "Can not get response by username", http.StatusBadRequest)
		return
	}

	vacancy, err := apiVacancies.GetVacancyByVacancyID(resp.VacancyID)
	if err != nil {
		http.Error(w, "Can not get vacancy by id", http.StatusBadRequest)
		return
	}

	seeker, err := apiSeekers.CheckSeeker(resp.Username)
	if err != nil {
		http.Error(w, "Can not get seeker by username", http.StatusBadRequest)
		return
	}

	// response := models.ResponseWithVacancy{
	// 	RespID:      resp.RespID,
	// 	VacancyID:   resp.VacancyID,
	// 	EmployerID:  resp.EmployerID,
	// 	Username:    resp.Username,
	// 	Status:      resp.Status,
	// 	Company:     vacancy.Company,
	// 	Title:       vacancy.Title,
	// 	Description: vacancy.Description,
	// 	Tags:        vacancy.Tags,
	// }

	data := Response{
		Response: resp,
		Vacancy:  vacancy,
		Seeker:   seeker,
	}

	renderResponse(w, "seekers/response.html", data)
}
