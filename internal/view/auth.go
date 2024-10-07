package view

import (
	"context"
	"fmt"
	"golab9/internal/domain/models"
	"golab9/internal/domain/usecase"
)

type Auth struct {
	authUseCase usecase.Auth
}

func NewAuth(authUseCase usecase.Auth) *Auth {
	return &Auth{
		authUseCase: authUseCase,
	}
}

func (a *Auth) Display(ctx context.Context) {
	for {
		var n int
		fmt.Println("\n\n введите номер операции:" +
			"\n 1. зарегистрироваться" +
			"\n 2. авторизироваться" +
			"\n 3. выйти")

		fmt.Print("\n number: ")
		fmt.Scan(&n)

		switch n {
		case 1:
			a.register(ctx)
		case 2:
			a.login(ctx)
		case 3:
			return
		default:
			fmt.Println("\n нет такой операции")
		}
	}
}

func (a *Auth) register(ctx context.Context) {
	var request models.RegisterReq
	fmt.Println("\n введите данные для регистрации")
	fmt.Print("\n логин: ")
	fmt.Scan(&request.Login)
	fmt.Print("\n пароль: ")
	fmt.Scan(&request.Password)

	id, err := a.authUseCase.Register(ctx, request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("id созданного аккаунта: %d", id.AccountId)
}

func (a *Auth) login(ctx context.Context) {
	var request models.LoginReq
	fmt.Println("\n введите данные для авторизации")
	fmt.Print("\n логин: ")
	fmt.Scan(&request.Login)
	fmt.Print("\n пароль: ")
	fmt.Scan(&request.Password)

	resp, err := a.authUseCase.Login(ctx, request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("токен: %s", resp.Token)
}
