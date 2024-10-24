package models

type Tag struct {
	VacancyID int64  `json:"vacancy_id" db:"vacancy_id"`
	Tag       string `json:"tag" db:"tag"`
} // @name tag
