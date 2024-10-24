package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/seekers/internal/models"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) AddSeeker(seeker models.Seeker) (bool, error) {
	const query = `
INSERT INTO main.seekers (
	username,
	f_name,
	password_hash,
	resume
) VALUES (
	:username,
	:f_name,
	:password_hash,
	:resume
) ON CONFLICT (username) DO NOTHING;`

	result, err := pg.db.NamedExec(query, seeker)
	if err != nil {
		return false, fmt.Errorf("Invalid INSERT INTO main.seekers: %s", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid INSERT INTO main.seekers: %s", err)
	}

	return affected == 1, nil
}

func (pg *Database) CheckSeekerUsername(username string) (models.Seeker, error) {
	const query = "SELECT id, username, f_name, password_hash, resume FROM main.seekers WHERE username = $1;"

	var seeker models.Seeker
	if err := pg.db.Get(&seeker, query, username); err != nil {
		return models.Seeker{}, fmt.Errorf("User does not exist in main.seekers: %w", err)
	}

	return seeker, nil
}

// func (pg *Database) GetAllVacancies() ([]models.Vacancy, error) {
// 	const query = `
// 	SELECT vacancy_id, company, title, description, employer_id, status
// 	FROM main.vacancies
// 	ORDER BY
//     CASE
//         WHEN status = 'Открыта' THEN 1
//         ELSE 2
//     END,
//     vacancy_id DESC;`

// 	var vacancies []models.Vacancy
// 	if err := pg.db.Select(&vacancies, query); err != nil {
// 		return nil, fmt.Errorf("Invalid SELECT main.vacancies: %s", err)
// 	}

// 	return vacancies, nil
// }

// func (pg *Database) AddVacancy(vacancy autogen.Vacancy) (int64, error) {
// 	const query = `
// INSERT INTO main.vacancies (
// 	company,
// 	title,
// 	description,
// 	chat_id_employer
// ) VALUES (
// 	$1,
// 	$2,
// 	$3,
// 	$4
// ) ON CONFLICT (vacancy_id) DO NOTHING
// RETURNING vacancy_id;`

// 	var vacancy_id int64
// 	err := pg.db.QueryRowx(query, vacancy.Company, vacancy.Title, vacancy.Description, vacancy.ChatIdEmployer).Scan(&vacancy_id)
// 	if err != nil {
// 		return 0, fmt.Errorf("Invalid INSERT INTO main.vacancies: %s", err)
// 	}

// 	return vacancy_id, nil
// }

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
