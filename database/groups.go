package database

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"

	"github.com/jonathanhecl/public-feedback-api/database/models"
	"github.com/jonathanhecl/public-feedback-api/extras"
)

var Groups []*models.GroupObject

func (ds DataStore) LoadGroups() {

	conf, err := google.JWTConfigFromJSON([]byte(ds.googleCert), spreadsheet.Scope)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet(ds.groupSpreadsheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	var newGroups []*models.GroupObject
	regTitle := regexp.MustCompile(`\[(\w*)\] ?(.*)`)
	for s := range spreadsheet.Sheets {
		m := regTitle.FindStringSubmatch(spreadsheet.Sheets[s].Properties.Title)
		if len(m) == 3 {
			group := &models.GroupObject{
				GroupID: m[1],
				Label:   m[2],
			}
			sheet := spreadsheet.Sheets[s]
			for _, row := range sheet.Rows {
				member := models.MemberGroup{}
				for _, cell := range row {
					if cell.Row > 0 {
						if cell.Column == 0 {
							if extras.ValidateEmail(cell.Value) {
								member.Email = cell.Value
							}
						} else if cell.Column == 1 && len(member.Email) > 0 {
							member.Name = cell.Value
						}
					}
				}
				if len(member.Email) > 0 {
					group.Members = append(group.Members, member)
				}
			}
			group.UpdatedAt = time.Now()
			newGroups = append(newGroups, group)
		}
	}
	Groups = newGroups

	//go DebugGroups()

}

func DebugGroups() {

	for i := range Groups {
		fmt.Println()
		fmt.Println("GroupID: ", Groups[i].GroupID)
		fmt.Println("Label: ", Groups[i].Label)
		fmt.Println("UpdatedAt: ", Groups[i].UpdatedAt.String())
		fmt.Println("Members: ", len(Groups[i].Members))
		for m := range Groups[i].Members {
			fmt.Println("Member: ", Groups[i].Members[m].Name, "<", Groups[i].Members[m].Email, ">")
		}
	}
}

func (ds DataStore) GetGroup(GroupID string) (*models.GroupObject, error) {

	if len(Groups) == 0 {
		return nil, errors.New("No groups settings")
	}
	for i := range Groups {
		if Groups[i].GroupID == GroupID {
			return Groups[i], nil
		}
	}
	return nil, errors.New("Group not found")

}

func (ds DataStore) GetGroups() ([]*models.GroupObject, error) {

	if len(Groups) == 0 {
		return nil, errors.New("No groups settings")
	}
	return Groups, nil

}
