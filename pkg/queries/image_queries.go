// Package queries provides access to all the database queries for all models
package queries

import (
	"github.com/google/uuid"
	"github.com/j-weigle/blogserver/pkg/models"
	"github.com/jmoiron/sqlx"
)

// ImageQueries is a type to collate all queries associated with Images
type ImageQueries struct {
	*sqlx.DB
}

// GetImage will get the image with the given ID
func (q *ImageQueries) GetImage(id uuid.UUID) (models.Image, error) {
	image := models.Image{}

	query := "SELECT * FROM blog_images WITH id = $1"

	err := q.Get(&image, query, id)

	return image, err
}

// GetImages will get all images available
func (q *ImageQueries) GetImages() ([]models.Image, error) {
	images := []models.Image{}

	query := "SELECT * FROM blog_images"

	err := q.Get(&images, query)

	return images, err
}

// AddImage adds an image
func (q *ImageQueries) AddImage(img *models.Image) error {
	query := "INSERT INTO blog_images (id, url, name, size) VALUES (:id, :url, :name, :size)"

	_, err := q.Exec(query, img)

	return err
}

// UpdateImage will update an image to have the provided information in Image
func (q *ImageQueries) UpdateImage(img *models.Image) error {
	query := "UPDATE blog_images SET url = $2, name = $3, size = $4 WHERE id = $1"

	_, err := q.Exec(query, img.ID, img.URL, img.Size)

	return err
}

// DeleteImage will delete an image with the provided ID
func (q *ImageQueries) DeleteImage(id uuid.UUID) error {
	query := "DELETE FROM blog_images WHERE id = $1"

	_, err := q.Exec(query, id)

	return err
}
