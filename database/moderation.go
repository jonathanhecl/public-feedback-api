package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/jonathanhecl/public-feedback-api/database/models"
)

func (db DataStore) GetModerationVote(MessageID string) (models.ModerationObject, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.ModerationObject
	q := bson.M{"id": MessageID}
	if err := db.moderation.FindOne(ctx, q).Decode(&msg); err != nil {
		return msg, errors.New("Message not moderated")
	}
	return msg, nil

}

func (db DataStore) SetModerationVote(MessageID string, Email string, IsApprove bool, IP string, UserAgent string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	new := false
	var msg models.ModerationObject
	q := bson.M{"id": MessageID}
	if err := db.moderation.FindOne(ctx, q).Decode(&msg); err != nil {
		new = true
	}
	if new {
		msg.MessageID = MessageID
		msg.Votes = []models.VoteObject{}
		msg.Votes = append(msg.Votes, models.VoteObject{
			Email:     Email,
			IsApprove: IsApprove,
			IP:        IP,
			UserAgent: UserAgent,
			VotedAt:   time.Now(),
		})
		if _, err := db.messages.InsertOne(ctx, msg); err != nil {
			log.Println("Database->SetModerationVote: " + err.Error())
			return err
		}
	} else {
		for v := range msg.Votes {
			if msg.Votes[v].Email == Email {
				return errors.New("Moderator already voted")
			}
		}
		msg.Votes = append(msg.Votes, models.VoteObject{
			Email:     Email,
			IsApprove: IsApprove,
			IP:        IP,
			UserAgent: UserAgent,
			VotedAt:   time.Now(),
		})
		set := bson.M{"$set": bson.M{
			"votes": msg.Votes,
		}}
		if _, err := db.moderation.UpdateOne(ctx, q, set); err != nil {
			log.Println("Database->SetModerationVote: " + err.Error())
			return err
		}
	}
	return nil

}
