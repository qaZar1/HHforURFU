package renders

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type Vacancy struct {
	Vacancy models.Vacancy
}

func renderResp(w http.ResponseWriter, templateName string, data Vacancy) {
	tmpl, err := template.ParseFiles("internal/temp/" + templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderVacancy(w http.ResponseWriter, r *http.Request) {
	id_str := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	api := api.NewApiVacancies()
	vacancy, _ := api.GetVacancyByVacancyID(id)

	data := Vacancy{
		Vacancy: vacancy,
	}
	renderResp(w, "employers/vacancy.html", data)
}
