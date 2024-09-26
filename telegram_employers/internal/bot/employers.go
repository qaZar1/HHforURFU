package bot

import (
	"fmt"
	"log"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/qaZar1/HHforURFU/telegram_employers/internal/api"
	"github.com/qaZar1/HHforURFU/telegram_employers/internal/models"

	logrus "github.com/sirupsen/logrus"
)

type BotEmployers struct {
	tgApi          *tg.BotAPI
	dbApiEmployers api.APIEmployers
	dbApiVacancies api.APIVacancies
	dbApiResponses api.APIResponses
	dbApiSeekers   api.APISeekers
	dbApiTags      api.APITags
	tgMessage      api.APIMessage
	tgApiMsg       APITelegramMsg
	responses      []models.Responses // Массив записей из базы
	userStates     map[int64]int
}

type EmployerState struct {
	WaitingFor string
	Company    string
}

type VacancyState struct {
	WaitingFor  string
	Title       string
	Description string
	Tags        []string
}

const (
	waitForCompany   = "waitForCompany"
	waitForTitle     = "waitForTitle"
	waitForDesc      = "waitForDesc"
	waitForTags      = "waitForTags"
	toAddVacancy     = "redirect_to_add_vacancy"
	toCheckResponses = "redirect_to_check_responses"
	toTarifs         = "redirect_to_tarifs"
	toMain           = "redirect_to_main"
)

var (
	employerStates = make(map[int64]*EmployerState)
	vacancyStates  = make(map[int64]*VacancyState)
)

func NewBotEmployers(token string,
	httpEmployer string,
	httpVacancies string,
	httpResponses string,
	httpSeekers string,
	httpTags string,
	message string,
	channelID string,
	botURL string,
	telegramAPI string,
) {
	botApi, err := tg.NewBotAPI(token)
	if err != nil {
		logrus.Errorf("Invalid token: %s", err)
	}

	bot := &BotEmployers{
		tgApi:          botApi,
		dbApiEmployers: *api.NewApiEmployers(httpEmployer),
		dbApiVacancies: *api.NewApiVacancies(httpVacancies),
		dbApiResponses: *api.NewApiResponses(httpResponses),
		dbApiSeekers:   *api.NewApiSeekers(httpSeekers),
		dbApiTags:      *api.NewApiTags(httpTags),
		tgMessage:      *api.NewApiMessage(message),
		tgApiMsg:       *NewApiTelegramMsg(channelID, botURL, telegramAPI),
		userStates:     make(map[int64]int),
	}

	bot.consume()
}

func (bot *BotEmployers) consume() {
	updater := tg.NewUpdate(0)
	updater.Timeout = 60
	updates := bot.tgApi.GetUpdatesChan(updater)

	for update := range updates {
		if err := bot.handle(update); err != nil {
			logrus.Errorf("Invalid update: %s", err)
			continue
		}
	}
}

func (bot *BotEmployers) handle(update tg.Update) error {
	if update.Message != nil {
		chatID := update.Message.Chat.ID
		switch update.Message.Text {
		case "/start":
			_, ok, err := bot.dbApiEmployers.GetEmployerByChatID(chatID)
			if err != nil {
				logrus.Errorf("Invalid get employer by chat id: %s", err)
			}

			if !ok {
				// Начало процесса регистрации
				msg := tg.NewMessage(chatID, "Введите название компании:")
				_, err := bot.tgApi.Send(msg)
				if err != nil {
					return err
				}

				// Устанавливаем состояние ожидания ввода имени
				employerStates[chatID] = &EmployerState{WaitingFor: waitForCompany}
				return nil
			}
		case "В главное меню":
			delete(vacancyStates, update.Message.From.ID)
			delete(employerStates, update.Message.From.ID)
			msg := tg.NewMessage(update.Message.Chat.ID, "Выберите, что вы хотите сделать:")

			buttons := tg.NewReplyKeyboard(
				tg.NewKeyboardButtonRow(
					tg.NewKeyboardButton("Создать вакансию"),
				),
				tg.NewKeyboardButtonRow(
					tg.NewKeyboardButton("Просмотреть отклики"),
				),
			)

			msg.ReplyMarkup = buttons
			if _, err := bot.tgApi.Send(msg); err != nil {
				log.Println(err)
			}
		case "Создать вакансию":
			buttons := tg.NewReplyKeyboard(
				tg.NewKeyboardButtonRow(
					tg.NewKeyboardButton("В главное меню"),
				),
			)
			msg := tg.NewMessage(update.Message.From.ID, "Введите название вакансии")
			msg.ReplyMarkup = buttons

			if _, err := bot.tgApi.Send(msg); err != nil {
				log.Println(err)
			}
			// Устанавливаем состояние ожидания ввода имени
			vacancyStates[update.Message.From.ID] = &VacancyState{WaitingFor: waitForTitle}
			return nil
		case "Просмотреть отклики":
			if err := bot.LoadRecords(update.Message.From.ID); err != nil {
				logrus.Errorf("Invalid add map: %v", err)

				msg := tg.NewMessage(update.Message.Chat.ID, "Нет доступных откликов")

				buttons := tg.NewReplyKeyboard(
					tg.NewKeyboardButtonRow(
						tg.NewKeyboardButton("Создать вакансию"),
					),
					tg.NewKeyboardButtonRow(
						tg.NewKeyboardButton("Просмотреть отклики"),
					),
				)

				msg.ReplyMarkup = buttons
				if _, err := bot.tgApi.Send(msg); err != nil {
					log.Println(err)
				}
				return nil
			}
			if err := bot.StartSendingRecords(update.Message.From.ID); err != nil {
				logrus.Errorf("Invalid start read map: %v", err)
			}
		case "Принять", "Отклонить":
			// Получаем текущую запись
			if len(bot.userStates) <= 0 {
				bot.sendConfirmationMessage(chatID, "Зайдите в \"Просмотреть отклики\" заново!")
				msg := tg.NewMessage(update.Message.Chat.ID, "Выберите, что вы хотите сделать:")

				buttons := tg.NewReplyKeyboard(
					tg.NewKeyboardButtonRow(
						tg.NewKeyboardButton("Создать вакансию"),
					),
					tg.NewKeyboardButtonRow(
						tg.NewKeyboardButton("Просмотреть отклики"),
					),
				)

				msg.ReplyMarkup = buttons
				if _, err := bot.tgApi.Send(msg); err != nil {
					log.Println(err)
				}
				return nil
			}
			userIndex := bot.userStates[chatID]
			currentRecord := bot.responses[userIndex]

			// TODO сделать обратную связь для искателя
			if update.Message.Text == "Принять" {
				ok, err := bot.dbApiResponses.UpdateResponse(currentRecord.Vacancy_ID, models.Responses{
					Chat_ID:          currentRecord.Chat_ID,
					Vacancy_ID:       currentRecord.Vacancy_ID,
					Status:           "Принято",
					Chat_ID_Employer: currentRecord.Chat_ID_Employer,
				})
				if err != nil {
					logrus.Errorf("Invalid UPDATE response: %v", err)
				}
				if ok {
					employer, ok, err := bot.dbApiEmployers.GetEmployerByChatID(currentRecord.Chat_ID_Employer)
					if err != nil || !ok {
						logrus.Errorf("Invalid get vacancy by vacancy_id: %v", err)
						return nil
					}
					notify := models.Notify{
						Chat_ID:          currentRecord.Chat_ID,
						Chat_ID_Employer: currentRecord.Chat_ID_Employer,
						Username:         employer.Nickname,
						Text:             "Ваш отклик на вакансию был одобрен!",
					}

					ok, err = bot.tgMessage.SendMsg(notify)
					if err != nil {
						logrus.Errorf("Invalid get vacancy by vacancy_id: %v", err)
						return nil
					}
					if ok {
						bot.sendConfirmationMessage(chatID, "Вы приняли запись.")
					}
				} else {
					bot.sendConfirmationMessage(chatID, "Вы не можете принять запись.")
				}
			} else {
				ok, err := bot.dbApiResponses.UpdateResponse(currentRecord.Vacancy_ID, models.Responses{
					Chat_ID:          currentRecord.Chat_ID,
					Vacancy_ID:       currentRecord.Vacancy_ID,
					Status:           "Отказано",
					Chat_ID_Employer: currentRecord.Chat_ID_Employer,
				})
				if err != nil {
					logrus.Errorf("Invalid UPDATE response: %v", err)
				}
				if ok {
					bot.sendConfirmationMessage(chatID, "Вы отклонили запись.")
				} else {
					bot.sendConfirmationMessage(chatID, "Вы не можете отклонить запись.")
				}
			}
			// После нажатия кнопки "Принять" или "Отклонить", отправляем следующую запись
			bot.handleResponse(update)
		default:
			if vacancyStates[update.Message.From.ID] == nil && employerStates[update.Message.From.ID] == nil {
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

		if employerState, exists := employerStates[chatID]; exists {
			if err := bot.RegEmployer(employerState, update, employerStates); err != nil {
				return err
			}

			buttons := tg.NewInlineKeyboardMarkup(
				tg.NewInlineKeyboardRow(
					tg.NewInlineKeyboardButtonData("Начать", toMain),
				),
			)

			msg := tg.NewMessage(update.Message.Chat.ID, "Добро пожаловать! Нажмите кнопку, чтобы начать испольование бота.")
			msg.ReplyMarkup = buttons

			if _, err := bot.tgApi.Send(msg); err != nil {
				log.Println(err)
			}
		}
		if vacancyState, exists := vacancyStates[chatID]; exists {
			if err := bot.NewVacancy(vacancyState, update, vacancyStates, bot.tgApi.Token); err != nil {
				return err
			}

			return nil
		}
	}
	return nil
}

func (bot *BotEmployers) MakeButton(update tg.Update, desc string, buttons tg.InlineKeyboardMarkup) {
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, desc)
	msg.ReplyMarkup = &buttons

	if _, err := bot.tgApi.Send(msg); err != nil {
		log.Println(err)
	}
}

func (bot *BotEmployers) RegEmployer(employerState *EmployerState, update tg.Update, employerStates map[int64]*EmployerState) error {
	switch employerState.WaitingFor {
	case waitForCompany:
		// Получаем имя пользователя и переходим к запросу фамилии
		employerState.Company = update.Message.Text

		_, err := bot.dbApiEmployers.AddEmployer(models.Employers{
			Chat_ID:  update.Message.From.ID,
			Nickname: update.Message.From.UserName,
			Company:  employerState.Company,
		})
		if err != nil {
			logrus.Errorf("Invalid add user: %s", err)
		}

		msg := tg.NewMessage(update.Message.From.ID, "Спасибо! Ваши данные сохранены.")
		if _, err := bot.tgApi.Send(msg); err != nil {
			log.Println(err)
		}

		// Удаляем состояние пользователя
		delete(employerStates, update.Message.From.ID)
		return nil
	}
	return nil
}

func (bot *BotEmployers) NewVacancy(newVacancy *VacancyState, update tg.Update, newVacancies map[int64]*VacancyState, token string) error {
	switch newVacancy.WaitingFor {
	case waitForTitle:
		// Получаем имя пользователя и переходим к запросу фамилии
		newVacancy.Title = update.Message.Text
		newVacancy.WaitingFor = waitForDesc

		buttons := tg.NewReplyKeyboard(
			tg.NewKeyboardButtonRow(
				tg.NewKeyboardButton("В главное меню"),
			),
		)

		msg := tg.NewMessage(update.Message.Chat.ID, "Введите описание вакансии:")
		msg.ReplyMarkup = buttons

		if _, err := bot.tgApi.Send(msg); err != nil {
			log.Println(err)
		}
		return nil

	case waitForDesc:
		newVacancy.Description = update.Message.Text
		buttons := tg.NewReplyKeyboard(
			tg.NewKeyboardButtonRow(
				tg.NewKeyboardButton("В главное меню"),
			),
		)

		msg := tg.NewMessage(update.Message.Chat.ID, "Введите метки для вакансии через \", \"\n(Пример: временная работа, гибкий график):")
		msg.ReplyMarkup = buttons

		if _, err := bot.tgApi.Send(msg); err != nil {
			log.Println(err)
		}

		newVacancy.WaitingFor = waitForTags
		return nil

	case waitForTags:
		text := update.Message.Text
		newVacancy.Tags = strings.Split(text, ", ")

		company, err := bot.GetCompany(&update)
		if err != nil {
			logrus.Errorf("Invalid get company: %d", err)
		}

		vacancy := models.Vacancies{
			Company:          company,
			Title:            newVacancy.Title,
			Description:      newVacancy.Description,
			Chat_ID_Employer: update.Message.From.ID,
		}

		ok, vacancy_id, err := bot.dbApiVacancies.AddVacancy(vacancy)
		if err != nil {
			logrus.Errorf("Invalid add user: %s", err)
		}

		vacancy, err = bot.dbApiVacancies.GetVacancyByVacancyID(vacancy_id)
		if err != nil {
			logrus.Errorf("Invalid get vacancy: %s", err)
		}

		if ok {
			for _, tag := range newVacancy.Tags {
				bot.dbApiTags.AddFilters(models.Filters{
					Tags:       tag,
					Vacancy_ID: vacancy_id,
				})
			}
			// Удаляем состояние пользователя
			delete(vacancyStates, update.Message.From.ID)

			msg := tg.NewMessage(update.Message.Chat.ID, "Ваша вакансия успешно добавлена")

			if _, err := bot.tgApi.Send(msg); err != nil {
				log.Println(err)
			}

			err = bot.tgApiMsg.SendMessage(token, vacancy_id, vacancy, newVacancy.Tags)
			return nil
		}

		return err

	}
	return nil
}

func (bot *BotEmployers) GetCompany(update *tg.Update) (string, error) {
	employers, _, err := bot.dbApiEmployers.GetEmployerByChatID(update.Message.From.ID)
	if err != nil {
		return "", err
	}

	return employers.Company, nil
}

// Получаем записи из базы (например, по лимиту)
func (bot *BotEmployers) LoadRecords(chatIdEmployer int64) error {
	records, err := bot.dbApiResponses.GetResponsesByChatIDEmployer(chatIdEmployer)
	if err != nil {
		return err
	}
	bot.responses = records
	return nil
}

// Начинаем отправку записей пользователю
func (bot *BotEmployers) StartSendingRecords(chatID int64) error {
	// Сбрасываем пользователя на первую запись
	bot.userStates[chatID] = 0

	if err := bot.sendNextRecord(chatID); err != nil {
		return err
	}
	return nil
}

// Отправляем следующую запись
func (bot *BotEmployers) sendNextRecord(chatID int64) error {
	userIndex := bot.userStates[chatID]
	if userIndex >= len(bot.responses) {
		// Если все записи закончились
		msg := tg.NewMessage(chatID, "Больше нет откликов")

		buttons := tg.NewReplyKeyboard(
			tg.NewKeyboardButtonRow(
				tg.NewKeyboardButton("Создать вакансию"),
			),
			tg.NewKeyboardButtonRow(
				tg.NewKeyboardButton("Просмотреть отклики"),
			),
		)

		msg.ReplyMarkup = buttons
		if _, err := bot.tgApi.Send(msg); err != nil {
			return err
		}

		return nil
	}

	record := bot.responses[userIndex]
	// Генерируем сообщение для записи
	vacancy, err := bot.dbApiVacancies.GetVacancyByVacancyID(record.Vacancy_ID)
	if err != nil {
		return err
	}
	seekers, err := bot.dbApiSeekers.GetSeekerByChatID(record.Chat_ID)
	if err != nil {
		return err
	}

	messageText := fmt.Sprintf(
		"Вакансия: %s\nОписание: %s\nКомпания: %s\n\n\nДанные об откликнувшемся\nИмя: %s\nФамилия: %s\nСсылка на резюме: %s\n\nСсылка на пользователя: @%s",
		vacancy.Title, vacancy.Description, vacancy.Company, seekers.Fname, seekers.Sname, seekers.Resume, seekers.Nickname)

	// Генерируем кнопки
	buttons := tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton("Принять"),
			tg.NewKeyboardButton("Отклонить"),
		),
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton("В главное меню"),
		),
	)

	// Отправляем сообщение с кнопками
	msg := tg.NewMessage(chatID, messageText)
	msg.ReplyMarkup = buttons
	_, err = bot.tgApi.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

// Обрабатываем ответ пользователя на запись
func (bot *BotEmployers) handleResponse(update tg.Update) error {
	chatID := update.Message.Chat.ID

	// Получаем текущее состояние пользователя
	userIndex := bot.userStates[chatID]
	if userIndex >= len(bot.responses) {
		return nil // Записей больше нет
	}

	// Переходим к следующей записи
	bot.userStates[chatID]++
	return bot.sendNextRecord(chatID)
}

// Отправляем сообщение с подтверждением
func (bot *BotEmployers) sendConfirmationMessage(chatID int64, text string) {
	msg := tg.NewMessage(chatID, text)
	bot.tgApi.Send(msg)
}
