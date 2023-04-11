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
	Author   string    `db:"author" json:"author"`
	ReadTime string    `db:"readTime" json:"readTime"`
	Html     string    `db:"html" json:"html"`
}

// BlogPostWithSource includes the source representation of the html from a BlogPost in markdown
type BlogPostWithSource struct {
	BlogPost *BlogPost `json:"blogPost"`
	Source   string    `db:"source" json:"source"`
}
