package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/responses/autogen"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) GetAllResponses() ([]autogen.Response, error) {
	const query = "SELECT vacancy_id, chat_id, status, chat_id_employer FROM main.responses;"

	var vacancies []autogen.Response
	if err := pg.db.Select(&vacancies, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.responses: %s", err)
	}

	return vacancies, nil
}

func (pg *Database) GetResponsesByVacancyIDAndChatIDEmployer(vacancyId int64, chatIdEmployer int64) (*autogen.Response, error) {
	const query = "SELECT vacancy_id, chat_id, status FROM main.responses WHERE vacancy_id = $1 AND chat_id_employer = $2;"

	var vacancy autogen.Response
	if err := pg.db.Get(&vacancy, query, vacancyId, chatIdEmployer); err != nil {
		return nil, fmt.Errorf("User does not exist in main.vacancies: %w", err)
	}

	return &vacancy, nil
}

func (pg *Database) AddResponses(respons autogen.Response) error {
	const query = `
INSERT INTO main.responses (
	vacancy_id,
	chat_id,
	status,
	chat_id_employer
) VALUES (
	:vacancy_id,
	:chat_id,
	:status,
	:chat_id_employer
);`

	if _, err := pg.db.NamedExec(query, respons); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.responses: %s", err)
	}

	return nil
}

func (pg *Database) RemoveResponses(vacancyId int64) (bool, error) {
	const query = "DELETE FROM main.responses WHERE vacancy_id = $1"

	exec, err := pg.db.Exec(query, vacancyId)
	if err != nil {
		return false, fmt.Errorf("Invalid DELETE main.responses: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected responses: %s", err)
	}

	return affected == 0, nil
}

func (pg *Database) UpdateResponses(vacancyId int64, updateRespons autogen.Response) (bool, error) {
	const query = `
UPDATE main.responses
SET	status = :status
WHERE vacancy_id = :vacancy_id;`

	exec, err := pg.db.NamedExec(query, updateRespons)
	if err != nil {
		return false, fmt.Errorf("Invalid UPDATE main.responses: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected description: %s", err)
	}

	return affected == 0, nil
}

func (pg *Database) GetResponsesByChatIDEmployer(chat_id_employer int64) ([]autogen.Response, error) {
	const query = "SELECT vacancy_id, chat_id, status, chat_id_employer FROM main.responses WHERE chat_id_employer = $1 AND status = 'Ожидает проверки';"

	var vacancies []autogen.Response
	if err := pg.db.Select(&vacancies, query, chat_id_employer); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.responses: %s", err)
	}

	return vacancies, nil
}

func (pg *Database) GetResponsesByChatID(chat_id int64) ([]autogen.Response, error) {
	const query = "SELECT vacancy_id, chat_id, status, chat_id_employer FROM main.responses WHERE chat_id = $1;"

	var responses []autogen.Response
	if err := pg.db.Select(&responses, query, chat_id); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.responses: %s", err)
	}

	return responses, nil
}

func (pg *Database) GetResponsesByVacancyID(vacancyId int64) ([]autogen.Response, error) {
	const query = "SELECT vacancy_id, chat_id, status, chat_id_employer FROM main.responses WHERE vacancy_id = $1;"

	var responses []autogen.Response
	if err := pg.db.Select(&responses, query, vacancyId); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.responses: %s", err)
	}

	return responses, nil
}

func (pg *Database) GetResponsesByVacancyIDAndChatID(vacancyId int64, chatId int64) (*autogen.Response, error) {
	const query = "SELECT vacancy_id, chat_id, status FROM main.responses WHERE vacancy_id = $1 AND chat_id = $2;"

	var response autogen.Response
	if err := pg.db.Get(&response, query, vacancyId, chatId); err != nil {
		return nil, fmt.Errorf("User does not exist in main.vacancies: %w", err)
	}

	return &response, nil
}
