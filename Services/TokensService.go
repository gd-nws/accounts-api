package Services

import (
	"../Models"
	"errors"
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

func VerifyToken(token string, claims *Models.Claims) error {
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		key := os.Getenv("JWT_KEY")
		return []byte(key), nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		return errors.New("invalid token")
	}

	return nil
}