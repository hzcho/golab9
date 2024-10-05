package main

import (
	"context"
	"golab9/internal/config"
	placeholder "golab9/internal/place_holder"
	"golab9/internal/service"
	"golab9/internal/usecase"
	"golab9/internal/view"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	services := service.NewServices(cfg)
	usecases := usecase.NewUsecases(services)
	views := view.NewViews(usecases)
	ph := placeholder.NewPlaceHolder(*views)

	ctx := context.Background()
	ph.Start(ctx)
}
