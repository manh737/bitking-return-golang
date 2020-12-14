package ticket

import "github.com/manh737/bitking-return-golang/model"

type Store interface {
	GetByID(uint) (*model.Ticket, error)
	GetByEmail(string) (*model.Ticket, error)
	GetByUsername(string) (*model.Ticket, error)
	Create(*model.Ticket) error
	Update(*model.Ticket) error
	AddFollower(ticket *model.Ticket, followerID uint) error
	RemoveFollower(ticket *model.Ticket, followerID uint) error
	IsFollower(ticketID, followerID uint) (bool, error)
}
