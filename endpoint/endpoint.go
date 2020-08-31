package endpoint

import "github.com/jonathanhecl/public-feedback-api/database"

type epStr struct {
	db *database.DataStore
}

var ep *epStr

// InitEndpoint - Init Endpoint
func InitEndpoint(db *database.DataStore) {
	ep = &epStr{
		db,
	}
}
