package bot

import (
	"fmt"

	"github.com/qaZar1/HHforURFU/telegram_employers/internal/api"
	"github.com/qaZar1/HHforURFU/telegram_employers/internal/models"
)

type APITelegramMsg struct {
	ChannelID   string
	BotURL      string
	telegramAPI string
}

func NewApiTelegramMsg(channelID string, botURL string, telegramAPI string) *APITelegramMsg {
	return &APITelegramMsg{
		ChannelID:   channelID,
		BotURL:      botURL,
		telegramAPI: telegramAPI,
	}
}

func (apiTgMsg *APITelegramMsg) SendMessage(token string, vacancy_id int64, vacancy models.Vacancies, tags []string) error {
	const company = "Компания: "

	text_tags := "Метки: "

	for index, tag := range tags {
		if index == 0 {
			text_tags = text_tags + tag
		} else {
			text_tags = text_tags + ", " + tag
		}
	}

	text := fmt.Sprintf(vacancy.Title + "\n\n" + vacancy.Description + "\n\n" + company + vacancy.Company + "\n\n" + text_tags)
	url := fmt.Sprintf(apiTgMsg.telegramAPI, token)
	rightBotURL := fmt.Sprintf(apiTgMsg.BotURL, vacancy_id)

	// Создаем кнопку с URL
	button := models.InlineKeyboardButton{
		Text: "Перейти к боту",
		URL:  rightBotURL,
	}

	// Создаем разметку для кнопки
	keyboard := models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				button,
			},
		},
	}

	// Формируем запрос для отправки сообщения с кнопкой
	reqBody := models.SendMessageRequest{
		ChannelID:   apiTgMsg.ChannelID,
		Text:        text,
		ReplyMarkup: keyboard,
	}

	apiTG := api.NewApiTelegram(url)
	_, err := apiTG.AddEmployer(reqBody)
	if err != nil {
		return err
	}

	return nil
}
