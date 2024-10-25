package middleware

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Secret key used to sign tokens
var jwtSecret []byte

// InitializeJWT initializes the secret key for JWT
func InitializeJWT(secret string) {
	jwtSecret = []byte(secret)
}

// GenerateJWT creates a new JWT token for a user
func GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"sub": strconv.Itoa(userID),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates a JWT and returns the claims
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
