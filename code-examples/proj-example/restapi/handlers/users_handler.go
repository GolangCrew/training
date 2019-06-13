package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/domain"
	"server/infrastructure"
)

type UsersHandler interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

type usersHandler struct {
	usersRepository infrastructure.UsersRepository
}

func NewUsersHandler(usersRepository infrastructure.UsersRepository) UsersHandler {
	return &usersHandler{
		usersRepository: usersRepository,
	}
}

func (h *usersHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reg domain.RegistrationData
	decoder.Decode(&reg)

	user := &domain.User{
		ID:       "id",
		Login:    reg.Login,
		Password: reg.Password,
		Role:     domain.DefaultUser,
	}

	h.usersRepository.Add(user)

	fmt.Printf("New User: \n Login - %s, Password - %s \n", reg.Login, reg.Password)
	fmt.Fprintf(w, "New User: \n Login - %s, Password - %s \n", reg.Login, reg.Password)
}

func (h *usersHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	user, _ := h.usersRepository.GetByLogin("login")
	fmt.Fprintln(w, user)

}
