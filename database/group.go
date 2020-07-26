package database

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"./models"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db DataStore) GetGroup(GroupID string) (models.GroupObject, error) {

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var grp models.GroupObject
	q := bson.M{"id": GroupID}
	if err := db.groups.FindOne(ctx, q).Decode(&grp); err != nil {
		log.Println("Database->GetGroup: " + err.Error())
		return grp, errors.New("Group not found")
	}
	return grp, nil

}

func (db DataStore) GetGroups() ([]models.GroupObject, error) {

	var results []models.GroupObject = []models.GroupObject{}
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.groups.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		log.Println("Database->GetGroups-Find: " + err.Error())
		return results, err
	}
	for cur.Next(ctx) {
		var grp models.GroupObject
		if err := cur.Decode(&grp); err != nil {
			log.Println("Database->GetGroups-Decode: " + err.Error())
			return nil, err
		}
		results = append(results, grp)
	}
	if err := cur.Err(); err != nil {
		log.Println("Database->GetGroups: " + err.Error())
		return nil, err
	}
	cur.Close(ctx)
	return results, nil

}

func (db DataStore) DeleteGroup(GroupID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var grp models.GroupObject
	q := bson.M{"id": GroupID}
	if err := db.groups.FindOne(ctx, q).Decode(&grp); err != nil {
		return errors.New("Group not found")
	}
	if _, err := db.groups.DeleteOne(ctx, q); err != nil {
		log.Println("Database->DeleteOne: " + err.Error())
		return err
	}
	return nil

}

func (db DataStore) SetGroup(GroupID string, Title string, Enabled bool) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var grp models.GroupObject
	if len(GroupID) != 0 {
		q := bson.M{"id": GroupID}
		if err := db.groups.FindOne(ctx, q).Decode(&grp); err != nil {
			return "", errors.New("Group not found")
		}
		grp.Title = Title
		grp.Enabled = Enabled
		set := bson.M{"$set": bson.M{
			"title":      grp.Title,
			"enabled":    grp.Enabled,
			"updated_at": time.Now(),
		}}
		if _, err := db.groups.UpdateOne(ctx, q, set); err != nil {
			log.Println("Database->SetGroup: " + err.Error())
			return "", err
		}
	} else {
		grp.GroupID = uuid.New().String()
		grp.Title = Title
		grp.Enabled = Enabled
		grp.Members = "[]"
		grp.CreatedAt = time.Now()
		grp.UpdatedAt = grp.CreatedAt
		if _, err := db.groups.InsertOne(ctx, grp); err != nil {
			log.Println("Database->SetGroup: " + err.Error())
			return "", err
		}
	}
	return grp.GroupID, nil

}

func (db DataStore) SetMemberGroup(GroupID string, MembersJSON string) (models.GroupObject, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var grp models.GroupObject
	q := bson.M{"id": GroupID}
	if err := db.groups.FindOne(ctx, q).Decode(&grp); err != nil {
		log.Println("Database->SetMemberGroup: ", err)
		return grp, errors.New("Group not found")
	}
	var members []models.MemberGroupObject
	if err := json.Unmarshal([]byte(grp.Members), &members); err != nil {
		log.Println("Database->SetMemberGroup-Unmarshal[1]:", err)
		return grp, err
	}
	var Members []models.MemberGroupObject
	if err := json.Unmarshal([]byte(MembersJSON), &Members); err != nil {
		log.Println("Database->SetMemberGroup-Unmarshal[2]:", err)
		return grp, err
	}
	save := false
	for i := range Members {
		Members[i].Email = strings.ToLower(Members[i].Email)
		new := true
		for t := range members {
			if members[t].Email == Members[i].Email {
				new = false
				if members[t].Name != Members[i].Name || members[t].Enabled != Members[i].Enabled {
					members[t].Name = Members[i].Name
					members[t].Enabled = Members[i].Enabled
					save = true
				}
			}
		}
		if new {
			members = append(members, models.MemberGroupObject{
				Email:   Members[i].Email,
				Name:    Members[i].Name,
				Enabled: Members[i].Enabled,
			})
			save = true
		}
	}
	if save {
		jsonMembers, _ := json.Marshal(members)
		grp.Members = string(jsonMembers)
		q2 := bson.M{"$set": bson.M{"members": grp.Members, "updated_at": time.Now()}}
		if err := db.groups.FindOneAndUpdate(ctx, q, q2).Err(); err != nil {
			log.Println("Database->SetMemberGroup: ", err)
			return grp, err
		}
	}
	return grp, nil

}
