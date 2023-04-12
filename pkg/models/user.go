// Package models holds the custom data types that the databases hold
package models

import (
	"time"

	"github.com/google/uuid"
)

// BlogUser represents a user that can create blogs
type BlogUser struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Name     string    `db:"name" json:"name"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"password"`
	Admin    bool      `db:"admin" json:"admin"`
	Active   bool      `db:"active" json:"active"`
	Created  time.Time `db:"created" json:"created"`
}
