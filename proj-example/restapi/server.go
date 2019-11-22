package restapi

import (
	"fmt"
	"net/http"
	"server/restapi/handlers"
	"server/restapi/middleware"
)

type Server struct {
	port           string
	usersHandler   handlers.UsersHandler
	adminsHandler  handlers.AdminsHandler
	authMiddleware middleware.AuthMiddleware
}

func NewServer(port string, usersHandler handlers.UsersHandler, adminsHandler handlers.AdminsHandler, authMiddleware middleware.AuthMiddleware) *Server {
	return &Server{
		port:           port,
		usersHandler:   usersHandler,
		adminsHandler:  adminsHandler,
		authMiddleware: authMiddleware,
	}
}

func (server *Server) ConfigureAndRun() {

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/userslist", server.adminsHandler.GetUsersList)

	adminHandler := server.authMiddleware.CheckRole(adminMux)
	adminHandler = server.authMiddleware.CheckAuth(adminHandler)

	userMux := http.NewServeMux()
	userMux.HandleFunc("/user/profile", server.usersHandler.GetUser)

	userHandler := server.authMiddleware.CheckAuth(userMux)

	siteMux := http.NewServeMux()
	siteMux.Handle("/admin/", adminHandler)
	siteMux.Handle("/user/", userHandler)
	siteMux.HandleFunc("/register", server.usersHandler.AddUser)

	fmt.Printf("listening at %s", server.port)
	http.ListenAndServe(server.port, siteMux)
}
