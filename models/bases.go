package models

import "time"

type Base struct {
	//tableName  {}struct `pg:"booking_bases"`
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	OwnerUUID  string    `json:"owner_uuid"`
	Location   Location  `json:"location"`
	Pool       bool      `json:"pool"`
	PlayArea   bool      `json:"play_area"`
	DanceFloor bool      `json:"dance_floor"`
	Photos     []string  `json:"photos"`
	OpenAt     time.Time `json:"open_at"`
	CloseAt    time.Time `json:"close_at"`
	Blocked    bool      `json:"blocked"`
}
