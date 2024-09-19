package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/vacancies/autogen"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) GetAllVacancies() ([]autogen.Info, error) {
	const query = "SELECT vacancy_id, company, title, description FROM main.vacancies;"

	var vacancies []autogen.Info
	if err := pg.db.Select(&vacancies, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.vacancies: %s", err)
	}

	return vacancies, nil
}

func (pg *Database) GetVacancyByVacancyID(vacancyId int64) (*autogen.Info, error) {
	const query = "SELECT vacancy_id, company, title, description FROM main.vacancies WHERE vacancy_id = $1;"

	var vacancy autogen.Info
	if err := pg.db.Get(&vacancy, query, vacancyId); err != nil {
		return nil, fmt.Errorf("User does not exist in main.vacancies: %w", err)
	}

	return &vacancy, nil
}

func (pg *Database) AddVacancy(vacancy autogen.Vacancy) error {
	const query = `
INSERT INTO main.vacancies (
	vacancy_id,
	company,
	title,
	description
) VALUES (
	:vacancyid,
	:company,
	:title,
	:description
) ON CONFLICT (vacancy_id) DO NOTHING;`

	if _, err := pg.db.NamedExec(query, vacancy); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.vacancies: %s", err)
	}

	return nil
}

func (pg *Database) RemoveVacancy(vacancyId int64) (bool, error) {
	const query = "DELETE FROM main.vacancies WHERE vacancy_id = $1"

	exec, err := pg.db.Exec(query, vacancyId)
	if err != nil {
		return false, fmt.Errorf("Invalid DELETE main.vacancies: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected vacancies: %s", err)
	}

	return affected == 0, nil
}

func (pg *Database) UpdateVacancy(vacancyId int64, updateVacancy autogen.UpdateVacancy) (bool, error) {
	const query = `
UPDATE main.vacancies
SET	company = :company
	title = :title,
	description = :description
WHERE vacancy_id = :vacancy_id;`

	exec, err := pg.db.NamedExec(query, updateVacancy)
	if err != nil {
		return false, fmt.Errorf("Invalid UPDATE main.vacancies: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected description: %s", err)
	}

	return affected == 0, nil
}
