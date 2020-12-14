package ticket

import (
	"github.com/manh737/bitking-return-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store interface {
	GetByID(primitive.ObjectID) *model.Ticket
}
