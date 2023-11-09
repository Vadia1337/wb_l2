package crudCalendar

import "errors"

var EventIdNotFound = errors.New("события с таким ID нет в системе. не удалось обновить событие")
