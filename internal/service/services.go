package service

import (
	"golab9/internal/config"
	"golab9/internal/domain/service"
)

type Services struct {
	service.User
}

func NewServices(cfg *config.Config) *Services {
	return &Services{
		User: NewUser(cfg),
	}
}
