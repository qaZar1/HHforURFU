package database

import (
	"fmt"
	"strings"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/vacancies/internal/models"
	"github.com/sirupsen/logrus"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) GetAllVacancies() ([]models.Vacancy, error) {
	const query = `
	SELECT 
    main.vacancies.vacancy_id,
    main.vacancies.company,
    main.vacancies.title,
    main.vacancies.description,
    main.vacancies.employer_id,
    main.vacancies.status,
    STRING_AGG(main.filters.tag, ', ') AS tags
FROM 
    main.vacancies
LEFT JOIN 
    main.filters ON main.vacancies.vacancy_id = main.filters.vacancy_id
WHERE 
    main.vacancies.status = 'Открыта' AND main.filters.tag != 'NULL'
GROUP BY 
    main.vacancies.vacancy_id, 
    main.vacancies.company, 
    main.vacancies.title, 
    main.vacancies.description, 
    main.vacancies.employer_id, 
    main.vacancies.status
ORDER BY 
    main.vacancies.status;
`

	var vacanciesWithTagStr []models.Tags
	if err := pg.db.Select(&vacanciesWithTagStr, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.vacancies: %s", err)
	}

	var vacancies []models.Vacancy
	for _, vacancyWithTagStr := range vacanciesWithTagStr {
		tags := strings.Split(vacancyWithTagStr.Tag, ", ")
		vacancy := models.Vacancy{
			VacancyID:   vacancyWithTagStr.VacancyID,
			Company:     vacancyWithTagStr.Company,
			Title:       vacancyWithTagStr.Title,
			Description: vacancyWithTagStr.Description,
			EmployerID:  vacancyWithTagStr.EmployerID,
			Status:      vacancyWithTagStr.Status,
			Tag:         tags,
		}

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func (pg *Database) GetVacancyByVacancyID(vacancyId int64) (models.Vacancy, error) {
	const query = `
SELECT 
    main.vacancies.vacancy_id,
    main.vacancies.company,
    main.vacancies.title,
    main.vacancies.description,
    main.vacancies.employer_id,
    main.vacancies.status,
    STRING_AGG(main.filters.tag, ', ') AS tags
FROM 
    main.vacancies
LEFT JOIN 
    main.filters ON main.vacancies.vacancy_id = main.filters.vacancy_id
WHERE 
    main.vacancies.vacancy_id = $1
GROUP BY 
    main.vacancies.vacancy_id, 
    main.vacancies.company, 
    main.vacancies.title, 
    main.vacancies.description, 
    main.vacancies.employer_id, 
    main.vacancies.status;
`

	var vacancyWithTagStr models.Tags
	if err := pg.db.Get(&vacancyWithTagStr, query, vacancyId); err != nil {
		return models.Vacancy{}, fmt.Errorf("User does not exist in main.vacancies: %w", err)
	}

	tags := strings.Split(vacancyWithTagStr.Tag, ", ")
	vacancy := models.Vacancy{
		VacancyID:   vacancyWithTagStr.VacancyID,
		Company:     vacancyWithTagStr.Company,
		Title:       vacancyWithTagStr.Title,
		Description: vacancyWithTagStr.Description,
		EmployerID:  vacancyWithTagStr.EmployerID,
		Status:      vacancyWithTagStr.Status,
		Tag:         tags,
	}

	return vacancy, nil
}

func (pg *Database) GetAllVacanciesByEmployerID(employer_id int64) ([]models.Vacancy, error) {
	const query = `
	SELECT 
    main.vacancies.vacancy_id,
    main.vacancies.company,
    main.vacancies.title,
    main.vacancies.description,
    main.vacancies.employer_id,
    main.vacancies.status,
    STRING_AGG(main.filters.tag, ', ') AS tags
FROM 
    main.vacancies
LEFT JOIN 
    main.filters ON main.vacancies.vacancy_id = main.filters.vacancy_id
WHERE 
    main.vacancies.status = 'Открыта' AND main.filters.tag != 'NULL' AND main.vacancies.employer_id = $1
GROUP BY 
    main.vacancies.vacancy_id, 
    main.vacancies.company, 
    main.vacancies.title, 
    main.vacancies.description, 
    main.vacancies.employer_id, 
    main.vacancies.status
ORDER BY 
    main.vacancies.status;
`

	var vacanciesWithTagStr []models.Tags
	if err := pg.db.Select(&vacanciesWithTagStr, query, employer_id); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.vacancies: %s", err)
	}

	var vacancies []models.Vacancy
	for _, vacancyWithTagStr := range vacanciesWithTagStr {
		tags := strings.Split(vacancyWithTagStr.Tag, ", ")
		vacancy := models.Vacancy{
			VacancyID:   vacancyWithTagStr.VacancyID,
			Company:     vacancyWithTagStr.Company,
			Title:       vacancyWithTagStr.Title,
			Description: vacancyWithTagStr.Description,
			EmployerID:  vacancyWithTagStr.EmployerID,
			Status:      vacancyWithTagStr.Status,
			Tag:         tags,
		}

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func (pg *Database) AddVacancy(vacancy models.AddVacancy) (int64, error) {
	const query = `
INSERT INTO main.vacancies (
	company,
	title,
	description,
	employer_id,
	status
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
) RETURNING vacancy_id;`

	var vacancy_id int64
	if err := pg.db.QueryRowx(query, vacancy.Company, vacancy.Title, vacancy.Description, vacancy.EmployerID, vacancy.Status).Scan(&vacancy_id); err != nil {
		logrus.Errorf("Invalid add vacancy: %s", err)
		return 0, err
	}

	return vacancy_id, nil
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
