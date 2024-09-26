package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/Impisigmatus/service_core/middlewares"
	"github.com/qaZar1/HHforURFU/telegram_seekers/autogen"
	"github.com/qaZar1/HHforURFU/telegram_seekers/internal/service"
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
		token         = "TOKEN_SEEKERS"
		auth          = "APIS_AUTH_BASIC"
		address       = "ADDRESS"
		api_seekers   = "API_SEEKERS"
		api_vacancies = "API_VACANCIES"
		api_responses = "API_RESPONSES"
		api_tags      = "API_TAGS"
	)

	transport := service.NewTransport(os.Getenv(token),
		os.Getenv(api_seekers),
		os.Getenv(api_vacancies),
		os.Getenv(api_responses),
		os.Getenv(api_tags))

	router := http.NewServeMux()
	router.Handle("/api/", middlewares.Use(autogen.Handler(transport), middlewares.Logger()))

	server := &http.Server{
		Addr:    os.Getenv(address),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Panic(err)
	}
}
