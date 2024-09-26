package bot

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/qaZar1/HHforURFU/telegram_seekers/internal/api"
	"github.com/qaZar1/HHforURFU/telegram_seekers/internal/models"

	logrus "github.com/sirupsen/logrus"
)

//go:generate mockgen -source=bot.go -package mocks -destination ../../autogen/mocks/telegram.go
type ITelegramAPI interface {
	Send(c tg.Chattable) (tg.Message, error)
	GetUpdatesChan(cfg tg.UpdateConfig) tg.UpdatesChannel
}

type BotSeekers struct {
	tgApi          *tg.BotAPI
	dbApiSeekers   api.APISeekers
	dbApiVacancies api.APIVacancies
	dbApiResponses api.APIResponses
	dbApiTags      api.APITags
}

type UserState struct {
	WaitingFor  string
	FirstName   string
	LastName    string
	Resume      string
	CommandArgs string
}

const (
	waitForName     = "waitForName"
	waitForLastName = "waitForLastName"
	waitForResume   = "waitForResume"
	responseMsg     = "Отклик %d:\nНазвание вакансии:%s\nОписание вакансии:%s\nКомпания:%s\n\n"
)

var userStates = make(map[int64]*UserState)

func NewBotSeekers(token string, seekers string, vacancies string, responses string, tags string) *BotSeekers {
	botApi, err := tg.NewBotAPI(token)
	if err != nil {
		logrus.Errorf("Invalid token: %s", err)
	}

	bot := &BotSeekers{
		tgApi:          botApi,
		dbApiSeekers:   *api.NewApiSeekers(seekers),
		dbApiVacancies: *api.NewApiVacancies(vacancies),
		dbApiResponses: *api.NewApiResponses(responses),
		dbApiTags:      *api.NewApiTags(tags),
	}

	bot.consume()
	return bot
}

func (bot *BotSeekers) consume() {
	go func() {
		updater := tg.NewUpdate(0)
		updater.Timeout = 60
		updates := bot.tgApi.GetUpdatesChan(updater)

		for update := range updates {
			if err := bot.handle(update); err != nil {
				logrus.Errorf("Invalid update: %s", err)
				continue
			}
		}
	}()
}

func (bot *BotSeekers) handle(update tg.Update) error {
	if update.Message != nil {
		chatID := update.Message.Chat.ID
		if update.Message.IsCommand() && update.Message.Command() == "start" {
			ok, err := bot.dbApiSeekers.GetSeekerByChatID(chatID)
			if err != nil {
				logrus.Errorf("Invalid get user by chat id: %s", err)
			}

			if !ok {
				// Начало процесса регистрации
				msg := tg.NewMessage(chatID, "Введите свое имя:")
				bot.tgApi.Send(msg)

				userStates[chatID] = &UserState{
					WaitingFor:  waitForName,
					CommandArgs: update.Message.CommandArguments(),
				}
				return nil
			}

			args := update.Message.CommandArguments()
			if args != "" {
				return bot.handleCommandArgs(chatID, args)
			}

			// Стандартное поведение для /start
			msg := tg.NewMessage(chatID, "Добро пожаловать в бота!")
			buttons := tg.NewReplyKeyboard(
				tg.NewKeyboardButtonRow(
					tg.NewKeyboardButton("Просмотреть мои отклики"),
				),
			)

			msg.ReplyMarkup = buttons
			_, err = bot.tgApi.Send(msg)
			if err != nil {
				logrus.Errorf("Invalid send msg")
			}

		}

		switch update.Message.Text {
		case "/start", "Начать":
			ok, err := bot.dbApiSeekers.GetSeekerByChatID(chatID)
			if err != nil {
				logrus.Errorf("Invalid get user by chat id: %s", err)
			}

			if !ok {
				// Начало процесса регистрации
				msg := tg.NewMessage(chatID, "Введите свое имя:")
				bot.tgApi.Send(msg)

				// Устанавливаем состояние ожидания ввода имени и сохраняем аргументы команды, если они есть
				userStates[chatID] = &UserState{
					WaitingFor:  waitForName,
					CommandArgs: update.Message.CommandArguments(), // Сохраняем аргументы
				}
				return nil
			}

			args := update.Message.CommandArguments() // Получаем переданные аргументы
			if args != "" {
				return bot.handleCommandArgs(chatID, args) // Вынесли обработку команды с аргументами в отдельную функцию
			}

			// Стандартное поведение для /start
			msg := tg.NewMessage(chatID, "Добро пожаловать в бота!")
			buttons := tg.NewReplyKeyboard(
				tg.NewKeyboardButtonRow(
					tg.NewKeyboardButton("Просмотреть мои отклики"),
				),
			)

			msg.ReplyMarkup = buttons
			_, err = bot.tgApi.Send(msg)
			if err != nil {
				logrus.Errorf("Invalid send msg: %v", err)
			}
		case "Просмотреть мои отклики":
			responses, err := bot.dbApiResponses.GetResponsesByChatID(update.Message.From.ID)
			if err != nil {
				logrus.Errorf("Invalid get responses: %v", err)
			}

			text := ""

			for index, value := range responses {
				vacancy, err := bot.dbApiVacancies.GetVacancyByVacancyID(value.Vacancy_ID)
				if err != nil {
					logrus.Errorf("Invalid get vacancy by vacancy_id: %v", err)
				}

				text = fmt.Sprintf(text+responseMsg, index+1, vacancy.Title, vacancy.Description, vacancy.Company)
			}

			msg := tg.NewMessage(chatID, text)

			buttons := tg.NewReplyKeyboard(
				tg.NewKeyboardButtonRow(
					tg.NewKeyboardButton("Просмотреть мои отклики"),
				),
			)

			msg.ReplyMarkup = buttons
			_, err = bot.tgApi.Send(msg)
			if err != nil {
				logrus.Errorf("Invalid send msg: %v", err)
			}
		case "В главное меню":
			buttons := tg.NewReplyKeyboard(
				tg.NewKeyboardButtonRow(
					tg.NewKeyboardButton("Просмотреть мои отклики"),
				),
			)

			msg := tg.NewMessage(update.Message.Chat.ID, "Выберите, что вы хотите сделать:")
			msg.ReplyMarkup = buttons

			_, err := bot.tgApi.Send(msg)
			if err != nil {
				logrus.Errorf("Invalid send msg: %v", err)
			}
		default:
			if userStates[update.Message.From.ID] == nil {
				msg := tg.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду :(\nЧтобы начать мной пользоваться, отправьте мне слово \"В главное меню\"")

				buttons := tg.NewReplyKeyboard(
					tg.NewKeyboardButtonRow(
						tg.NewKeyboardButton("В главное меню"),
					),
				)

				msg.ReplyMarkup = buttons
				if _, err := bot.tgApi.Send(msg); err != nil {
					log.Println(err)
				}
			}
		}

		// Обработка состояний пользователя
		if userState, exists := userStates[chatID]; exists {
			if err := bot.RegSeeker(userState, update, userStates); err != nil {
				return err
			}
			return nil
		}

	}

	if update.CallbackQuery != nil {
		// Обработка нажатий на кнопки (если есть)
		data := strings.Split(update.CallbackQuery.Data, ":")
		if len(data) == 2 && data[0] == "accept" {
			vacancyID, err := strconv.Atoi(data[1])
			if err != nil {
				logrus.Errorf("Invalid vacancy ID: %s", err)
				return err
			}

			// Получаем вакансию из базы данных по vacancyID
			vacancy, err := bot.dbApiVacancies.GetVacancyByVacancyID(int64(vacancyID))
			if err != nil {
				logrus.Errorf("Invalid get vacancy by ID: %s", err)
				return err
			}

			ok, err := bot.dbApiResponses.GetResponsesByVacancyIDAndChatID(vacancy.Vacancy_ID, update.CallbackQuery.From.ID)
			if err != nil {
				logrus.Errorf("Invalid get responses by ID: %s", err)
			}
			if ok {
				bot.MakeButton(update, "Вы уже откликнулись на данную вакансию! Повторно это сделать нельзя!")
				return nil
			}
			// Добавляем отклик на вакансию
			ok, err = bot.dbApiResponses.AddResponses(models.Responses{
				Chat_ID:          update.CallbackQuery.From.ID,
				Vacancy_ID:       vacancy.Vacancy_ID,
				Status:           "Ожидает проверки",
				Chat_ID_Employer: vacancy.Chat_ID_Employer,
			})
			if err != nil || !ok {
				logrus.Errorf("Invalid add response: %s", err)
			} else {
				bot.MakeButton(update, "Вы успешно откликнулись на вакансию!")
			}
		} else {
			bot.MakeButton(update, "Вы отклонили вакансию!")
		}
	}
	return nil
}

func (bot *BotSeekers) getVacancyFromChannel(messageID string) (models.Vacancies, error) {
	id, err := strconv.Atoi(messageID)
	if err != nil {
		return models.Vacancies{}, err
	}
	vacancy, err := bot.dbApiVacancies.GetVacancyByVacancyID(int64(id))
	if err != nil {
		return models.Vacancies{}, err
	}

	return vacancy, nil
}

func (bot *BotSeekers) getMessageFromChannel(vacancy models.Vacancies, tags []string) (string, error) {
	text_tags := "Метки: "
	for index, tag := range tags {
		if index == 0 {
			text_tags = text_tags + tag
		} else {
			text_tags = text_tags + ", " + tag
		}
	}
	text := fmt.Sprintf(vacancy.Title + "\n\n" + vacancy.Description + "\n\n" + "Компания: " + vacancy.Company + "\n\n" + text_tags)

	return text, nil
}

func (bot *BotSeekers) MakeButton(update tg.Update, desc string) {
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, desc)

	if _, err := bot.tgApi.Send(msg); err != nil {
		log.Println(err)
	}
}

func (bot *BotSeekers) RegSeeker(userState *UserState, update tg.Update, userStates map[int64]*UserState) error {
	switch userState.WaitingFor {
	case waitForName:
		// Получаем имя пользователя и переходим к запросу фамилии
		userState.FirstName = update.Message.Text
		msg := tg.NewMessage(update.Message.From.ID, "Введите свою фамилию:")
		bot.tgApi.Send(msg)

		// Обновляем состояние на ожидание фамилии
		userState.WaitingFor = waitForLastName
		return nil

	case waitForLastName:
		// Получаем фамилию и завершаем процесс регистрации
		userState.LastName = update.Message.Text

		msg := tg.NewMessage(update.Message.From.ID, "Вставьте ссылку на резюме:\nПример ссылки: https://google.com")
		bot.tgApi.Send(msg)

		// Обновляем состояние на ожидание резюме
		userState.WaitingFor = waitForResume
		return nil

	case waitForResume:
		userState.Resume = update.Message.Text
		usrname := update.Message.From.UserName
		logrus.Print(usrname)
		_, err := bot.dbApiSeekers.AddSeeker(models.Seekers{
			Chat_ID:  update.Message.From.ID,
			Nickname: update.Message.From.UserName,
			F_Name:   userState.FirstName,
			S_Name:   userState.LastName,
			Resume:   userState.Resume,
		})
		if err != nil {
			logrus.Errorf("Invalid add user: %s", err)
		}

		msg := tg.NewMessage(update.Message.From.ID, "Спасибо! Ваши данные сохранены.")
		bot.tgApi.Send(msg)

		// Если есть сохраненные аргументы команды, обрабатываем их
		if userState.CommandArgs != "" {
			bot.handleCommandArgs(update.Message.Chat.ID, userState.CommandArgs)
		}

		// Удаляем состояние пользователя
		delete(userStates, update.Message.From.ID)
		return nil
	}
	return nil
}

func (bot *BotSeekers) handleCommandArgs(chatID int64, args string) error {
	vacancy, err := bot.getVacancyFromChannel(args)
	if err != nil {
		return err
	}

	filters, err := bot.dbApiTags.GetFiltersByVacancyID(vacancy.Vacancy_ID)
	if err != nil {
		logrus.Errorf("Invalid get filters by vacancy_ID: %d", err)
	}

	var tags []string
	for _, filter := range filters {
		tags = append(tags, filter.Tags)
	}

	messageText, err := bot.getMessageFromChannel(vacancy, tags)
	if err != nil {
		logrus.Errorf("Invalid get vacancy by ID: %d", err)
	}

	buttons := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Принять", fmt.Sprintf("accept:%d", vacancy.Vacancy_ID)),
			tg.NewInlineKeyboardButtonData("Отклонить", fmt.Sprintf("deny:%d", vacancy.Vacancy_ID)),
		),
	)

	msg := tg.NewMessage(chatID, messageText)
	msg.ReplyMarkup = buttons
	bot.tgApi.Send(msg)

	return nil
}

func (bot *BotSeekers) Send(chatID int64, text string) error {
	msg := tg.NewMessage(chatID, text)
	_, err := bot.tgApi.Send(msg)
	if err != nil {
		return fmt.Errorf("Invalid send msg for %d: %s", err)
	}

	return nil
}
