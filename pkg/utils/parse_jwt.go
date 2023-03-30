// utils offers utility functions for parsing information from headers and cookies
package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type JWTMetadata struct {
	Expiration int64
}

func GetJWTMetadata(c *fiber.Ctx) (*JWTMetadata, error) {
	token, err := validateToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expiration := int64(claims["expiration"].(float64))

		return &JWTMetadata{
			Expiration: expiration,
		}, nil
	} else {
		err = fmt.Errorf("invalid token claims")
	}

	return nil, err
}

func validateToken(c *fiber.Ctx) (*jwt.Token, error) {
	rawToken := getTokenFromHeaders(c)
	if rawToken == "" {
		return nil, fmt.Errorf("no token found")
	}

	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func getTokenFromHeaders(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")

	token := strings.Split(bearerToken, " ")
	if len(token) == 2 {
		return token[1]
	}

	return ""
}
