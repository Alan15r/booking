package server

import (
	"github.com/labstack/echo/v4"

	handler "booking/booking.core/hendler"
)

const (
	apiPrefix = "/api/v3"
)

type Rest struct {
	Router  *echo.Echo
	Handler *handler.Handler
}

// Route defines all the application rest endpoints
func (r *Rest) Route() {
	//web.UseHealthCheck(r.Router)

	g := r.Router.Group(apiPrefix)

	//g.POST("/messages", r.CreateMessage)
}
