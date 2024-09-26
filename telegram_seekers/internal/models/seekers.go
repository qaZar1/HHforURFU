package models

type Seekers struct {
	Chat_ID  int64  `validate:"required"`
	Nickname string `validate:"required"`
	F_Name   string `validate:"required"`
	S_Name   string `validate:"required"`
	Resume   string `validate:"required"`
}
