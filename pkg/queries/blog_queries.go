// Package queries provides access to all the database queries for all models
package queries

import (
	"github.com/google/uuid"
	"github.com/j-weigle/blogserver/pkg/models"
	"github.com/jmoiron/sqlx"
)

// BlogQueries is a type to collate all queries associated with BlogPosts
type BlogQueries struct {
	*sqlx.DB
}

// GetBlogPost gets the blog post with the provided ID
func (q *BlogQueries) GetBlogPost(id uuid.UUID) (models.BlogPost, error) {
	blog := models.BlogPost{}

	query := "SELECT * FROM blog_post WHERE id = $1"

	err := q.Get(&blog, query, id)

	return blog, err
}

// GetBlogPosts gets all blog posts
func (q *BlogQueries) GetBlogPosts() ([]models.BlogPost, error) {
	blogs := []models.BlogPost{}

	query := "SELECT * FROM blog_posts"

	err := q.Get(&blogs, query)

	return blogs, err
}

// CreateBlogPost creates a new blog post with the provided BlogPostWithSource
func (q *BlogQueries) CreateBlogPost(b *models.BlogPostWithSource) error {
	query1 := "INSERT INTO blog_posts (id, author, read_time, content) VALUES (:id, :author, :read_time, :content)"
	query2 := "INSERT INTO blog_post_source (id, source) VALUES ($1, $2)"

	// use transaction to rollback on error on either post or source
	tx, err := q.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query1, b.BlogPost)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(query2, b.ID, b.Source)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	return err
}
