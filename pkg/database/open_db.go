// Package database gives access to opening and querying available databases
package database

import "github.com/j-weigle/blogserver/app/queries"

// Queries collects all possible app queries into one place
type Queries struct {
	*queries.BlogQueries
	*queries.ImageQueries
}

// OpenDB opens any db connections and returns available queries
func OpenDB() (*Queries, error) {
	db, err := OpenPostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		BlogQueries:  &queries.BlogQueries{DB: db},
		ImageQueries: &queries.ImageQueries{DB: db},
	}, nil
}
