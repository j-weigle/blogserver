// Package main is to run the server for interacting with a blog database
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/j-weigle/blogserver/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Static("/images", "/srv/blog/images")

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	// TODO routes.FourOhFour(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3100"
	}
	log.Fatalln(app.Listen(fmt.Sprintf("localhost:%v", port)))
}
