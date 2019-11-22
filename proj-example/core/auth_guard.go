package core

import (
	"server/domain"
	"server/infrastructure"
)

//go:generate mockery -name=AuthGuard -case=underscore -dir ./ -output ./mocks/

type AuthGuard interface {
	IsAuthorized(login string, role domain.UserRole) bool
	IsLoggedIn(login string) bool
}

type authGuard struct {
	usersRepository infrastructure.UsersRepository
}

func NewAuthGuard(usersRepository infrastructure.UsersRepository) AuthGuard {
	return &authGuard{
		usersRepository: usersRepository,
	}
}

func (g *authGuard) IsAuthorized(login string, role domain.UserRole) bool {

	user, err := g.usersRepository.GetByLogin(login)
	if err != nil {
		return false
	}
	if user != nil {
		if user.Role == role {
			return true
		}
	}
	return false

}

func (g *authGuard) IsLoggedIn(login string) bool {
	user, err := g.usersRepository.GetByLogin(login)
	if err != nil {
		return false
	}
	if user != nil {
		return true
	}
	return false
}
