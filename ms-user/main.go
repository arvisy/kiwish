package main

import (
	"ms-user/cmd"
	"ms-user/config"
	"ms-user/handler"
	"ms-user/repository"
)

func main() {
	db := config.ConnectPostgresDB()

	UserRepository := repository.NewUserRepository(db)
	UserHandler := handler.NewUserHandler(*UserRepository)

	cmd.InitGrpc(*UserHandler)
}
