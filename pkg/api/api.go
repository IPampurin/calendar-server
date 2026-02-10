package api

import (
	"net/http"

	"github.com/IPampurin/calendar-server/pkg/storage"
)

type API struct {
	storage storage.Repository
}

func NewAPI(db storage.Repository) *API {
	return &API{storage: db}
}

func Init(db storage.Repository) {

	api := NewAPI(db)

	http.HandleFunc("/create_event", api.CreateEventHandler) // POST — создание нового события

	http.HandleFunc("/update_event", api.UpdateEventHandler) // POST — обновление существующего

	http.HandleFunc("/delete_event", api.DeleteEventHandler) // POST — удаление

	http.HandleFunc("/events_for_day", api.GetEventsForDayHandler) // GET — получить все события на день

	http.HandleFunc("/events_for_week", api.GetEventsForWeekHandler) // GET — события на неделю

	http.HandleFunc("/events_for_month", api.GetEventsForMonthHandler) // GET — события на месяц
}

// Answer - ответ на запрос к календарю
type Answer struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}
