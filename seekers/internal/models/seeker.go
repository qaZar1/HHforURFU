package models

type Seeker struct {
	ID           int64  `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	FirstName    string `json:"f_name" db:"f_name"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Resume       string `json:"resume" db:"resume"`
} // @name seeker
