package models

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
	Mu     sync.RWMutex
	Events map[int][]*Event // user_id -> events
	NextID int              // номер (ID) следующего Event
}

// Answer - ответ на запрос к календарю
type Answer struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}
