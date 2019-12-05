package infrastructure

import (
	"server/domain"
	"server/infrastructure/client"
)

type UsersRepository interface {
	Add(user *domain.User) error
	GetByLogin(login string) (*domain.User, error)
	GetUsersList() ([]*domain.User, error)
}

type usersRepository struct {
	clientFactoryFunc client.ClientFactoryFunc
}

func NewUsersRepository(clientFactoryFunc client.ClientFactoryFunc) UsersRepository {
	return &usersRepository{
		clientFactoryFunc: clientFactoryFunc,
	}
}

func (repo *usersRepository) Add(user *domain.User) error {
	// client, err := repo.clientFactoryFunc()
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
