package handler

import (
	"calendarApi/usecase"
	"calendarApi/utils"
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
	if r.Method != http.MethodPost {
		utils.SendJsonResponse(w, 405, "разрешены только POST запросы")
		return
	}

	validEvent := utils.CreateEventValidate(r)
	if validEvent == nil {
		utils.SendJsonResponse(w, 400, "ваши данные невалидны")
		return
	}

	eventWithID, err := h.crudCalendarUC.CreateEvent(validEvent)
	if err != nil {
		utils.SendJsonResponse(w, 503, "ошибка бизнес-логики")
		return
	}

	utils.SendJsonResponse(w, 200, eventWithID)
}

func (h *DefaultHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendJsonResponse(w, 405, "разрешены только POST запросы")
		return
	}

	validEvent := utils.UpdateEventValidate(r)
	if validEvent == nil {
		utils.SendJsonResponse(w, 400, "ваши данные невалидны")
		return
	}

	updateEvent, err := h.crudCalendarUC.UpdateEvent(validEvent)
	if err == crudCalendar.EventIdNotFound {
		utils.SendJsonResponse(w, 400, "ваши данные невалидны")
		return
	}

	if err != nil {
		utils.SendJsonResponse(w, 503, "ошибка бизнес-логики")
		return
	}

	utils.SendJsonResponse(w, 200, updateEvent)
}

func (h *DefaultHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendJsonResponse(w, 405, "разрешены только POST запросы")
		return
	}

	eventId := utils.DeleteEventValidate(r)
	if eventId == 0 {
		utils.SendJsonResponse(w, 400, "ваши данные невалидны")
		return
	}

	err := h.crudCalendarUC.DeleteEvent(eventId)
	if err == crudCalendar.EventIdNotFound {
		utils.SendJsonResponse(w, 400, "ваши данные невалидны")
		return
	}

	if err != nil {
		utils.SendJsonResponse(w, 503, "ошибка бизнес-логики")
		return
	}

	utils.SendJsonResponse(w, 200, eventId)
}

func (h *DefaultHandler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendJsonResponse(w, 405, "разрешены только GET запросы")
		return
	}

	validDay, isValid := utils.GetDayValidate(r)
	if !isValid {
		utils.SendJsonResponse(w, 400, "ваши данные невалидны")
		return
	}

	eventsForDay, err := h.crudCalendarUC.GetForDay(validDay)
	if err != nil {
		utils.SendJsonResponse(w, 503, "ошибка бизнес-логики")
		return
	}

	utils.SendJsonResponse(w, 200, eventsForDay)
}

func (h *DefaultHandler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendJsonResponse(w, 405, "разрешены только GET запросы")
		return
	}

	firstDay, lastDay, isValid := utils.GetWeekValidate(r)
	if !isValid {
		utils.SendJsonResponse(w, 400, "ваши данные невалидны")
		return
	}

	eventsForWeek, err := h.crudCalendarUC.GetForWeek(firstDay, lastDay)
	if err != nil {
		utils.SendJsonResponse(w, 503, "ошибка бизнес-логики")
		return
	}

	utils.SendJsonResponse(w, 200, eventsForWeek)
}

func (h *DefaultHandler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendJsonResponse(w, 405, "разрешены только GET запросы")
		return
	}

	validMonth, isValid := utils.GetMonthValidate(r)
	if !isValid {
		utils.SendJsonResponse(w, 400, "ваши данные невалидны")
		return
	}

	eventsForMonth, err := h.crudCalendarUC.GetForMonth(validMonth)
	if err != nil {
		utils.SendJsonResponse(w, 503, "ошибка бизнес-логики")
		return
	}

	utils.SendJsonResponse(w, 200, eventsForMonth)
}

func (h *DefaultHandler) MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		log.Printf("%s %s %s Body: %s", r.Method, r.RemoteAddr, r.URL.Path, r.Form)
	}
}
