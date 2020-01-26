package Services

import (
	"../Models"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToken(user Models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Models.Claims{
		Id: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := os.Getenv("JWT_KEY")
	return token.SignedString([]byte(key))
}