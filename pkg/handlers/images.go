// Package handlers defines the method handlers for different routes
package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/j-weigle/blogserver/pkg/utils"
)

// GetImages returns a list of all images on the server's metadata
func GetImages(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

// GetImage returns a single image's metadata
func GetImage(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

// CreateImage creates a new image and adds its metadata to the database
func CreateImage(c *fiber.Ctx) error {
	// db, err := database.OpenDB()

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	metadata, err := utils.ParseJWTClaims(claims)

	if metadata.Expiration < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"error": fiber.Map{
				"message": "unauthorized, expired token",
			},
		})
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": "error", "error": fiber.Map{
			"message": "Server error",
		}})
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.SendStatus(500)
		return c.JSON(fiber.Map{"status": "error", "error": fiber.Map{
			"message": "Server error",
		}})
	}
	defer file.Close()

	d, err := io.ReadAll(file)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(fiber.Map{"status": "error", "error": fiber.Map{
			"message": "Server error",
		}})
	}

	fType := http.DetectContentType(d)

	if !isAcceptableFileType(fType) {
		c.SendStatus(415)
		return c.JSON(fiber.Map{"status": "error", "error": fiber.Map{
			"message":  "valid file types: jpeg, png",
			"fileType": fType,
		}})
	}

	uniqueID := uuid.New()

	fileName := strings.Replace(uniqueID.String(), "-", "", -1)

	fileExt := strings.Split(fileHeader.Filename, ".")[1]

	image := fmt.Sprintf("%s.%s", fileName, fileExt)

	err = c.SaveFile(fileHeader, fmt.Sprintf("/srv/blog/images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	imageURL := fmt.Sprintf("http://129.80.4.181/images/%s", image)

	data := map[string]interface{}{
		"imageName": image,
		"imageURL":  imageURL,
		"header":    fileHeader.Header,
		"size":      fileHeader.Size,
	}

	// TODO insert into database as well

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

// UpdateImage replaces and image and updates its metadata in the database
func UpdateImage(c *fiber.Ctx) error {
	// is protected by JWT middleware
	return c.SendString("HELLO")
}

// DeleteImage deletes an image from the server and removes its metadata from the database
func DeleteImage(c *fiber.Ctx) error {
	// is protected by JWT middleware
	return c.SendString("HELLO")
}

func isAcceptableFileType(ft string) bool {
	acceptedFileTypes := []string{"image/jpeg", "image/png"}

	for _, t := range acceptedFileTypes {
		if t == ft {
			return true
		}
	}
	return false
}
