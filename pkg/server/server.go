package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IPampurin/calendar-server/pkg/api"
	"github.com/IPampurin/calendar-server/pkg/storage"
)

const calendarPortDefault = "8081"

// Run запускает сервер и передаёт объект хранилища хэндлерам
func Run(db storage.Repository) error {

	port, ok := os.LookupEnv("CALENDAR_PORT")
	if !ok {
		port = calendarPortDefault
	}

	// инициализируем api
	api.Init(db)

	// настраиваем логирование
	logger, logFile, err := SetupLogging()
	if err != nil {
		return fmt.Errorf("ошибка настройки логирования: %w", err)
	}
	defer logFile.Close()

	// оборачиваем в middleware
	handler := LoggingMiddleware(logger)(http.DefaultServeMux)

	// http.Handle("/", http.FileServer(http.Dir("web"))) // здесь может быть запуск фронтэнда

	// запускаем сервер
	addr := fmt.Sprintf(":%s", port)
	logger.Printf("Сервер запущен на %s\n", addr)

	return http.ListenAndServe(addr, handler)
}
