package renders

import (
	"html/template"
	"net/http"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
	"github.com/sirupsen/logrus"
)

type Vacancies struct {
	Vacancies []models.Vacancy
}

func renderVacancies(w http.ResponseWriter, templateName string, data Vacancies) {
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

func RenderVacancies(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		logrus.Errorf("Invalid parse cookie: %s", err)

		renderVacancies(w, "employers/vacanciesEmployers.html", Vacancies{})
		return
	}

	apiEmployers := api.NewApiEmployers()
	apiVacancies := api.NewApiVacancies()

	employer, err := apiEmployers.CheckEmployerByUsername(cookie.Value)
	if err != nil {
		logrus.Errorf("Invalid get employer: %s", err)
	}

	_ = employer
	vacancies, err := apiVacancies.GetAllVacanciesByEmployerID(employer.EmployerID)
	if err != nil {
		logrus.Errorf("Invalid get vacancies: %s", err)
	}

	data := Vacancies{
		Vacancies: vacancies,
	}
	renderVacancies(w, "employers/vacanciesEmployers.html", data)
}
