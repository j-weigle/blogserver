// Package main is to run the server for interacting with a blog database
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// BlogPost represents a blog post with image links
type BlogPost struct {
	ID      int      `json:"id"`
	Content string   `json:"content"`
	Images  []string `json:"images"`
}

func handleGet(c *fiber.Ctx, _ *sql.DB) error {
	// example for reading from database: TODO make table for blog posts
	//type Todo struct {
	//	ID   int      `json:"id"`
	//	Item string   `json:"item"`
	//}
	// var res Todo
	// var todos []Todo
	// rows, err := db.Query("SELECT * FROM todos")
	// defer rows.Close()
	//
	//	if err != nil {
	//		c.JSON(`"error": "issue getting data from SELECT"`)
	//	}
	//
	//	for rows.Next() {
	//		rows.Scan(&res.Id, &res.Item)
	//		todos = append(todos, res)
	//	}
	//
	//	return c.JSON(fiber.Map{
	//		"todos": todos,
	//	})
	return c.SendString("Hello")
}

func handlePost(c *fiber.Ctx, _ *sql.DB) error {
	return c.SendString("Hello")
}

func handlePostImage(c *fiber.Ctx, _ *sql.DB) error {

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	uniqueID := uuid.New()

	fmt.Println(uniqueID)

	fileName := strings.Replace(uniqueID.String(), "-", "", -1)

	fmt.Println(file.Filename)

	fileExt := strings.Split(file.Filename, ".")[1]

	image := fmt.Sprintf("%s.%s", fileName, fileExt)

	err = c.SaveFile(file, fmt.Sprintf("/srv/blog/images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	imageURL := fmt.Sprintf("http://localhost:3100/images/%s", image)

	data := map[string]interface{}{
		"imageName": image,
		"imageURL":  imageURL,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func handlePut(c *fiber.Ctx, _ *sql.DB) error {
	return c.SendString("Hello")
}

func handleDelete(c *fiber.Ctx, _ *sql.DB) error {
	return c.SendString("Hello")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	postgresUser := os.Getenv("POSTGRESUSER")
	postgresPass := os.Getenv("POSTGRESPASS")
	connStr := fmt.Sprintf("postgresql://%s:%s@127.0.0.1:5432/todoapp?sslmode=disable", postgresUser, postgresPass)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Static("/images", "/srv/blog/images")

	app.Get("/", func(c *fiber.Ctx) error {
		return handleGet(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return handlePost(c, db)
	})

	app.Put("/", func(c *fiber.Ctx) error {
		return handlePut(c, db)
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		return handleDelete(c, db)
	})

	// image group
	image := app.Group("/image", func(c *fiber.Ctx) error {
		fmt.Println("IMAGE")
		// TODO add authentication here so image posting is protected
		return c.Next()
	})

	image.Post("/", func(c *fiber.Ctx) error {
		return handlePostImage(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3100"
	}
	log.Fatalln(app.Listen(fmt.Sprintf("localhost:%v", port)))
}
