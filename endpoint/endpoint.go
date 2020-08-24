package endpoint

import "../database"

type epStr struct {
	db             *database.DataStore
	minModApproves int
}

var ep *epStr

// InitEndpoint - Init Endpoint
func InitEndpoint(db *database.DataStore, minModApproves int) {
	ep = &epStr{
		db,
		minModApproves,
	}
}
