package handler

import (
	crudCalendar "calendarApi/usecase"
	"log"
	"net/http"
)

type DefaultHandler struct {
	crudCalendarUC crudCalendar.UseCase
}

func NewHandler(useCase crudCalendar.UseCase) RESTHandler {
	return &DefaultHandler{crudCalendarUC: useCase}
}

func (h *DefaultHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodPost {
	//	http.Error(w, "", http.StatusMethodNotAllowed)
	//	return
	//}

}

func (h *DefaultHandler) MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		log.Printf("%s %s %s Body: %s", r.Method, r.RemoteAddr, r.URL.Path, r.Body)
	}
}
