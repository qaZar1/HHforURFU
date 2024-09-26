package bot

import (
	"testing"

	"github.com/qaZar1/HHforURFU/telegram_employers/internal/models"
	"github.com/sirupsen/logrus"
)

// mockEmployers реализует интерфейс APIEmployers
type mockEmployers struct {
	employers map[int64]models.Employers
}

func NewMockEmployers() *mockEmployers {
	return &mockEmployers{
		employers: make(map[int64]models.Employers),
	}
}

func (m *mockEmployers) GetEmployerByChatID(chatID int64) (models.Employers, bool, error) {
	employer, ok := m.employers[chatID]
	return employer, ok, nil
}

func (m *mockEmployers) AddEmployer(employer models.Employers) (models.Employers, error) {
	m.employers[employer.Chat_ID] = employer
	return employer, nil
}

func TestBotEmployers_HandleStartCommand(t *testing.T) {
	//var mockEmployers api.APIEmployers = // NewMockEmployers() // Убедитесь, что передаете интерфейс
	//bot := &BotEmployers{
	//	dbApiEmployers: mockEmployers,
	//}

	// Остальная часть теста...
	logrus.Print("123")
}
