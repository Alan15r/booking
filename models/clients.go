package models

type Client struct {
	tableName  struct{} `pg:"booking_clients"`
	UUID       string   `json:"uuid"`
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Bookings   []string `json:"bookings"`
	Requisites string   `json:"requisites"`
	Assessment float32  `json:"assessment"`
	Blocked    bool     `json:"blocked"`
}
