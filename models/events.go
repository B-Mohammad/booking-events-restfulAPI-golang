package models

import "time"

type Event struct {
	ID          int
	Name        string
	description string
	Location    string
	UserID      int
	DateTime    time.Time
}

var allEvents []Event = []Event{}

func (e Event) Save() {
	allEvents = append(allEvents, e)
}

func GetAllEvents() []Event {
	return allEvents
}
