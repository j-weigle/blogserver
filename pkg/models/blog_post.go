// Package models holds the custom data types that the databases hold
package models

import (
	"time"

	"github.com/google/uuid"
)

// BlogPost represents a blog post with some extra metadata
type BlogPost struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Author   *BlogUser `db:"author" json:"author"`
	ReadTime int       `db:"read_time" json:"readTime"`
	Content  string    `db:"content" json:"content"`
	Created  time.Time `db:"created" json:"created"`
	Updated  time.Time `db:"updated" json:"updated"`
}

// FullBlogPost includes the source representation of the html from a BlogPost in markdown
type FullBlogPost struct {
	BlogPost
	Source string `db:"source" json:"source"`
}
