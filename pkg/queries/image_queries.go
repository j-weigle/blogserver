// Package queries provides access to all the database queries for all models
package queries

import (
	"github.com/j-weigle/blogserver/pkg/models"
	"github.com/jmoiron/sqlx"
)

// ImageQueries is a type to collate all queries associated with Images
type ImageQueries struct {
	*sqlx.DB
}

// GetImages will select all images from the database
func (q *ImageQueries) GetImages() ([]models.Image, error) {
	images := []models.Image{}

	query := "SELECT * FROM images"

	err := q.Get(&images, query)

	return images, err
}
