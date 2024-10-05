package view

import "golab9/internal/usecase"

type Views struct {
	*User
}

func NewViews(usecases *usecase.Usecases) *Views {
	return &Views{
		NewUser(usecases.User),
	}
}
