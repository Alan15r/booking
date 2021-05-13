package proto

import "gitlab.com/tuloev_alan/booking.core/models"

type OwnerFront struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Requisites string `json:"requisites"`
	Blocked    bool   `json:"blocked"`
}

type OwnerFilter struct {
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Blocked bool   `json:"blocked"`
}

type OwnerUpdate struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Requisites string `json:"requisites"`
	Blocked    bool   `json:"blocked"`
}

func OwnerToCore(of *OwnerFront) *models.Owner {
	return &models.Owner{
		UUID:       of.UUID,
		Name:       of.Name,
		Phone:      of.Phone,
		Requisites: of.Requisites,
	}
}

func OwnerToFront(oc *models.Owner) *OwnerFront {
	return &OwnerFront{
		UUID:       oc.UUID,
		Name:       oc.Name,
		Phone:      oc.Phone,
		Requisites: oc.Requisites,
		Blocked:    oc.Blocked,
	}
}

func ArrOwnerToFront(onrs []models.Owner) []OwnerFront {
	owners := make([]OwnerFront, len(onrs))
	for i, v := range onrs {
		owners[i] = *OwnerToFront(&v)
	}
	return owners
}
