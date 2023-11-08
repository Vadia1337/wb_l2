package main

import (
	"calendarApi/config"
	"calendarApi/handler"
	"calendarApi/repository"
	"calendarApi/usecase"
	"net/http"
)

func main() {
	cfg := config.GetConfig()

	rep := repository.NewMapRepository()

	crudCalendarUC := crudCalendar.NewUseCase(rep)

	h := handler.NewHandler(crudCalendarUC)

	http.HandleFunc("/create_event", h.MiddlewareLogger(h.CreateEvent))

	err := http.ListenAndServe(cfg.HttpServerPort, nil)
	if err != nil {
		return
	}

}
