// Package models holds the custom data types that the databases hold
package models

import (
	"time"

	"github.com/google/uuid"
)

// BlogPost represents a blog post with some extra metadata
type BlogPost struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Created  time.Time `db:"created" json:"created"`
	Updated  time.Time `db:"updated" json:"updated"`
	ReadTime int       `db:"read_time" json:"readTime"`
	Author   *BlogUser `db:"author" json:"author"`
	Content  string    `db:"content" json:"content"`
}

// BlogPostWithSource includes the source representation of the html from a BlogPost in markdown
type BlogPostWithSource struct {
	BlogPost
	Source string `db:"source" json:"source"`
}
