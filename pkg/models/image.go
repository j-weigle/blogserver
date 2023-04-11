package models

// Image represents information about an image on the server
type Image struct {
	ImageName string `db:"imageName" json:"imageName"`
	ImageURL  string `db:"imageURL" json:"imageURL"`
	Size      int64  `db:"size" json:"size"`
}
