package models

import "time"

type Booking struct {
	tableName   struct{}  `pg:"booking_bookings"`
	OwnerUUID   string    `json:"owner_uuid"`
	ClientUUID  string    `json:"client_uuid"`
	HouseUUID   string    `json:"house_uuid"`
	GazeboUUID  string    `json:"gazebo_uuid"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	CancelledAt time.Time `json:"cancelled_at"`
	CheckInTime time.Time `json:"check_in_time"`
	MoveOutTime time.Time `json:"move_out_time"`
}
