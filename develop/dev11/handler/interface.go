package handler

import "net/http"

type RESTHandler interface {
	MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc

	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)

	GetEventsForDay(w http.ResponseWriter, r *http.Request)
	GetEventsForWeek(w http.ResponseWriter, r *http.Request)
	GetEventsForMonth(w http.ResponseWriter, r *http.Request)
}
