package models

type Employers struct {
	Chat_ID  int64  `validate:"required"`
	Nickname string `validate:"required"`
	Company  string `validate:"required"`
}
