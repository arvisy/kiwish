package main

import (
	"log"
	"ms-user/cmd"
	"ms-user/config"
	"ms-user/handler"
	"ms-user/repository"
)

func main() {
	db := config.ConnectPostgresDB()

	cache, err := config.InitCache(config.DefaultRedisConfig())
	if err != nil {
		log.Println(err)
	}

	UserRepository := repository.NewUserRepository(db)
	UserHandler := handler.NewUserHandler(*UserRepository, cache)

	cmd.InitGrpc(*UserHandler)
}
