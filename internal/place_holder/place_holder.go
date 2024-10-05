package placeholder

import (
	"context"
	"fmt"
	"golab9/internal/view"
)

type PlaceHolder struct {
	user view.User
}

func NewPlaceHolder(views view.Views) *PlaceHolder {
	return &PlaceHolder{
		user: *views.User,
	}
}

func (ph *PlaceHolder) Start(ctx context.Context) {
	for {
		var n int
		fmt.Print("введите номер операции:" +
			"\n 1. открыть пользовательский интерфейс" +
			"\n 2. выйти")

		fmt.Print("\n номер: ")
		fmt.Scan(&n)

		switch n {
		case 1:
			ph.user.Display(ctx)
		case 2:
			return
		default:
			fmt.Println("нет такой операции")
		}
	}
}
