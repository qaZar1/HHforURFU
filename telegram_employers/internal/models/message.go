package models

type SendMessageRequest struct {
	ChannelID   string               `json:"chat_id"`
	Text        string               `json:"text"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}
