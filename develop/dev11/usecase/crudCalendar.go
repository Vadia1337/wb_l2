package crudCalendar

import (
	"calendarApi/entity"
	"calendarApi/repository"
	"fmt"
	"time"
)

type CrudCalendar struct {
	repository repository.Repository
}

func NewUseCase(repository repository.Repository) UseCase {
	return &CrudCalendar{repository: repository}
}

func (c *CrudCalendar) CreateEvent(event *entity.Event) (*entity.Event, error) {
	createdEvent, err := c.repository.Create(event)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return createdEvent, nil
}

func (c *CrudCalendar) UpdateEvent(event *entity.Event) (*entity.Event, error) {
	updatedEvent, err := c.repository.Update(event)
	if err == repository.IdNotFound {
		return nil, EventIdNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return updatedEvent, nil
}

func (c *CrudCalendar) DeleteEvent(eventId uint) error {
	err := c.repository.Delete(eventId)
	if err == repository.IdNotFound {
		return EventIdNotFound
	}

	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}

	return nil
}

func (c *CrudCalendar) GetForDay(day time.Time) ([]*entity.Event, error) {
	forDay, err := c.repository.GetForDay(day)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return forDay, nil
}

func (c *CrudCalendar) GetForWeek(firstDay time.Time, lastDay time.Time) ([]*entity.Event, error) {
	forWeek, err := c.repository.GetForWeek(firstDay, lastDay)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return forWeek, nil
}

func (c *CrudCalendar) GetForMonth(month time.Time) ([]*entity.Event, error) {
	forMonth, err := c.repository.GetForMonth(month)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return forMonth, nil
}
