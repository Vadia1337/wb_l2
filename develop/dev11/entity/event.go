package entity

import "time"

type Event struct {
	ID          uint      `json:"ID"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Date        time.Time `json:"Date"`
}
