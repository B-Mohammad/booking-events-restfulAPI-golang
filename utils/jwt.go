package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix()})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedT, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid method!")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("token could  not parsed!")
	}

	if !parsedT.Valid {
		return 0, errors.New("token not valid!")
	}

	data, ok := parsedT.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("claims could not parsed!")
	}

	// email := data["email"].(string)
	userId := int64(data["userId"].(float64))

	return userId, nil
}
