package Models

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id int       `json:"id"`
	jwt.StandardClaims
}