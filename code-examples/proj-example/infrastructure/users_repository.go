package infrastructure

import (
	"server/domain"
)

type UsersRepository interface {
	Add(user *domain.User) error
	GetByLogin(login string) (*domain.User, error)
	GetUsersList() ([]*domain.User, error)
}

type usersRepository struct{}

func NewUsersRepository() UsersRepository {
	return &usersRepository{}
}

func (repo *usersRepository) Add(user *domain.User) error {
	return nil
}

func (repo *usersRepository) GetByLogin(login string) (*domain.User, error) {

	if login == "" {
		return nil, nil
	}

	return &domain.User{
		ID:       "12345",
		Login:    login,
		Password: "somePassword",
		Role:     domain.Admin,
	}, nil
}

func (repo *usersRepository) GetUsersList() ([]*domain.User, error) {
	return []*domain.User{
		&domain.User{
			ID:       "1",
			Login:    "login1",
			Password: "somePassword",
			Role:     domain.DefaultUser,
		},
		&domain.User{
			ID:       "2",
			Login:    "login2",
			Password: "somePassword",
			Role:     domain.Admin,
		},
	}, nil
}
