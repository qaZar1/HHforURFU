package models

type Responses struct {
	Chat_ID          int64  `validate:"required"`
	Vacancy_ID       int64  `validate:"required"`
	Status           string `validate:"required"`
	Chat_ID_Employer int64  `validate:"required"`
}
