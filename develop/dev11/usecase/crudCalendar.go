package crudCalendar

import "calendarApi/repository"

type CrudCalendar struct {
	repository repository.Repository
}

func NewUseCase(repository repository.Repository) UseCase {
	return &CrudCalendar{repository: repository}
}
