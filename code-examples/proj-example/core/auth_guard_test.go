package core

import (
	"server/domain"
	"server/infrastructure/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Проверяем positive case в котором пользователь является админом
func TestIsAuthorized_ReturnsTrue(t *testing.T) {
	t.Parallel()
	usersRepoMock := &mocks.UsersRepositoryMock{} // создаем mock репозитория

	// user которого мы хотим вернуть из GetByLogin
	user := &domain.User{
		ID:       "id",
		Login:    "login",
		Password: "password",
		Role:     domain.Admin, // устанавливаем роль, чтобы тест прошел
	}

	usersRepoMock.On("GetByLogin", mock.Anything).Return(user, nil) // описываем поведение метода GetByLogin (входные параметры и возврат)

	authGuard := NewAuthGuard(usersRepoMock) // подставляем mock вместо реального репозитория

	expectedResult := true
	actualResult := authGuard.IsAuthorized("login", domain.Admin)

	assert.Equal(t, expectedResult, actualResult) // сравниваем ожидаемый результат с реальным

}
