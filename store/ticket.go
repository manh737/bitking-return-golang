package store

import (
	"context"
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/manh737/bitking-return-golang/db"
	"github.com/manh737/bitking-return-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketStore struct {
	db *db.DB
}

func NewTicketStore(db *db.DB) *TicketStore {
	return &TicketStore{
		db: db,
	}
}

func (t *TicketStore) GetByID(ticketId primitive.ObjectID) *model.Ticket {
	ticket := &model.Ticket{}
	collection := t.db.DB.Collection("usersTicket")

	err := collection.FindOne(context.TODO(), bson.M{
		"_id": ticketId,
	}).Decode(&ticket)

	if err != nil {

		log.Fatal(err)
		return ticket

	}
	return nil
}
