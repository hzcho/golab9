package view

import "golab9/internal/usecase"

type Views struct {
	*User
	*Auth
}

func NewViews(usecases *usecase.Usecases) *Views {
	return &Views{
		User: NewUser(usecases.User),
		Auth: NewAuth(usecases.Auth),
	}
}
