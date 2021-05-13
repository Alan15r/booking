package handler

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/tuloev_alan/booking.core/models"
	"gitlab.com/tuloev_alan/booking.core/proto"
)

func (h *Handler) CreateOwner(ctx context.Context, owner *models.Owner) (*models.Owner, error) {
	owners, err := h.DB.GetOwnerByFilter(ctx, &proto.OwnerFilter{Phone: owner.Phone})
	if err != nil {
		return nil, errors.Wrap(err, "fail to get owners")
	}

	if len(owners) > 0 {
		return nil, errors.New("owner with this phone already exists")
	}

	onr, err := h.DB.CreateOwner(ctx, owner)
	if err != nil {
		return nil, errors.Wrap(err, "fail to create owner")
	}
	return onr, nil
}

func (h *Handler) GetOwnerByFilter(ctx context.Context, filter *proto.OwnerFilter) ([]models.Owner, error) {
	owners, err := h.DB.GetOwnerByFilter(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get chats")
	}
	return owners, nil
}

func (h *Handler) GetOwnerByUUID(ctx context.Context, uuid string) (*models.Owner, error) {
	owner, err := h.DB.GetOwnerByUUID(ctx, uuid)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get chats")
	}
	return owner, nil
}
