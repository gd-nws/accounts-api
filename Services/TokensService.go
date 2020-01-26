package Services

import (
	"../Models"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToken(user Models.User, expires time.Duration) (string, error) {
	claims := &Models.Claims{
		Id: user.Id,
		StandardClaims: jwt.StandardClaims{
		},
	}

	if expires != 0 {
		expirationTime := time.Now().Add(expires)
		claims.StandardClaims.ExpiresAt = expirationTime.Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := os.Getenv("JWT_KEY")
	return token.SignedString([]byte(key))
}