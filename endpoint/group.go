package endpoint

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"../extras"
	"./models"
)

// HandleGetGroupsMessage - Handle Get Groups Message
func HandleGetGroupsMessage(w http.ResponseWriter, r *http.Request) {

	grps, err := ep.db.GetGroups()
	if err != nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	var result []models.GroupSimpleObject = []models.GroupSimpleObject{}
	for i := range grps {
		if grps[i].Enabled {
			membersCount := 0
			members := convertMembersFromDB(grps[i].Members)
			for t := range members {
				if members[t].Enabled == true {
					membersCount++
				}
			}
			if membersCount > 0 {
				result = append(result, models.GroupSimpleObject{
					GroupID:      grps[i].GroupID,
					Title:        grps[i].Title,
					MembersCount: membersCount,
				})
			}
		}
	}

	SuccessResponseInterface(w, r, result)

}

// HandleAdminGetGroupsMessage - Handle Admin Get Groups Message
func HandleAdminGetGroupsMessage(w http.ResponseWriter, r *http.Request) {

	grps, err := ep.db.GetGroups()
	if err != nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	var result []models.GroupObject = []models.GroupObject{}
	for i := range grps {
		result = append(result, models.GroupObject{
			GroupID:   grps[i].GroupID,
			Title:     grps[i].Title,
			Enabled:   grps[i].Enabled,
			Members:   convertMembersFromDB(grps[i].Members),
			CreatedAt: grps[i].CreatedAt,
			UpdatedAt: grps[i].UpdatedAt,
		})
	}

	SuccessResponseInterface(w, r, result)

}

// HandleAdminDeleteGroupMessage - Handle Admin Delete Group Message
func HandleAdminDeleteGroupMessage(w http.ResponseWriter, r *http.Request) {

	var req models.AdminDeleteGroupMessageRequest

	// Body parser
	err := DecodeRequest(w, r, &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	if len(req.GroupID) == 0 {
		ErrorResponse(w, r, errors.New("Group ID required"))
		return
	}

	err = ep.db.DeleteGroup(req.GroupID)
	if err != nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}

	SuccessResponse(w, r)

}

// HandleAdminSetGroupMessage - Handle Admin Set Group Message
func HandleAdminSetGroupMessage(w http.ResponseWriter, r *http.Request) {

	var req models.AdminSetGroupMessageRequest

	// Body parser
	err := DecodeRequest(w, r, &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	// Validations
	if len(req.Group.Members) > 0 {
		for _, m := range req.Group.Members {
			if len(m.Email) == 0 || !extras.ValidateEmail(m.Email) {
				ErrorResponse(w, r, errors.New(m.Email+" E-mail invalid"))
				return
			}
		}
	} else if len(req.Group.Members) == 0 {
		ErrorResponse(w, r, errors.New("At least one member is required"))
		return
	}
	if len(req.Group.Title) == 0 {
		ErrorResponse(w, r, errors.New("Title required"))
		return
	}

	gid, err := ep.db.SetGroup(req.Group.GroupID, req.Group.Title, req.Group.Enabled)
	if err != nil {
		ErrorResponse(w, r, errors.New("Internal error"))
		return
	}
	if len(gid) != 0 {
		members := convertMembersToDB(req.Group.Members)
		grp, err := ep.db.SetMemberGroup(gid, members)
		if err != nil {
			ErrorResponse(w, r, errors.New("Internal error"))
			return
		}
		var result models.AdminSetGroupMessageResponse
		result.Group.GroupID = grp.GroupID
		result.Group.Title = grp.Title
		result.Group.Enabled = grp.Enabled
		result.Group.Members = convertMembersFromDB(grp.Members)
		result.Group.CreatedAt = grp.CreatedAt
		result.Group.UpdatedAt = grp.UpdatedAt
		SuccessResponseInterface(w, r, result)
	} else {
		ErrorResponse(w, r, errors.New("Group wasn't created"))
	}

}

func convertMembersToDB(members []models.MemberGroup) string {
	result, err := json.Marshal(members)
	if err != nil {
		log.Println("convertMembersToDB: ", err)
		return "[]"
	}
	return string(result)
}

func convertMembersFromDB(members string) []models.MemberGroup {
	var result []models.MemberGroup = []models.MemberGroup{}
	err := json.Unmarshal([]byte(members), &result)
	if err != nil {
		log.Println("convertMembersFromDB: ", err)
		return []models.MemberGroup{}
	}
	return result
}
