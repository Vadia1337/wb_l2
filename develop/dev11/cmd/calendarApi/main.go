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
	http.HandleFunc("/update_event", h.MiddlewareLogger(h.UpdateEvent))
	http.HandleFunc("/delete_event", h.MiddlewareLogger(h.DeleteEvent))
	http.HandleFunc("/events_for_day", h.MiddlewareLogger(h.GetEventsForDay))
	http.HandleFunc("/events_for_week", h.MiddlewareLogger(h.GetEventsForWeek))
	http.HandleFunc("/events_for_month", h.MiddlewareLogger(h.GetEventsForMonth))

	err := http.ListenAndServe(cfg.HttpServerPort, nil)
	if err != nil {
		return
	}

}
