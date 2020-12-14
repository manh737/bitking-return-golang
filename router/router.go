package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/manh737/bitking-return-golang/db"
	"github.com/manh737/bitking-return-golang/handler"
	"github.com/manh737/bitking-return-golang/store"
)

// Create creates the gin engine with all routes.
func Create(db *db.DB) *echo.Echo {

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	ticketStore := store.NewTicketStore(db)
	ticketHandler := handler.NewTicketHandler(ticketStore)

	postT := e.Group("/ticket")
	{
		postT.POST("/buy", ticketHandler.BuyTicket)
	}
	return e
}
