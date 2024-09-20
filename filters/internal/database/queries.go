package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/filters/autogen"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) GetAllFilters() ([]autogen.Info, error) {
	const query = "SELECT vacancy_id, tags FROM main.filters;"

	var vacancies []autogen.Info
	if err := pg.db.Select(&vacancies, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.filters: %s", err)
	}

	return vacancies, nil
}

func (pg *Database) GetFiltersByVacancyID(vacancyId int64) (*[]autogen.Info, error) {
	const query = "SELECT vacancy_id, tags FROM main.filters WHERE vacancy_id = $1;"

	var vacancy []autogen.Info
	if err := pg.db.Select(&vacancy, query, vacancyId); err != nil {
		return nil, fmt.Errorf("User does not exist in main.filters: %w", err)
	}

	return &vacancy, nil
}

func (pg *Database) AddFilters(vacancy autogen.Filters) error {
	const query = `
INSERT INTO main.filters (
	vacancy_id,
	tags
) VALUES (
	:vacancyid,
	:tags
);`

	if _, err := pg.db.NamedExec(query, vacancy); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.filters: %s", err)
	}

	return nil
}

func (pg *Database) RemoveFilters(vacancyId int64) (bool, error) {
	const query = "DELETE FROM main.filters WHERE vacancy_id = $1"

	exec, err := pg.db.Exec(query, vacancyId)
	if err != nil {
		return false, fmt.Errorf("Invalid DELETE main.filters: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected vacancies: %s", err)
	}

	return affected == 0, nil
}