package view

import (
	"context"
	"fmt"
	"golab9/internal/domain/models"
	"golab9/internal/domain/usecase"
)

type User struct {
	userUseCase usecase.User
}

func NewUser(userUseCase usecase.User) *User {
	return &User{
		userUseCase: userUseCase,
	}
}

func (u *User) Display(ctx context.Context) {
	for {
		var n int
		fmt.Println("\n введите номер операции:" +
			"\n 1. получить пользователей по фильтру" +
			"\n 2. получить пользователя по id" +
			"\n 3. создать пользователя user" +
			"\n 4. обновить пользователя" +
			"\n 5. удалить пользователя" +
			"\n 6. выйти")

		fmt.Print("\n number: ")
		fmt.Scan(&n)

		switch n {
		case 1:
			u.get(ctx)
		case 2:
			u.getById(ctx)
		case 3:
			u.add(ctx)
		case 4:
			u.update(ctx)
		case 5:
			u.delete(ctx)
		case 6:
			return
		default:
			fmt.Println("\n нет такой операции")
		}
	}
}

func (u *User) get(ctx context.Context) {
	var filter models.GetUserFilter

	fmt.Println("\n введите данные для фильтрации")

	fmt.Print("\n имя пользователя (нажмите Enter для пропуска): ")
	var nameInput string
	fmt.Scanln(&nameInput)
	if nameInput != "" {
		filter.Name = nameInput
	}

	fmt.Print("\n возраст пользователя (нажмите Enter для пропуска): ")
	var ageInput string
	fmt.Scanln(&ageInput)
	if ageInput != "" {
		var age uint8
		fmt.Sscanf(ageInput, "%d", &age)
		filter.Age = int(age)
	}

	fmt.Print("\n страница: ")
	fmt.Scan(&filter.Page)

	fmt.Print("\n количество: ")
	fmt.Scan(&filter.Limit)

	users, err := u.userUseCase.Get(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Printf("\n user"+
			"\n("+
			"\n id:%d"+
			"\n name:%s"+
			"\n age:%d"+
			"\n)",
			user.ID, user.Name, user.Age,
		)
	}
}

func (u *User) getById(ctx context.Context) {
	var id uint64

	fmt.Print("\n id пользователя: ")
	fmt.Scan(&id)

	user, err := u.userUseCase.GetById(ctx, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n user"+
		"\n("+
		"\n id:%d"+
		"\n name:%s"+
		"\n age:%d"+
		"\n)",
		user.ID, user.Name, user.Age,
	)
}

func (u *User) add(ctx context.Context) {
	var addUser models.AddUser

	fmt.Println("\n введите данные для создания пользователя")
	fmt.Print("\n имя пользователя: ")
	fmt.Scan(&addUser.Name)
	fmt.Print("\n возраст пользователя: ")
	fmt.Scan(&addUser.Age)

	id, err := u.userUseCase.Add(ctx, addUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("id созданного пользователя: %d", id)
}

func (u *User) update(ctx context.Context) {
	var updateUser models.UpdateUser

	fmt.Println("\n введите данные для обновления пользователя")
	fmt.Print("\n id пользователя: ")
	fmt.Scan(&updateUser.ID)
	fmt.Print("\n имя пользователя: ")
	fmt.Scan(&updateUser.Name)
	fmt.Print("\n возраст пользователя: ")
	fmt.Scan(&updateUser.Age)

	user, err := u.userUseCase.Update(ctx, updateUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n user"+
		"\n("+
		"\n id:%d"+
		"\n name:%s"+
		"\n age:%d"+
		"\n)",
		user.ID, user.Name, user.Age,
	)
}

func (u *User) delete(ctx context.Context) {
	var id uint64

	fmt.Print("\n id пользователя: ")
	fmt.Scan(&id)

	isDelete, err := u.userUseCase.Delete(ctx, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isDelete {
		fmt.Println("\n пользователь удален")
	} else {
		fmt.Println("\n пользователь не удален")
	}
}
