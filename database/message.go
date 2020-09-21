package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jonathanhecl/public-feedback-api/database/models"
	"github.com/jonathanhecl/public-feedback-api/extras"
)

func (db DataStore) SetMessageSended(MessageID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.MessageObject
	q := bson.M{"id": MessageID}
	if err := db.messages.FindOne(ctx, q).Decode(&msg); err != nil {
		return errors.New("Message not found")
	}
	set := bson.M{"$set": bson.M{
		"sended_at": time.Now(),
		"closed_at": time.Now(),
	}}
	if _, err := db.messages.UpdateOne(ctx, q, set); err != nil {
		log.Println("Database->SetMessageSended: " + err.Error())
		return err
	}
	return nil

}

func (db DataStore) SetMessageClosed(MessageID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.MessageObject
	q := bson.M{"id": MessageID}
	if err := db.messages.FindOne(ctx, q).Decode(&msg); err != nil {
		return errors.New("Message not found")
	}
	set := bson.M{"$set": bson.M{
		"closed_at": time.Now(),
	}}
	if _, err := db.messages.UpdateOne(ctx, q, set); err != nil {
		log.Println("Database->SetMessageClosed: " + err.Error())
		return err
	}
	return nil

}

func (db DataStore) GetMessagesPending() ([]models.MessageObject, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msgs []models.MessageObject
	or := bson.M{"$or": bson.M{
		"confirmed_at": time.Unix(0, 0),
		"sended_at":    time.Unix(0, 0),
	}}
	cur, err := db.messages.Find(ctx, or, options.Find())
	if err != nil {
		return msgs, errors.New("Messages not found")
	}
	for cur.Next(ctx) {
		var elem models.MessageObject
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		if elem.ClosedAt == time.Unix(0, 0) { // Ignore closed
			msgs = append(msgs, elem)
		}
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return msgs, nil

}

func (db DataStore) GetMessage(MessageID string) (models.MessageObject, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.MessageObject
	q := bson.M{"id": MessageID}
	if err := db.messages.FindOne(ctx, q).Decode(&msg); err != nil {
		return msg, errors.New("Message not found")
	}

	return msg, nil

}

func (db DataStore) ConfirmMessage(MessageID string, ConfirmationCode string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.MessageObject
	q := bson.M{"id": MessageID}
	if err := db.messages.FindOne(ctx, q).Decode(&msg); err != nil {
		return errors.New("Message not found")
	}
	if msg.ConfirmedAt.Unix() > 0 {
		return errors.New("Message already confirmed")
	}
	if msg.ConfirmationCode != ConfirmationCode {
		return errors.New("Wrong confirmation code")
	}
	q = bson.M{"id": MessageID}
	set := bson.M{"$set": bson.M{
		"confirmed_at": time.Now(),
	}}
	if _, err := db.messages.UpdateOne(ctx, q, set); err != nil {
		log.Println("Database->ConfirmMessage: " + err.Error())
		return err
	}
	return nil

}

func (db DataStore) NewMessage(Email string, Name string, Message string, GroupID string, IP string, UserAgent string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.MessageObject
	msg.MessageID = uuid.New().String()
	msg.Email = Email
	msg.Name = Name
	msg.Message = Message
	msg.ToGroup = GroupID
	msg.IP = IP
	msg.UserAgent = UserAgent
	msg.ConfirmationCode = extras.RandomCode()
	msg.CreatedAt = time.Now()
	msg.ConfirmedAt = time.Unix(0, 0) // Not confirmed
	msg.SendedAt = time.Unix(0, 0)    // Not sended
	msg.ClosedAt = time.Unix(0, 0)    // Not closed
	if _, err := db.messages.InsertOne(ctx, msg); err != nil {
		log.Println("Database->NewMessage: " + err.Error())
		return "", err
	}
	return msg.MessageID, nil

}
