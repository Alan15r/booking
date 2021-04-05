package models

type House struct {
	//tableName  {}struct `pg:"booking_houses"`
	UUID                string   `json:"uuid"`
	OwnerUUID           string   `json:"owner_uuid"`
	Location            Location `json:"lacation"`
	Price               int      `json:"price"`
	Photos              []string `json:"photos"`
	SleepingPlaces      int      `json:"sleeping_places"`
	AllowedNumberPeople int      `json:"allowed_number_people"`
	CheckInTime         string   `json:"check_in_time"`
	MoveOutTime         string   `json:"move_out_time"`
	Assessment          float32  `json:"assessment"`
	Prepaymen           int      `json:"prepaymen"`
	Pool                bool     `json:"pool"`
	Description         string   `json:"description"`
	Blocked             bool     `json:"blocked"`
}

type Location struct {
	Locality string  `json:"locality"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Comment  string  `json:"comment"`
}
