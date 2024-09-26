package service

import (
	"fmt"

	"github.com/qaZar1/HHforURFU/telegram_seekers/autogen"
	"github.com/qaZar1/HHforURFU/telegram_seekers/internal/bot"
)

type Service struct {
	bot *bot.BotSeekers
}

func NewService(token string, seekers string, vacancies string, responses string, tags string) *Service {
	return &Service{
		bot: bot.NewBotSeekers(token, seekers, vacancies, responses, tags),
	}
}

func (srv *Service) Send(notification autogen.Notification) error {
	text := fmt.Sprintf("%s\n\nСвязь с работодателем: @%s", notification.Text, notification.Username)
	return srv.bot.Send(notification.ChatId, text)
}
