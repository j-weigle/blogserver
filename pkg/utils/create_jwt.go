// Package utils offers utility functions for parsing information from headers and cookies
package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CreateJWT creates a new JSON Web Token
//
// environment variable JWT_KEY is used to sign the token
// environment variable JWT_EXPIRATION_MINUTES is used to set the expiration
func CreateJWT() (string, error) {
	key := os.Getenv("JWT_KEY")

	duration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_MINUTES"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * time.Duration(duration)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return t, err
}
