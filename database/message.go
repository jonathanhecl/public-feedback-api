package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"

	"../extras"
	"./models"
)

func (db DataStore) GetMessage(MessageID string) (models.MessageObject, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.MessageObject
	q := bson.M{"id": MessageID}
	if err := db.messages.FindOne(ctx, q).Decode(&msg); err == nil {
		return msg, errors.New("Message not found")
	}
	return msg, nil

}

func (db DataStore) ConfirmMessage(MessageID string, ConfirmationCode string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.MessageObject
	q := bson.M{"id": MessageID}
	if err := db.messages.FindOne(ctx, q).Decode(&msg); err == nil {
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

func (db DataStore) NewMessage(Email string, Message string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var msg models.MessageObject
	msg.MessageID = uuid.New().String()
	msg.Email = Email
	msg.Message = Message
	msg.ConfirmationCode = extras.RandomCode()
	msg.CreatedAt = time.Now()
	if _, err := db.messages.InsertOne(ctx, msg); err != nil {
		log.Println("Database->NewMessage: " + err.Error())
		return "", err
	}
	return msg.MessageID, nil

}
