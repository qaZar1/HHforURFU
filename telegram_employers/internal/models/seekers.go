package models

type Seekers struct {
	Chat_ID  int64  `validate:"required"`
	Nickname string `validate:"required"`
	Fname    string `validate:"required" json="fname"`
	Sname    string `validate:"required" json="sname"`
	Resume   string `validate:"required"`
}

type Notify struct {
	Chat_ID          int64  `validate:"required"`
	Text             string `validate:"required"`
	Chat_ID_Employer int64  `validate:"required"`
	Username         string `validate:"required"`
}
