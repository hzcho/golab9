package service

import (
	"context"
	"golab9/internal/domain/models"
)

type User interface {
	Get(ctx context.Context, filter models.GetUserFilter) ([]models.User, error)
	GetById(ctx context.Context, id uint64) (models.User, error)
	Add(ctx context.Context, user models.AddUser) (uint64, error)
	Update(ctx context.Context, user models.UpdateUser) (models.User, error)
	Delete(ctx context.Context, id uint64) (bool, error)

	SetToken(token string) error
}
