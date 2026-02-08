package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IPampurin/calendar-server/pkg/api"
)

const calendarPortDefault = "8081"

func Run() error {

	port, ok := os.LookupEnv("CALENDAR_PORT")
	if !ok {
		port = calendarPortDefault
	}

	api.Init()

	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
