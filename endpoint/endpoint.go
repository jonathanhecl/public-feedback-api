package endpoint

import "github.com/jonathanhecl/public-feedback-api/database"

type epStr struct {
	db             *database.DataStore
	brandTitle     string
	minModApproves int
}

var ep *epStr

// InitEndpoint - Init Endpoint
func InitEndpoint(db *database.DataStore, brandTitle string, minModApproves int) {
	ep = &epStr{
		db,
		brandTitle,
		minModApproves,
	}
}
