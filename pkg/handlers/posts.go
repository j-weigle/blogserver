// Package handlers defines the method handlers for different routes
package handlers

import "github.com/gofiber/fiber/v2"

// GetPosts gets a list of all Blog Posts
func GetPosts(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

// GetPost gets a single blog post
func GetPost(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

// GetPostWithSource gets a single blog post with its source markdown
func GetPostWithSource(c *fiber.Ctx) error {
	// is protected by JWT middleware
	return c.SendString("HELLO")
}

// CreatePost creates a new blog post
func CreatePost(c *fiber.Ctx) error {
	// is protected by JWT middleware
	return c.SendString("HELLO")
}

// UpdatePost updates a blog post to contain the new contents provided
func UpdatePost(c *fiber.Ctx) error {
	// is protected by JWT middleware
	return c.SendString("HELLO")
}

// DeletePost deletes a blog post
func DeletePost(c *fiber.Ctx) error {
	// is protected by JWT middleware
	return c.SendString("HELLO")
}
