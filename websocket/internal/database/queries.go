package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/websocket/internal/models"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) AddVacancy(message models.Message) error {
	const query = `
INSERT INTO main.messages (
	response_id,
	sender_type,
	sender_id,
	receiver_type,
	receiver_id,
	message
) VALUES (
	:response_id,
	:sender_type,
	:sender_id,
	:receiver_type,
	:receiver_id,
	:message
);`

	if _, err := pg.db.NamedExec(query, message); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.messages: %s", err)
	}

	return nil
}

func (db *Database) GetMessagesByVacancy(responseID int) ([]models.Message, error) {
	var messages []models.Message
	query := `SELECT response_id, sender_type, sender_id, receiver_type, receiver_id, message FROM main.messages WHERE response_id = $1`
	err := db.db.Select(&messages, query, responseID)
	return messages, err
}

// func (pg *Database) RemoveVacancy(vacancyId int64) (bool, error) {
// 	const query = "DELETE FROM main.vacancies WHERE vacancy_id = $1"

// 	exec, err := pg.db.Exec(query, vacancyId)
// 	if err != nil {
// 		return false, fmt.Errorf("Invalid DELETE main.vacancies: %s", err)
// 	}

// 	affected, err := exec.RowsAffected()
// 	if err != nil {
// 		return false, fmt.Errorf("Invalid affected vacancies: %s", err)
// 	}

// 	return affected == 0, nil
// }

// func (pg *Database) UpdateVacancy(vacancyId int64, updateVacancy autogen.UpdateVacancy) (bool, error) {
// 	const query = `
// UPDATE main.vacancies
// SET	company = :company
// 	title = :title,
// 	description = :description
// 	chat_id_employer = :chatIdEmployer
// WHERE vacancy_id = :vacancy_id;`

// 	exec, err := pg.db.NamedExec(query, updateVacancy)
// 	if err != nil {
// 		return false, fmt.Errorf("Invalid UPDATE main.vacancies: %s", err)
// 	}

// 	affected, err := exec.RowsAffected()
// 	if err != nil {
// 		return false, fmt.Errorf("Invalid affected description: %s", err)
// 	}

// 	return affected == 0, nil
// }
