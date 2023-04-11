package models

import "github.com/google/uuid"

// Image represents information about an image on the server
type Image struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
	URL  string    `db:"url" json:"url"`
	Size int64     `db:"size" json:"size"`
}
