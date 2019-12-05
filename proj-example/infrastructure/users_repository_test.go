package infrastructure

import (
	"errors"
	"server/domain"
	"server/infrastructure/client"
	"server/infrastructure/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersRepository_Add(t *testing.T) {
	t.Parallel()
	clientMock := &mocks.Client{}
	expectedErr := errors.New("someErr")
	clientMockFunc := client.NewClientFactoryFuncMock(clientMock, expectedErr)
	repo := NewUsersRepository(clientMockFunc)

	user := &domain.User{}
	actualErr := repo.Add(user)

	assert.Equal(t, expectedErr, actualErr)
}
