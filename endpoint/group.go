package endpoint

import (
	"errors"
	"net/http"

	"./models"
)

// HandleGetGroupsMessage - Handle Get Groups Message
func HandleGetGroupsMessage(w http.ResponseWriter, r *http.Request) {

	var res models.GetGroupsMessageResponse

	grps, err := ep.db.GetGroups()
	if err != nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}
	grs := []models.GroupSimpleObject{}
	for i := range grps {
		if len(grps[i].Label) > 0 && grps[i].GroupID != "MOD" {
			grs = append(grs, models.GroupSimpleObject{
				GroupID:      grps[i].GroupID,
				Label:        grps[i].Label,
				MembersCount: len(grps[i].Members),
			})
		}
	}
	if len(grs) > 0 {
		res.Groups = grs
	} else {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	SuccessResponseInterface(w, r, res)

}
