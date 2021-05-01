package models

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"goBoilterplate/config"
	"log"
)

// JwtClaims struct
type JwtClaims struct {
	jwt.StandardClaims
	User User
}

// Login struct
type Login struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

// AuthLogin Login Query
func AuthLogin(email string, password string) *User {
	user := new(User)
	res := config.DB.Where("email = ? and password = ?", email, password).First(&user)
	if res.Error == nil {
		_, err := json.Marshal(&user)
		if err != nil {
			log.Println(err)
		}

		return user
	}
	return nil
}
