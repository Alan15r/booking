package server

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/tuloev_alan/booking.core/handler"
)

type Rest struct {
	Router  *echo.Echo
	Handler *handler.Handler
}

// Route defines all the application rest endpoints
func (r *Rest) Route() {

	g := r.Router.Group("/v1")

	//owners
	g.POST("/owners", r.CreateOwner)
	g.GET("/owners", r.GetOwnerByFilter)
	g.GET("/owners/:uuid", r.GetOwnerByUUID)
	//g.PUT("/owners", r.UpdateOwner)
}
