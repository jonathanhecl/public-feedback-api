package database

import (
	"context"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/jonathanhecl/public-feedback-api/database/models"
)

func (db DataStore) SetTracking(MessageID string, GroupID string, Email string, IP string, UserAgent string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.TrackingObject
	new := false
	q := bson.M{"id": MessageID, "group_id": GroupID}
	if err := db.tracking.FindOne(ctx, q).Decode(&msg); err != nil {
		new = true
	}
	if new {
		msg.MessageID = MessageID
		msg.GroupID = GroupID
		msg.Members = append(msg.Members, models.MemberTrackingObject{
			Email:         Email,
			IP:            IP,
			UserAgent:     UserAgent,
			Readed:        1,
			FirstReadedAt: time.Now(),
			LastReadedAt:  time.Now(),
		})
		msg.SendedAt = msg.SendedAt
		if _, err := db.tracking.InsertOne(ctx, msg); err != nil {
			log.Println("Database->SetTracking: " + err.Error())
			return err
		}
	} else {
		newMember := true
		for v := range msg.Members {
			if msg.Members[v].Email == Email {
				msg.Members[v].Readed++
				msg.Members[v].LastReadedAt = time.Now()
				newMember = false
				break
			}
		}
		if newMember {
			msg.Members = append(msg.Members, models.MemberTrackingObject{
				Email:         Email,
				IP:            IP,
				UserAgent:     UserAgent,
				Readed:        1,
				FirstReadedAt: time.Now(),
				LastReadedAt:  time.Now(),
			})
		}
		set := bson.M{"$set": bson.M{
			"members": msg.Members,
		}}
		if _, err := db.tracking.UpdateOne(ctx, q, set); err != nil {
			log.Println("Database->SetTracking: " + err.Error())
			return err
		}
	}
	return nil

}
