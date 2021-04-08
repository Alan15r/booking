package handler

import "gitlab.com/tuloev_alan/booking.core/config"

type Repository interface {
	MessageRepository
}
type Handler struct {
	DB     Repository
	Config *config.Config
}

type MessageRepository interface {
}
