package main

import (
	"server/core"
	"server/infrastructure"
	"server/restapi"
	"server/restapi/handlers"
	"server/restapi/middleware"
)

func main() {

	usersRepository := infrastructure.NewUsersRepository()

	authGuard := core.NewAuthGuard(usersRepository)

	authMiddleware := middleware.NewAuthMiddleware(authGuard)

	usersHandler := handlers.NewUsersHandler(usersRepository)
	adminsHandler := handlers.NewAdminsHandler(usersRepository)

	server := restapi.NewServer(":8080", usersHandler, adminsHandler, authMiddleware)
	server.ConfigureAndRun()

}
