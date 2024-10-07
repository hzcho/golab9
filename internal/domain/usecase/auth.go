package usecase

import (
	"context"
	"golab9/internal/domain/models"
)

type Auth interface {
	Register(ctx context.Context, request models.RegisterReq) (models.RegisterResponse, error)
	Login(ctx context.Context, request models.LoginReq) (models.LoginResponse, error)
}
