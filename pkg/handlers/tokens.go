// Package handlers defines the method handlers for different routes
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j-weigle/blogserver/pkg/utils"
)

// GetNewToken returns a new JSON Web Token to a client
func GetNewToken(c *fiber.Ctx) error {

	// TODO login validation with refresh token checking first

	token, err := utils.CreateJWT()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error": fiber.Map{
				"message": err.Error(),
			},
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"message":      "",
			"access_token": token,
		},
	})
}
