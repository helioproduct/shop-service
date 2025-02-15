package main

import (
	"shop-service/internal/repository/user/postgres"
)

func main() {
	// fmt.Println("hello")
	repo := postgres.NewUserRepository(nil)
	repo.CreateUser(nil, nil)

	// zerolog.CallerMarshalFunc = func()

}
