package api

import "net/http"

func Init() {

	http.HandleFunc("/api/create_event") // POST — создание нового события

	http.HandleFunc("/api/update_event") // POST — обновление существующего

	http.HandleFunc("/api/delete_event") // POST — удаление

	http.HandleFunc("/api/events_for_day") // GET — получить все события на день

	http.HandleFunc("/api/events_for_week") // GET — события на неделю

	http.HandleFunc("/api/events_for_month") // GET — события на месяц
}
