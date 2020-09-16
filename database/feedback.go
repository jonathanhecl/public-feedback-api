package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"github.com/jonathanhecl/public-feedback-api/database/models"
)

func (db DataStore) GetFeedback(FeedbackID string) (models.FeedbackObject, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.FeedbackObject
	q := bson.M{"id": FeedbackID}
	if err := db.feedback.FindOne(ctx, q).Decode(&msg); err != nil {
		return msg, errors.New("Feedback not found")
	}
	return msg, nil

}

func (db DataStore) NewFeedback(MessageID string, Email string, ToGroup string, Message string, IP string, UserAgent string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.FeedbackObject
	q := bson.M{"message_id": MessageID, "email": Email}
	if err := db.feedback.FindOne(ctx, q).Decode(&msg); err == nil {
		return "", errors.New("Message already replied")
	}
	msg.FeedbackID = uuid.New().String()
	msg.MessageID = MessageID
	msg.Email = Email
	msg.ToGroup = ToGroup
	msg.Message = Message
	msg.IP = IP
	msg.UserAgent = UserAgent
	msg.CreatedAt = time.Now()
	if _, err := db.feedback.InsertOne(ctx, msg); err != nil {
		log.Println("Database->NewFeedback: " + err.Error())
		return "", err
	}
	return msg.FeedbackID, nil

}
