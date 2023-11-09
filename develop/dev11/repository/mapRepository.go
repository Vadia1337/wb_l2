package repository

import (
	"calendarApi/entity"
	"sync"
	"time"
)

type MapRepository struct {
	mapEvents map[uint]*entity.Event
	lastId    uint
	mu        *sync.RWMutex
}

func NewMapRepository() Repository {
	return &MapRepository{mapEvents: map[uint]*entity.Event{}, lastId: 0, mu: &sync.RWMutex{}}
}

func (m *MapRepository) Create(event *entity.Event) (*entity.Event, error) {
	m.mu.Lock()

	m.lastId++

	event.ID = m.lastId
	m.mapEvents[m.lastId] = event

	m.mu.Unlock() // defer не использую т.к оверхед в 40 ns, у нас сама операция занимает ns, и может выполняться часто

	return event, nil
}

func (m *MapRepository) Update(event *entity.Event) (*entity.Event, error) {
	m.mu.Lock()

	e, ok := m.mapEvents[event.ID]
	if !ok {
		return nil, IdNotFound
	}

	e.Name = event.Name
	e.Description = event.Description
	e.Date = event.Date

	m.mapEvents[event.ID] = e

	m.mu.Unlock()

	return e, nil
}

func (m *MapRepository) Delete(eventId uint) error {
	m.mu.Lock()
	_, ok := m.mapEvents[eventId]
	if !ok {
		return IdNotFound
	}

	delete(m.mapEvents, eventId)
	m.mu.Unlock()

	return nil
}

func (m *MapRepository) GetForDay(day time.Time) ([]*entity.Event, error) {
	m.mu.RLock()
	eventsForDay := make([]*entity.Event, 0, 2)

	for _, v := range m.mapEvents {
		if v.Date.Year() == day.Year() && v.Date.Month() == day.Month() && v.Date.Day() == day.Day() {
			eventsForDay = append(eventsForDay, v)
		}
	}
	m.mu.RUnlock()
	return eventsForDay, nil
}

func (m *MapRepository) GetForWeek(firstDay time.Time, lastDay time.Time) ([]*entity.Event, error) {
	m.mu.RLock()
	eventsForWeek := make([]*entity.Event, 0, 2)

	for _, v := range m.mapEvents {
		if v.Date.Unix() >= firstDay.Unix() && v.Date.Unix() < lastDay.Unix() {
			eventsForWeek = append(eventsForWeek, v)
		}
	}

	m.mu.RUnlock()

	return eventsForWeek, nil
}

func (m *MapRepository) GetForMonth(month time.Time) ([]*entity.Event, error) {
	m.mu.RLock()
	eventsForMonth := make([]*entity.Event, 0, 2)

	for _, v := range m.mapEvents {
		if v.Date.Year() == month.Year() && v.Date.Month() == month.Month() {
			eventsForMonth = append(eventsForMonth, v)
		}
	}

	m.mu.RUnlock()

	return eventsForMonth, nil
}
