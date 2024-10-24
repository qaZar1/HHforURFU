package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	employers "github.com/qaZar1/HHforURFU/web/internal/renders/employers"
	seekers "github.com/qaZar1/HHforURFU/web/internal/renders/seekers"
)

func main() {
	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./internal/static/"))))
	r.Get("/home", seekers.RenderSeekers)
	r.Get("/vacancies", seekers.RenderVacancies)
	r.Get("/vacancy/id-{id}", seekers.RenderVacancy)
	r.Get("/login", seekers.RenderLogIn)
	r.Get("/register", seekers.RenderRegister)
	r.Get("/responses", seekers.RenderResp)
	r.Get("/response/id-{id}", seekers.RenderResponse)
	r.HandleFunc("/submit", seekers.Submit)
	r.HandleFunc("/submitToLogin", seekers.SubmitToLogin)
	r.HandleFunc("/acceptVacancy", seekers.Accept)

	r.Get("/registerEmployers", employers.RenderRegisterFor)
	r.Get("/loginEmployers", employers.RenderLogInEmployers)
	r.Get("/employers", employers.RenderSeekers)
	r.Get("/employers/vacancies", employers.RenderVacancies)
	r.Get("/employers/vacancy/id-{id}", employers.RenderVacancy)
	r.Get("/employers/responses", employers.RenderResp)
	r.Get("/employers/response/id-{id}", employers.RenderResponse)
	r.Get("/employers/addVacancy", employers.RenderAddVacancy)
	r.HandleFunc("/submitEmployers", employers.SubmitEmployers)
	r.HandleFunc("/submitToLoginEmployers", employers.SubmitToLogin)
	r.HandleFunc("/acceptResponse", employers.AcceptResponse)
	r.HandleFunc("/refuseResponse", employers.RefuseResponse)
	r.HandleFunc("/add", employers.AddVacancyInList)
	r.HandleFunc("/deleteVacancy", employers.DeleteVacancy)

	http.ListenAndServe(":3000", r)
}
