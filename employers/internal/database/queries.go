package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/employers/internal/models"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) AddEmployer(employer models.Employer) (bool, error) {
	const query = `
INSERT INTO main.employers (
	username,
	f_name,
	password_hash,
	company
) VALUES (
	:username,
	:f_name,
	:password_hash,
	:company
) ON CONFLICT (employer_id) DO NOTHING;`

	result, err := pg.db.NamedExec(query, employer)
	if err != nil {
		return false, fmt.Errorf("Invalid INSERT INTO main.employers: %s", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid INSERT INTO main.employers: %s", err)
	}

	return affected == 1, nil
}

func (pg *Database) CheckEmployerUsername(employerID int64) (models.Employer, error) {
	const query = "SELECT employer_id, username, f_name, password_hash, company FROM main.employers WHERE employer_id = $1;"

	var employer models.Employer
	if err := pg.db.Get(&employer, query, employerID); err != nil {
		return models.Employer{}, fmt.Errorf("User does not exist in main.employers: %w", err)
	}

	return employer, nil
}

func (pg *Database) CheckEmployerByUsername(username string) (models.Employer, error) {
	const query = "SELECT employer_id, username, f_name, password_hash, company FROM main.employers WHERE username = $1;"

	var employer models.Employer
	if err := pg.db.Get(&employer, query, username); err != nil {
		return models.Employer{}, fmt.Errorf("User does not exist in main.employers: %w", err)
	}

	return employer, nil
}
