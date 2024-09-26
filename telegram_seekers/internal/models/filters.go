package models

type Filters struct {
	Vacancy_ID int64  `validate:"required"`
	Tags       string `validate:"required" json="tags"`
}
