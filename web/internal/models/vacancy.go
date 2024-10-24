package models

type Vacancy struct {
	ID          int64    `json:"vacancy_id"`
	Company     string   `json:"company"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	EmployerID  int64    `json:"employer_id"`
	Tags        []string `json:"tags"`
}

type AddVacancy struct {
	ID          int64  `json:"vacancy_id"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	EmployerID  int64  `json:"employer_id"`
}

type Temp struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	Username    string   `json:"username"`
	Tags        []string `json:"tags"`
}
