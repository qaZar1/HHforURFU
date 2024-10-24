package models

type Message struct {
	VacancyID    int    `json:"vacancy_id"`
	SenderType   string `json:"sender_type"` // "seeker" или "employer"
	SenderID     int    `json:"sender_id"`
	ReceiverType string `json:"receiver_type"`
	ReceiverID   int    `json:"receiver_id"`
	Text         string `json:"text"`
}
