package db

import (
	"sync"
	"time"
)

// Event описывает запись в календаре событий
type Event struct {
	ID      int       `json:"id"`                // id события
	UserID  int       `json:"user_id"`           // id пользователя
	Date    time.Time `json:"date"`              // дата события
	Title   string    `json:"title"`             // заголовок события
	Content string    `json:"content,omitempty"` // содержание события
}

// Storage используем для хранения информации календаря событий
type Storage struct {
	mu     sync.RWMutex
	Events map[int][]*Event // user_id -> events
	nextID int              // номер (ID) следующего Event
}

func InitDB() error {

	storage := make(map[int][]*Event)

	return nil
}
