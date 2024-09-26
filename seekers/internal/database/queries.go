package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/seekers/autogen"
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
	const query = "SELECT chat_id, nickname, f_name, s_name, resume FROM main.seekers;"

	var users []autogen.Info
	if err := pg.db.Select(&users, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.seekers: %s", err)
	}

	return users, nil
}

func (pg *Database) GetUserByChatID(chatId int64) (*autogen.Info, error) {
	const query = "SELECT chat_id, nickname, f_name, s_name, resume FROM main.seekers WHERE chat_id = $1;"

	var users autogen.Info
	if err := pg.db.Get(&users, query, chatId); err != nil {
		return nil, fmt.Errorf("User does not exist in main.seekers: %w", err)
	}

	return &users, nil
}

func (pg *Database) AddUser(user autogen.User) error {
	const query = `
INSERT INTO main.seekers (
	chat_id,
	nickname,
	f_name,
	s_name,
	resume
) VALUES (
	:chatid,
	:nickname,
	:fname,
	:sname,
	:resume
) ON CONFLICT (chat_id) DO NOTHING;`

	if _, err := pg.db.NamedExec(query, user); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.seekers: %s", err)
	}

	return nil
}

func (pg *Database) RemoveUser(chatId int64) (bool, error) {
	const query = "DELETE FROM main.seekers WHERE chat_id = $1"

	exec, err := pg.db.Exec(query, chatId)
	if err != nil {
		return false, fmt.Errorf("Invalid DELETE main.seekers: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected seekers: %s", err)
	}

	return affected == 0, nil
}

func (pg *Database) UpdateUser(chat_id int64, updateUser autogen.UpdateUser) (bool, error) {
	const query = `
UPDATE main.seekers
SET nickname = :nickname,
	f_name = :f_name,
	s_name = :s_name,
	resume = :resume
WHERE chat_id = :chat_id;`

	exec, err := pg.db.NamedExec(query, updateUser)
	if err != nil {
		return false, fmt.Errorf("Invalid UPDATE main.versions: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected description: %s", err)
	}

	return affected == 0, nil
}
