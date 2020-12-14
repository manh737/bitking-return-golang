package api

import (
	"github.com/labstack/echo/v4"
	"github.com/manh737/bitking-return-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestBuyTicket struct {
	Ticket struct {
		TicketId string `json:"ticketid" validate:"required bson:"`
	} `json:"ticket"`
}

func (r *RequestBuyTicket) Bind(c echo.Context, t *model.Ticket) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	oid, err := primitive.ObjectIDFromHex(r.Ticket.TicketId)
	if err != nil {
		return err
	}
	t.ID = oid
	return nil
}

type buyTicketResponse struct {
	Ticket struct {
		TicketId string `json:"ticketid"`
	} `json:"data"`
}

func BuyTicketResponse(t *model.Ticket) *buyTicketResponse {
	r := new(buyTicketResponse)
	r.Ticket.TicketId = t.ID.String()
	return r
}
