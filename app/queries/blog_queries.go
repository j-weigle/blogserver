// Package queries provides access to all the database queries for all models
package queries

import (
	"github.com/j-weigle/blogserver/app/models"
	"github.com/jmoiron/sqlx"
)

// BlogQueries is a type to collate all queries associated with BlogPosts
type BlogQueries struct {
	*sqlx.DB
}

// GetBlogs will select all blog posts from the database
func (q *BlogQueries) GetBlogs() ([]models.BlogPost, error) {
	blogs := []models.BlogPost{}

	query := "SELECT * FROM posts"

	err := q.Get(&blogs, query)

	return blogs, err
}
