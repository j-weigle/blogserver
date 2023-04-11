// Package routes handles assigning public, private, and unfound route handlers
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j-weigle/blogserver/pkg/handlers"
)

// PublicRoutes assigns handlers for routes that are not protected
func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/posts", handlers.GetPosts)
	route.Get("/post/:id", handlers.GetPost)
	route.Get("/images", handlers.GetImages)
	route.Get("/image/:id", handlers.GetImage)
	route.Post("/token/jwt/new", handlers.GetNewToken)
}
