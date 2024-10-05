package usecase

import (
	"golab9/internal/domain/usecase"
	"golab9/internal/service"
)

type Usecases struct {
	usecase.User
}

func NewUsecases(services *service.Services) *Usecases {
	return &Usecases{
		User: NewUser(services.User),
	}
}
