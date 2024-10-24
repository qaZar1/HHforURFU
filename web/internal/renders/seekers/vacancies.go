package renders

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type Vacancies struct {
	Vacancies []models.Vacancy
}

// Функция для рендеринга шаблона с вакансиями
func renderVacancies(w http.ResponseWriter, templateName string, data Vacancies) {
	tmpl, err := template.ParseFiles("internal/temp/" + templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Функция для фильтрации вакансий по категории и названию
func filterVacancies(vacancies []models.Vacancy, searchQuery, category string) []models.Vacancy {
	filteredVacanciesMap := make(map[int]models.Vacancy) // Карта для уникальных вакансий

	// Приведение поискового запроса к нижнему регистру для некейсчувствительного поиска
	searchQuery = strings.ToLower(searchQuery)

	for _, vacancy := range vacancies {
		// Логика фильтрации
		// Если оба условия заданы, вакансия должна соответствовать и названию, и категории
		matchesSearch := searchQuery == "" || strings.Contains(strings.ToLower(vacancy.Title), searchQuery)
		matchesCategory := category == "" || containsTag(vacancy.Tags, category)

		if matchesSearch && matchesCategory {
			// Если вакансия соответствует обоим условиям или одно из них не задано
			filteredVacanciesMap[int(vacancy.ID)] = vacancy
		}
	}

	// Преобразуем карту обратно в срез
	var filteredVacancies []models.Vacancy
	for _, vacancy := range filteredVacanciesMap {
		filteredVacancies = append(filteredVacancies, vacancy)
	}

	return filteredVacancies
}

// Вспомогательная функция для проверки наличия категории в тегах вакансии
func containsTag(tags []string, category string) bool {
	for _, tag := range tags {
		if tag == category {
			return true
		}
	}
	return false
}

// Основная функция для отображения вакансий с фильтрацией
func RenderVacancies(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры фильтров из URL
	searchQuery := r.URL.Query().Get("search") // Поисковый запрос по названию
	category := r.URL.Query().Get("category")  // Фильтр по категории

	// Получаем все вакансии через API
	api := api.NewApiVacancies()
	vacancies, _ := api.GetAllVacancies()

	// Фильтруем вакансии на основе поискового запроса и категории
	filteredVacancies := filterVacancies(vacancies, searchQuery, category)

	// Подготавливаем данные для шаблона
	data := Vacancies{
		Vacancies: filteredVacancies,
	}

	// Рендерим шаблон с отфильтрованными вакансиями
	renderVacancies(w, "seekers/vacancies.html", data)
}
