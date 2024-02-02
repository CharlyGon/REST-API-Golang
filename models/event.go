package models

import "time"

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e *Event) Save() {
	// db.Create(&e)
	events = append(events, *e)
}

func GetEvents() []Event {
	return events
}
