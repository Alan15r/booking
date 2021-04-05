package models

type Gazebo struct {
	//tableName  {}struct `pg:"booking_gazebos"`
	UUID        string   `json:"uuid"`
	BaseUUID    string   `json:"base_uuid"`
	BaseName    string   `json:"base_name"`
	Numder      string   `json:"numder"`
	Price       string   `json:"price"`
	NumberSeats int      `json:"number_seats"`
	Closed      bool     `json:"closed"`
	Photos      []string `json:"photos"`
}
