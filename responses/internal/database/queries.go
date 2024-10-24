package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/responses/internal/models"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

// func (pg *Database) GetAllResponses() ([]autogen.Response, error) {
// 	const query = "SELECT vacancy_id, chat_id, status, chat_id_employer FROM main.responses;"

// 	var vacancies []autogen.Response
// 	if err := pg.db.Select(&vacancies, query); err != nil {
// 		return nil, fmt.Errorf("Invalid SELECT main.responses: %s", err)
// 	}

// 	return vacancies, nil
// }

// func (pg *Database) GetResponsesByVacancyIDAndChatIDEmployer(vacancyId int64, chatIdEmployer int64) (*autogen.Response, error) {
// 	const query = "SELECT vacancy_id, chat_id, status FROM main.responses WHERE vacancy_id = $1 AND chat_id_employer = $2;"

// 	var vacancy autogen.Response
// 	if err := pg.db.Get(&vacancy, query, vacancyId, chatIdEmployer); err != nil {
// 		return nil, fmt.Errorf("User does not exist in main.vacancies: %w", err)
// 	}

// 	return &vacancy, nil
// }

func (pg *Database) AddResponses(respons models.Response) error {
	const query = `
INSERT INTO main.responses (
	vacancy_id,
	employer_id,
	username,
	status
) VALUES (
	:vacancy_id,
	:employer_id,
	:username,
	:status
);`

	if _, err := pg.db.NamedExec(query, respons); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.responses: %s", err)
	}

	return nil
}

// func (pg *Database) RemoveResponses(vacancyId int64) (bool, error) {
// 	const query = "DELETE FROM main.responses WHERE vacancy_id = $1"

// 	exec, err := pg.db.Exec(query, vacancyId)
// 	if err != nil {
// 		return false, fmt.Errorf("Invalid DELETE main.responses: %s", err)
// 	}

// 	affected, err := exec.RowsAffected()
// 	if err != nil {
// 		return false, fmt.Errorf("Invalid affected responses: %s", err)
// 	}

// 	return affected == 0, nil
// }

func (pg *Database) UpdateResponses(response_id int, updateResponse models.Response) (bool, error) {
	const query = `
UPDATE main.responses
SET	status = :status
WHERE response_id = :response_id;`

	params := map[string]interface{}{
		"status":      updateResponse.Status,
		"response_id": response_id,
	}

	exec, err := pg.db.NamedExec(query, params)
	if err != nil {
		return false, fmt.Errorf("Invalid UPDATE main.responses: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected description: %s", err)
	}

	return affected == 0, nil
}

// func (pg *Database) GetResponsesByChatIDEmployer(chat_id_employer int64) ([]autogen.Response, error) {
// 	const query = "SELECT vacancy_id, chat_id, status, chat_id_employer FROM main.responses WHERE chat_id_employer = $1 AND status = 'Ожидает проверки';"

// 	var vacancies []autogen.Response
// 	if err := pg.db.Select(&vacancies, query, chat_id_employer); err != nil {
// 		return nil, fmt.Errorf("Invalid SELECT main.responses: %s", err)
// 	}

// 	return vacancies, nil
//}

func (pg *Database) GetResponsesByUsername(username string) ([]models.Response, error) {
	const query = "SELECT response_id, vacancy_id, employer_id, username, status FROM main.responses WHERE username = $1;"

	var responses []models.Response
	if err := pg.db.Select(&responses, query, username); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.responses: %s", err)
	}

	return responses, nil
}

func (pg *Database) GetResponsesById(id int64) (models.Response, error) {
	const query = "SELECT response_id, vacancy_id, employer_id, username, status FROM main.responses WHERE response_id = $1;"

	var response models.Response
	if err := pg.db.Get(&response, query, id); err != nil {
		return models.Response{}, fmt.Errorf("Invalid SELECT main.responses: %s", err)
	}

	return response, nil
}

func (pg *Database) GetResponsesByEmployersID(id int64) ([]models.Response, error) {
	const query = `
SELECT response_id, vacancy_id, employer_id, username, status
FROM main.responses
WHERE employer_id = $1 AND (status = 'Ожидает ответа' OR status = 'Принято');`

	var responses []models.Response
	if err := pg.db.Select(&responses, query, id); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.responses: %s", err)
	}

	return responses, nil
}

// func (pg *Database) GetResponsesByVacancyID(vacancyId int64) ([]autogen.Response, error) {
// 	const query = "SELECT vacancy_id, chat_id, status, chat_id_employer FROM main.responses WHERE vacancy_id = $1;"

// 	var responses []autogen.Response
// 	if err := pg.db.Select(&responses, query, vacancyId); err != nil {
// 		return nil, fmt.Errorf("Invalid SELECT main.responses: %s", err)
// 	}

// 	return responses, nil
// }

// func (pg *Database) GetResponsesByVacancyIDAndChatID(vacancyId int64, chatId int64) (*autogen.Response, error) {
// 	const query = "SELECT vacancy_id, chat_id, status FROM main.responses WHERE vacancy_id = $1 AND chat_id = $2;"

// 	var response autogen.Response
// 	if err := pg.db.Get(&response, query, vacancyId, chatId); err != nil {
// 		return nil, fmt.Errorf("User does not exist in main.vacancies: %w", err)
// 	}

// 	return &response, nil
// }
