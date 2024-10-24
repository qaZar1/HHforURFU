package models

type Message struct {
	ResponseID   int    `json:"response_id" db:"response_id"`
	SenderType   string `json:"sender_type" db:"sender_type"` // "seeker" или "employer"
	SenderID     int    `json:"sender_id" db:"sender_id"`
	ReceiverType string `json:"receiver_type" db:"receiver_type"`
	ReceiverID   int    `json:"receiver_id" db:"receiver_id"`
	Text         string `json:"text" db:"message"`
} // @name message
