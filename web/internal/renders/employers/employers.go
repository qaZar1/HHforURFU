package renders

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	tmpl, err := template.ParseFiles("internal/temp/" + templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil) // Здесь можно передать данные в шаблон, если нужно
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderSeekers(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "employers/home.html")
}
