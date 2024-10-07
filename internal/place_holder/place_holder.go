package placeholder

import (
	"context"
	"fmt"
	"golab9/internal/view"
)

type PlaceHolder struct {
	views view.Views
}

func NewPlaceHolder(views view.Views) *PlaceHolder {
	return &PlaceHolder{
		views: views,
	}
}

func (ph *PlaceHolder) Start(ctx context.Context) {
	for {
		var n int
		fmt.Print("введите номер операции:" +
			"\n 1. открыть меню авторизации" +
			"\n 2. открыть меню работы с пользователями" +
			"\n 3. выйти")

		fmt.Print("\n номер: ")
		fmt.Scan(&n)

		switch n {
		case 1:
			ph.views.Auth.Display(ctx)
		case 2:
			ph.views.User.Display(ctx)
		case 3:
			return
		default:
			fmt.Println("нет такой операции")
		}
	}
}
