package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/qaZar1/HHforURFU/telegram_employers/internal/bot"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			file := frame.File[len(path.Dir(os.Args[0]))+1:]
			line := frame.Line
			return "", fmt.Sprintf("%s:%d", file, line)
		},
	})
}

func main() {
	const (
		address        = "ADDRESS"
		auth           = "APIS_AUTH_BASIC"
		tokenEmployers = "TOKEN_EMPLOYERS"
		botURL         = "BOT_URL"
		channelID      = "CHANNEL_ID"
		telegramAPI    = "TELEGRAM_API"
		APIEmployers   = "API_EMPLOYERS"
		APIVacancies   = "API_VACANCIES"
		APIResponses   = "API_RESPONSES"
		APISeekers     = "API_SEEKERS"
		APIMessage     = "API_MESSAGE"
		APITags        = "API_TAGS"
	)

	bot.NewBotEmployers(os.Getenv(tokenEmployers),
		os.Getenv(APIEmployers),
		os.Getenv(APIVacancies),
		os.Getenv(APIResponses),
		os.Getenv(APISeekers),
		os.Getenv(APITags),
		os.Getenv(APIMessage),
		os.Getenv(channelID),
		os.Getenv(botURL),
		os.Getenv(telegramAPI))
}
