package usecase

import (
	"context"
	"golab9/internal/domain/models"
	"golab9/internal/domain/service"
)

type Auth struct {
	authService service.Auth
	userService service.User
}

func NewAuth(authService service.Auth, userService service.User) *Auth {
	return &Auth{
		authService: authService,
		userService: userService,
	}
}

func (a *Auth) Register(ctx context.Context, request models.RegisterReq) (models.RegisterResponse, error) {
	resp, err := a.authService.Register(ctx, request)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	return resp, nil
}

func (a *Auth) Login(ctx context.Context, request models.LoginReq) (models.LoginResponse, error) {
	resp, err := a.authService.Login(ctx, request)
	if err != nil {
		return models.LoginResponse{}, err
	}

	a.userService.SetToken(resp.Token)

	return resp, nil
}
