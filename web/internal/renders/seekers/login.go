package renders

import (
	"html/template"
	"net/http"
)

type StartPage struct{}

func renderLogIn(w http.ResponseWriter, templateName string) {
	tmpl, err := template.ParseFiles("internal/temp/" + templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderLogIn(w http.ResponseWriter, r *http.Request) {
	// id_str := chi.URLParam(r, "id")
	// id, err := strconv.ParseInt(id_str, 10, 64)
	// if err != nil {
	// 	http.Error(w, "Invalid ID", http.StatusBadRequest)
	// 	return
	// }

	// api := api.NewApiVacancies()
	// vacancy, _ := api.GetVacancyByVacancyID(id)

	// data := Vacancy{
	// 	Vacancy: vacancy,
	// }
	renderLogIn(w, "seekers/login.html")
}
