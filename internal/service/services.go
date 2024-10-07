package service

import (
	"golab9/internal/config"
	"golab9/internal/domain/service"
)

type Services struct {
	service.User
	service.Auth
}

func NewServices(cfg *config.Config) *Services {
	return &Services{
		User: NewUser(cfg),
		Auth: NewAuth(cfg),
	}
}
