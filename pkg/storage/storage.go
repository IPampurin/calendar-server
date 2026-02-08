package storage

import (
	"fmt"
	"sync"
	"time"
)

// Repository - интерфейс, реализующий требуемые методы
type Repository interface {
	Create(userID int, date time.Time, title, content string) (int, error) // добавляет event в хранилище, возвращает ID event или ошибку
	Update(event *Event) error                                             // обновляет event в хранилище, возвращает ошибку, если событие не найдено
	Delete(userID, eventID int) error                                      // удаляет event из хранилища, возвращает ошибку, если событие не найдено
	GetForDay(userID int, date time.Time) ([]*Event, error)                // возвращает перечень событий на день или ошибку
	GetForWeek(userID int, date time.Time) ([]*Event, error)               // возвращает перечень событий на неделю или ошибку
	GetForMonth(userID int, date time.Time) ([]*Event, error)              // возвращает перечень событий на месяц или ошибку
}

// Event описывает запись в календаре событий
type Event struct {
	ID      int       `json:"id"`                // id события (счётчик событий)
	UserID  int       `json:"user_id"`           // id пользователя
	Date    time.Time `json:"date"`              // дата события
	Title   string    `json:"title"`             // заголовок события
	Content string    `json:"content,omitempty"` // содержание события
}

// Storage используем для хранения информации календаря событий
type Storage struct {
	Mu     sync.RWMutex     // предполагаем конкурентный доступ к ресурсу
	Events map[int][]*Event // user_id -> events
	NextID int              // номер (ID) следующего Event (счётчик событий)
}

// Create добавляет event в хранилище, возвращает ID event или ошибку
func (s *Storage) Create(userID int, date time.Time, title, content string) (int, error) {

	// выполняем базовые проверки
	if userID < 0 {
		return 0, fmt.Errorf("ошибочный ID пользователя")
	}
	if title == "" {
		return 0, fmt.Errorf("поле title должно быть заполнено")
	}

	s.Mu.Lock()
	defer s.Mu.Unlock()

	// проверяем, что память под слайс событий есть и пользователь существует
	if _, ok := s.Events[userID]; !ok || s.Events[userID] == nil {
		s.Events[userID] = make([]*Event, 0)
	}

	// добавляем пользователю событие в список
	s.Events[userID] = append(s.Events[userID], &Event{
		ID:      s.NextID,
		UserID:  userID,
		Date:    date,
		Title:   title,
		Content: content,
	})
	// добаляем счётчик событий
	s.NextID++

	return s.NextID - 1, nil
}

// Update обновляет event в хранилище, возвращает ошибку, если событие не найдено
func (s *Storage) Update(event *Event) error {

	if event == nil {
		return fmt.Errorf("событие не может быть nil")
	}

	s.Mu.Lock()
	defer s.Mu.Unlock()

	events, ok := s.Events[event.UserID]
	if !ok {
		return fmt.Errorf("пользователь с %d не найден", event.UserID)
	}
	if events == nil {
		return fmt.Errorf("у пользователя с %d событий не найдено", event.UserID)
	}

	for i := 0; i < len(events); i++ {
		// если нашли событие - обновляем данные
		if event.ID == events[i].ID {
			events[i].Date = event.Date
			events[i].Title = event.Title
			events[i].Content = event.Content
			return nil
		}
	}

	// если событие не найдено - что-то пошло не так
	return fmt.Errorf("событие с %d не найдено", event.ID)
}

// Delete удаляет event из хранилища, возвращает ошибку, если событие не найдено
func (s *Storage) Delete(userID, eventID int) error {

	s.Mu.Lock()
	defer s.Mu.Unlock()

	events, ok := s.Events[userID]
	if !ok {
		return fmt.Errorf("пользователь с %d не найден", userID)
	}
	if events == nil {
		return fmt.Errorf("у пользователя с %d событий не найдено", userID)
	}

	for i := 0; i < len(events); i++ {
		// если нашли событие - удаляем событие
		if eventID == events[i].ID {
			events[i] = nil
			s.Events[userID] = append(events[:i], events[i+1:]...)
			// или events = slices.Delete(s.Events[userID], i, i+1)
			return nil
		}
	}

	// если событие не найдено - что-то пошло не так
	return fmt.Errorf("событие с %d не найдено", eventID)
}

// возвращает перечень событий на день или ошибку
func (s *Storage) GetForDay(userID int, date time.Time) ([]*Event, error) {

}

// возвращает перечень событий на неделю или ошибку
func (s *Storage) GetForWeek(userID int, date time.Time) ([]*Event, error) {

}

// возвращает перечень событий на месяц или ошибку
func (s *Storage) GetForMonth(userID int, date time.Time) ([]*Event, error) {

}

// NewStorage создаёт новое хранилище
func NewStorage() Repository {
	return &Storage{
		Events: make(map[int][]*Event),
		NextID: 1,
	}
}

func InitDB() error {

	return nil
}
