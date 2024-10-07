package usecase

import (
	"golab9/internal/domain/usecase"
	"golab9/internal/service"
)

type Usecases struct {
	usecase.User
	usecase.Auth
}

func NewUsecases(services *service.Services) *Usecases {
	return &Usecases{
		User: NewUser(services.User),
		Auth: NewAuth(services.Auth, services.User),
	}
}
