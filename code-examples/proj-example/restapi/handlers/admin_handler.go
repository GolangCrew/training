package handlers

import (
	"fmt"
	"net/http"
	"server/infrastructure"
)

type AdminsHandler interface {
	GetUsersList(w http.ResponseWriter, r *http.Request)
}

type adminsHandler struct {
	usersRepository infrastructure.UsersRepository
}

func NewAdminsHandler(usersRepository infrastructure.UsersRepository) AdminsHandler {
	return &adminsHandler{
		usersRepository: usersRepository,
	}
}

func (h *adminsHandler) GetUsersList(w http.ResponseWriter, r *http.Request) {

	users, _ := h.usersRepository.GetUsersList()

	fmt.Fprintln(w, *users[0], *users[1])

}
