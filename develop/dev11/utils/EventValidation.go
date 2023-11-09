package utils

import (
	"calendarApi/entity"
	"net/http"
	"strconv"
	"time"
)

func CreateEventValidate(r *http.Request) *entity.Event {
	name := r.FormValue("name")
	description := r.FormValue("description")
	date := r.FormValue("date")

	if name == "" || description == "" || date == "" {
		return nil
	}

	validDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil
	}

	validEvent := &entity.Event{
		Name:        name,
		Description: description,
		Date:        validDate,
	}

	return validEvent
}

func UpdateEventValidate(r *http.Request) *entity.Event {
	idString := r.FormValue("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	date := r.FormValue("date")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil
	}

	if name == "" || description == "" || date == "" {
		return nil
	}

	validDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil
	}

	validEvent := &entity.Event{
		ID:          uint(id),
		Name:        name,
		Description: description,
		Date:        validDate,
	}

	return validEvent
}

func DeleteEventValidate(r *http.Request) uint {
	idString := r.FormValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0
	}

	return uint(id)
}

func GetDayValidate(r *http.Request) (time.Time, bool) {
	day := r.FormValue("day")

	if day == "" {
		return time.Time{}, false
	}

	validDay, err := time.Parse("2006-01-02", day)
	if err != nil {
		return time.Time{}, false
	}

	return validDay, true
}

func GetWeekValidate(r *http.Request) (time.Time, time.Time, bool) {
	firstDayInWeek := r.FormValue("first_day")
	firstLastInWeek := r.FormValue("last_day")

	if firstDayInWeek == "" || firstLastInWeek == "" {
		return time.Time{}, time.Time{}, false
	}

	firstDay, err := time.Parse("2006-01-02", firstDayInWeek)
	if err != nil {
		return time.Time{}, time.Time{}, false
	}

	lastDay, err := time.Parse("2006-01-02", firstLastInWeek)
	if err != nil {
		return time.Time{}, time.Time{}, false
	}

	if lastDay.Unix() < firstDay.Unix() {
		return time.Time{}, time.Time{}, false
	}

	if (lastDay.Unix() - firstDay.Unix()) > (60 * 60 * 24 * 7) { // хотят диапазон больше недели
		return time.Time{}, time.Time{}, false
	}

	return firstDay, lastDay, true
}

func GetMonthValidate(r *http.Request) (time.Time, bool) {
	month := r.FormValue("month")

	if month == "" {
		return time.Time{}, false
	}

	validMonth, err := time.Parse("2006-01", month)
	if err != nil {
		return time.Time{}, false
	}

	return validMonth, true
}
