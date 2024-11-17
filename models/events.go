package models

import "time"

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	UserID      int
	DateTime    time.Time `binding:"required"`
}

var allEvents []Event = []Event{}

func (e Event) Save() {
	allEvents = append(allEvents, e)
}

func GetAllEvents() []Event {
	return allEvents
}
