package main

import (
	"github.com/vnniciusg/src/internal/app"
	"github.com/vnniciusg/src/internal/infra"
)

func main() {
	db := infra.NewDatabase()
	userRepository := infra.NewUserRepository(db)
	userUseCase := app.NewUserUseCase(userRepository)
	httpHandler := app.NewHTTPHandler(userUseCase)
	httpHandler.StartServer()
}
