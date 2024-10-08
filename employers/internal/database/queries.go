package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/employers/autogen"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) GetAllUsers() ([]autogen.Info, error) {
	const query = "SELECT chat_id, nickname, company FROM main.employers;"

	var users []autogen.Info
	if err := pg.db.Select(&users, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.employers: %s", err)
	}

	return users, nil
}

func (pg *Database) GetUserByChatID(chatId int64) (*autogen.Info, error) {
	const query = "SELECT chat_id, nickname, company FROM main.employers WHERE chat_id = $1;"

	var users autogen.Info
	if err := pg.db.Get(&users, query, chatId); err != nil {
		return nil, fmt.Errorf("User does not exist in main.employers: %w", err)
	}

	return &users, nil
}

func (pg *Database) AddUser(user autogen.User) (bool, error) {
	const query = `
INSERT INTO main.employers (
	chat_id,
	nickname,
	company
) VALUES (
	:chatid,
	:nickname,
	:company
) ON CONFLICT (chat_id) DO NOTHING;`

	exec, err := pg.db.NamedExec(query, user)
	if err != nil {
		return false, fmt.Errorf("Invalid INSERT INTO main.employers: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected employers: %s", err)
	}

	return affected == 0, nil
}

func (pg *Database) RemoveUser(chatId int64) (bool, error) {
	const query = "DELETE FROM main.employers WHERE chat_id = $1"

	exec, err := pg.db.Exec(query, chatId)
	if err != nil {
		return false, fmt.Errorf("Invalid DELETE main.employers: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected employers: %s", err)
	}

	return affected == 0, nil
}

func (pg *Database) UpdateUser(chat_id int64, updateUser autogen.UpdateUser) (bool, error) {
	const query = `
UPDATE main.employers
SET nickname = :nickname,
	company = :company
WHERE chat_id = :chat_id;`

	exec, err := pg.db.NamedExec(query, updateUser)
	if err != nil {
		return false, fmt.Errorf("Invalid UPDATE main.employers: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected description: %s", err)
	}

	return affected == 0, nil
}
