package handler

import (
	"context"

	"gitlab.com/tuloev_alan/booking.core/config"
	"gitlab.com/tuloev_alan/booking.core/models"
	"gitlab.com/tuloev_alan/booking.core/proto"
)

type Repository interface {
	OwnerRepository
}
type Handler struct {
	DB     Repository
	Config *config.Config
}

type OwnerRepository interface {
	CreateOwner(ctx context.Context, owner *models.Owner) (*models.Owner, error)
	GetOwnerByFilter(ctx context.Context, filter *proto.OwnerFilter) ([]models.Owner, error)
	GetOwnerByUUID(ctx context.Context, uuid string) (*models.Owner, error)
}
