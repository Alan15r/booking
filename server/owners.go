package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gitlab.com/tuloev_alan/booking.core/proto"
)

func (r *Rest) CreateOwner(c echo.Context) error {
	ctx := c.Request().Context()

	lg := log.WithFields(log.Fields{
		"event": "owner create",
	})

	onr := new(proto.OwnerFront)
	if err := c.Bind(onr); err != nil {
		lg.WithError(err).Error("owner bind error")
		return c.JSON(http.StatusBadRequest, lg)
	}

	owner := proto.OwnerToCore(onr)
	owner, err := r.Handler.CreateOwner(ctx, owner)
	if err != nil {
		lg.WithError(err).Error("owner create error")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	lg.Info()

	onr = proto.OwnerToFront(owner)
	return c.JSON(http.StatusOK, onr)
}

func (r *Rest) GetOwnerByFilter(c echo.Context) error {
	ctx := c.Request().Context()

	lg := log.WithFields(log.Fields{
		"event": "get owner by filter",
	})

	filter := new(proto.OwnerFilter)
	filter.Name = c.QueryParam("name")
	filter.Phone = c.QueryParam("phone")
	blocked, err := strconv.ParseBool(c.QueryParam("blocked"))
	if err == nil {
		filter.Blocked = blocked
	}

	owners, err := r.Handler.GetOwnerByFilter(ctx, filter)
	if err != nil {
		log.WithError(err).Error("geting owners error")
		return c.JSON(http.StatusBadRequest, err)
	}

	lg.Info()

	onrs := proto.ArrOwnerToFront(owners)
	return c.JSON(http.StatusOK, onrs)
}

func (r *Rest) GetOwnerByUUID(c echo.Context) error {
	ctx := c.Request().Context()

	uuid := c.Param("uuid")

	lg := log.WithFields(log.Fields{
		"event": "get owner by UUID",
		"uuid":  uuid,
	})

	owner, err := r.Handler.GetOwnerByUUID(ctx, uuid)
	if err != nil {
		log.WithError(err).Error("geting owners error")
		return c.JSON(http.StatusBadRequest, err)
	}

	lg.Info()

	onrs := proto.OwnerToFront(owner)
	return c.JSON(http.StatusOK, onrs)
}

// func (r *Rest) UpdateOwner(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	lg := log.WithFields(log.Fields{
// 		"event": "update owner",
// 	})

// 	onrUp := new(proto.OwnerUpdate)
// 	if err := c.Bind(onrUp); err != nil {
// 		lg.WithError(err).Error("updateOwner bind error")
// 		return c.JSON(http.StatusBadRequest, lg)
// 	}

// 	owners, err := r.Handler.UpdateOwner(ctx, onrUp)
// 	if err != nil {
// 		log.WithError(err).Error("geting owners error")
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	lg.Info()

// 	onrs := proto.ArrOwnerToFront(owners)
// 	return c.JSON(http.StatusOK, onrs)
// }
