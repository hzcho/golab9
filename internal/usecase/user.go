package usecase

import (
	"context"
	"golab9/internal/domain/models"
	"golab9/internal/domain/service"
)

type User struct {
	userService service.User
}

func NewUser(userService service.User) *User {
	return &User{
		userService: userService,
	}
}

func (u *User) Get(ctx context.Context, filter models.GetUserFilter) ([]models.User, error) {
	users, err := u.userService.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) GetById(ctx context.Context, id uint64) (models.User, error) {
	user, err := u.userService.GetById(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *User) Add(ctx context.Context, user models.AddUser) (uint64, error) {
	id, err := u.userService.Add(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) Update(ctx context.Context, updateUser models.UpdateUser) (models.User, error) {
	user, err := u.userService.Update(ctx, updateUser)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *User) Delete(ctx context.Context, id uint64) (bool, error) {
	deleted, err := u.userService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return deleted, nil
}
