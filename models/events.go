package models

import (
	"time"

	"bashiri.ir/booking_events_restfulAPI_golang/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	UserID      int64
	DateTime    time.Time `binding:"required"`
}

func (e *Event) Save() error {
	insertQuery := `
	INSERT INTO events (name, description, location, dateTime, userId) 
	Values (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	e.ID, err = result.LastInsertId()
	return err
}

func GetAllEvents() ([]Event, error) {
	selectQuery := `SELECT * FROM events`
	rows, err := db.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEvent(eventId int64) (*Event, error) {
	singleSelectQuery := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(singleSelectQuery, eventId)
	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e Event) UpdateEvent() error {
	updateQuery := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ? 
	WHERE id = ?`
	stmt, err := db.DB.Prepare(updateQuery)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e Event) DeleteEvent() error {
	deleteQuery := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(deleteQuery)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func (e Event) RegisterToEvent(userId int64) error {
	insertQuery := `INSERT INTO erg(userId, eventId)
	VALUES(?, ?)`

	stmt, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, e.ID)
	return err
}

func (e Event) DeleteRegister(userId int64) error {
	insertQuery := `DELETE FROM erg
	WHERE userId = ?  AND eventId = ?`

	stmt, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, e.ID)
	return err
}
