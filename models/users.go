package models

import (
	"errors"

	"bashiri.ir/booking_events_restfulAPI_golang/db"
	"bashiri.ir/booking_events_restfulAPI_golang/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	saveUQuery := `INSERT INTO users(email, password) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(saveUQuery)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedP, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedP)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	u.ID = id

	return err
}

func (u User) CheckCredential() error {
	selectQ := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(selectQ, u.Email)

	var hPass string
	err := row.Scan(&hPass)
	if err != nil {
		return err
	}

	result := utils.ComparePassword(u.Password, hPass)
	if result {
		return nil
	}
	return errors.New("password is incorrect!")
}
