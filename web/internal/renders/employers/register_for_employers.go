package renders

import (
	"html/template"
	"net/http"
)

func renderRegister(w http.ResponseWriter, templateName string) {
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

func RenderRegisterFor(w http.ResponseWriter, r *http.Request) {
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
	renderRegister(w, "employers/register_for_employers.html")
}
