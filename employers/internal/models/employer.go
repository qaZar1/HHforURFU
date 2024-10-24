package models

type Employer struct {
	EmployerID   int64  `json:"employer_id" db:"employer_id"`
	Username     string `json:"username" db:"username"`
	FirstName    string `json:"f_name" db:"f_name"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Company      string `json:"company" db:"company"`
} // @name employer
