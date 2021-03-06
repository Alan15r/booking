package models

import "time"

type Owner struct {
	tableName  struct{}  `pg:"booking_owners"`
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Requisites string    `json:"requisites"`
	Blocked    bool      `json:"blocked"`
	CreatedAt  time.Time `json:"created_at"`
}
