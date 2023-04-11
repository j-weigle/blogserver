// Package routes handles assigning public, private, and unfound route handlers
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j-weigle/blogserver/pkg/handlers"
	"github.com/j-weigle/blogserver/pkg/middleware"
)

// PrivateRoutes assigns handlers for routes that are protected
func PrivateRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	postRoute := route.Group("/post", middleware.UseJWT())
	postRoute.Post("/", handlers.CreatePost)
	postRoute.Put("/", handlers.UpdatePost)
	postRoute.Delete("/", handlers.DeletePost)

	imageRoute := route.Group("/image", middleware.UseJWT())
	//imageRoute := route.Group("/image")
	imageRoute.Post("/", handlers.CreateImage)
	imageRoute.Put("/", handlers.UpdateImage)
	imageRoute.Delete("/", handlers.DeleteImage)

	editorRoute := route.Group("/editor", middleware.UseJWT())
	editorRoute.Get("/post", handlers.GetPostWithSource)
}
