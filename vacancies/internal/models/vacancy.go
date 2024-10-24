package models

type Vacancy struct {
	VacancyID   int64    `json:"vacancy_id" db:"vacancy_id"`
	Company     string   `json:"company" db:"company"`
	Title       string   `json:"title" db:"title"`
	Description string   `json:"description" db:"description"`
	EmployerID  int64    `json:"employer_id" db:"employer_id"`
	Status      string   `json:"status" db:"status"`
	Tag         []string `json:"tags" db:"tags"`
} // @name vacancy

type Tags struct {
	VacancyID   int64  `json:"vacancy_id" db:"vacancy_id"`
	Company     string `json:"company" db:"company"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	EmployerID  int64  `json:"employer_id" db:"employer_id"`
	Status      string `json:"status" db:"status"`
	Tag         string `json:"tags" db:"tags"`
}

type AddVacancy struct {
	VacancyID   int64  `json:"vacancy_id" db:"vacancy_id"`
	Company     string `json:"company" db:"company"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	EmployerID  int64  `json:"employer_id" db:"employer_id"`
	Status      string `json:"status" db:"status"`
} // @name add_vacancy

type Num struct {
	VacancyID int64 `json:"vacancy_id" db:"vacancy_id"`
} // @name vacancyID
