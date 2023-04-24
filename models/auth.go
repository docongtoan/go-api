package models

import (
	"goserverapi/config/db"
	_ "goserverapi/config/db"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ValidateAuth(auth Auth) error {
	validate := validator.New()
	return validate.Struct(auth)
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {
	var result string

	sql := ` SELECT token FROM tb_accounts WHERE email = $1 AND password = $2 `

	if err := db.DB.QueryRow(sql, email, password).Scan(&result); err != nil {
		return "", err
	}

	return result, nil

}
