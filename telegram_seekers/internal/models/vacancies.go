package models

type Vacancies struct {
	Vacancy_ID       int64  `validate:"required"`
	Company          string `validate:"required"`
	Title            string `validate:"required"`
	Description      string `validate:"required"`
	Chat_ID_Employer int64  `validate:"required"`
}
