package handler

import "booking/booking.core/config"

type Repository interface {
	MessageRepository
}

type Publisher interface {
}

type Handler struct {
	DB     Repository
	Pub    Publisher
	Config *config.Config
}

type MessageRepository interface {
}
