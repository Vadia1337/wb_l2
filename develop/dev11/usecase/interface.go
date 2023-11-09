package crudCalendar

import (
	"calendarApi/entity"
	"time"
)

type UseCase interface {
	CreateEvent(event *entity.Event) (*entity.Event, error)
	UpdateEvent(event *entity.Event) (*entity.Event, error)
	DeleteEvent(eventId uint) error
	GetForDay(day time.Time) ([]*entity.Event, error)
	GetForWeek(firstDay time.Time, lastDay time.Time) ([]*entity.Event, error)
	GetForMonth(month time.Time) ([]*entity.Event, error)
}
