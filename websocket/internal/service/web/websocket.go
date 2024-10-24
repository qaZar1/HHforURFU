package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/HHforURFU/websocket/internal/database"
	"github.com/qaZar1/HHforURFU/websocket/internal/models"
)

// Структура для хранения информации о WebSocket и базе данных
type WebSocket struct {
	db *database.Database
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Проверку на происхождение можно настроить для безопасности
	},
}

// Структура для хранения соединений пользователей
type Client struct {
	Conn     *websocket.Conn
	UserID   int
	UserType string
}

// Карта для хранения соединений клиентов
var clients = make(map[string]*Client)

// Канал для пересылки сообщений
var broadcast = make(chan models.Message)

// Функция для создания нового WebSocket с подключенной базой данных
func NewWebSocket(db *sqlx.DB) *WebSocket {
	return &WebSocket{
		db: database.NewDatabase(db),
	}
}

func (socket *WebSocket) HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Обновляем HTTP соединение до WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ws.Close()

	// Получаем идентификатор пользователя и тип пользователя
	userIDStr := r.URL.Query().Get("user_id")
	userType := r.URL.Query().Get("user_type")
	responseIDStr := r.URL.Query().Get("response_id")

	// Преобразуем строки в int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println("Неверный user_id:", err)
		return
	}
	responseID, err := strconv.Atoi(responseIDStr)
	if err != nil {
		log.Println("Неверный vacancy_id:", err)
		return
	}

	// Создаем уникальный ключ для клиента
	clientKey := userType + "_" + userIDStr

	// Сохраняем клиента в карте подключений
	clients[clientKey] = &Client{
		Conn:     ws,
		UserID:   userID,
		UserType: userType,
	}

	// Загружаем предыдущие сообщения по вакансии
	messages, err := socket.db.GetMessagesByVacancy(responseID)
	if err != nil {
		log.Printf("Ошибка получения сообщений для вакансии %d: %v", responseID, err)
	} else {
		// Отправляем предыдущие сообщения клиенту
		for _, message := range messages {
			err := ws.WriteJSON(message)
			if err != nil {
				log.Printf("Ошибка отправки предыдущего сообщения: %v", err)
				break
			}
		}
	}

	// Обработка новых сообщений от клиента
	for {
		var msg models.Message
		err := ws.ReadJSON(&msg) // Чтение сообщения от клиента
		if err != nil {
			log.Printf("Ошибка чтения сообщения: %v", err)
			delete(clients, clientKey)
			break
		}

		// Сохраняем сообщение в базу данных
		if err := socket.db.AddVacancy(msg); err != nil {
			log.Println("Ошибка сохранения сообщения:", err)
		}

		// Транслируем сообщение через канал broadcast
		broadcast <- msg
	}
}

// Обработка и пересылка сообщений
func (socket *WebSocket) HandleMessages() {
	for {
		// Получаем новое сообщение из канала broadcast
		msg := <-broadcast

		// Генерируем ключ для получателя
		receiverKey := msg.ReceiverType + "_" + strconv.Itoa(msg.ReceiverID)

		// Отправляем сообщение получателю
		if client, ok := clients[receiverKey]; ok {
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				log.Printf("Ошибка отправки сообщения: %v", err)
				client.Conn.Close()
				delete(clients, receiverKey)
			}
		}
	}
}
