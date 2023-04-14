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

	query := "SELECT * FROM blog_posts WHERE id = $1"

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

// GetFullBlogPost gets the blog post with the provided ID (with source)
func (q *BlogQueries) GetFullBlogPost(id uuid.UUID) (models.FullBlogPost, error) {
	blog := models.FullBlogPost{}

	query := "SELECT * FROM blog_posts_full WHERE id = $1"

	err := q.Get(&blog, query, id)

	return blog, err
}

// GetFullBlogPosts gets all blog posts (with source)
func (q *BlogQueries) GetFullBlogPosts() ([]models.FullBlogPost, error) {
	blogs := []models.FullBlogPost{}

	query := "SELECT * FROM blog_posts_full"

	err := q.Get(&blogs, query)

	return blogs, err
}

// CreateBlogPost creates a new blog post with the provided FullBlogPost
func (q *BlogQueries) CreateBlogPost(b *models.FullBlogPost) error {
	query := "INSERT INTO blog_posts_full (id, author, read_time, content, source) VALUES (:id, :author, :read_time, :content, :source)"

	_, err := q.Exec(query, b)

	return err
}

// UpdateBlogPost updates a blog post with the information in the provided FullBlogPost
func (q *BlogQueries) UpdateBlogPost(b *models.FullBlogPost) error {
	query := "UPDATE blog_posts_full SET author = $1, read_time = $2, content = $3, source = $4, updated = $5"

	_, err := q.Exec(query, b.Author, b.ReadTime, b.Content, b.Source, b.Updated)

	return err
}

// DeleteBlogPost deletes a blog post with the given ID
func (q *BlogQueries) DeleteBlogPost(id uuid.UUID) error {
	query := "DELETE FROM blog_posts_full where id = $1"

	_, err := q.Exec(query, id)

	return err
}
