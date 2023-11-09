package repository

import (
	"calendarApi/entity"
	"time"
)

type Repository interface {
	Create(entity *entity.Event) (*entity.Event, error)
	Update(event *entity.Event) (*entity.Event, error)
	Delete(eventId uint) error
	GetForDay(day time.Time) ([]*entity.Event, error)
	GetForWeek(firstDay time.Time, lastDay time.Time) ([]*entity.Event, error)
	GetForMonth(month time.Time) ([]*entity.Event, error)
}
