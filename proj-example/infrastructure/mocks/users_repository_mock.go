package mocks

import (
	"server/domain"

	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	mock.Mock
}

func (mock *UsersRepositoryMock) GetByLogin(login string) (*domain.User, error) {
	args := mock.Called(login)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.User), args.Error(1)
}

func (mock *UsersRepositoryMock) GetUsersList() ([]*domain.User, error) {
	args := mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*domain.User), args.Error(1)
}

func (mock *UsersRepositoryMock) Add(user *domain.User) error {
	args := mock.Called(user)

	return args.Error(0)
}
