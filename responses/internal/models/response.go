package models

type Response struct {
	RespID     int64  `json:"response_id" db:"response_id"`
	VacancyID  int64  `json:"vacancy_id" db:"vacancy_id"`
	EmployerID int64  `json:"employer_id" db:"employer_id"`
	Username   string `json:"username" db:"username"`
	Status     string `json:"status" db:"status"`
} // @name response
