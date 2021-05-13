package repository

import (
	"context"
	"fmt"

	"gitlab.com/tuloev_alan/booking.core/models"
	"gitlab.com/tuloev_alan/booking.core/proto"
)

func (p *Pg) CreateOwner(ctx context.Context, owner *models.Owner) (*models.Owner, error) {
	timeout, cancelFunc := context.WithTimeout(ctx, queryTimeout)
	defer cancelFunc()

	_, err := p.Db.ModelContext(timeout, owner).Returning("*").Insert()
	fmt.Println(err)

	return owner, err
}

func (p *Pg) GetOwnerByFilter(ctx context.Context, filter *proto.OwnerFilter) ([]models.Owner, error) {
	timeout, cancelFunc := context.WithTimeout(ctx, queryTimeout)
	defer cancelFunc()

	owners := make([]models.Owner, 0)

	query := p.Db.ModelContext(timeout, &owners)

	if filter.Name != "" {
		query.Where("name = ?", filter.Name)
	}

	if filter.Phone != "" {
		query.Where("phone = ?", filter.Phone)
	}

	if filter.Blocked {
		query.Where("blocked = ?", true)
	}

	err := query.Select()

	return owners, err
}

func (p *Pg) GetOwnerByUUID(ctx context.Context, uuid string) (*models.Owner, error) {
	timeout, cancelFunc := context.WithTimeout(ctx, queryTimeout)
	defer cancelFunc()

	owner := new(models.Owner)

	err := p.Db.ModelContext(timeout, owner).
		Where("uuid = ?", uuid).Returning("*").
		Select()

	return owner, err
}
