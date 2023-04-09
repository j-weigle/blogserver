// Package middleware provides middleware functions to use in a go fiber application
package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// UseJWT creates a JWT middleware to apply to routes in a go fiber application
func UseJWT() func(*fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey:    []byte(os.Getenv("JWT_KEY")),
		SigningMethod: jwt.SigningMethodHS512.Alg(),
	}

	return jwtware.New(config)
}
