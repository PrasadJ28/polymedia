package repository_test

import (
	"testing"

	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/entities"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewUserRepository(t *testing.T) {
	repo := repository.NewUserRepository(nil)
	assert.NotNil(t, repo)
}

func TestUserEntity(t *testing.T) {
	user := &entities.User{
		ID:          1,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
	}

	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "1234567890", user.PhoneNumber)
}

func TestUserEntity_Empty(t *testing.T) {
	user := &entities.User{}

	assert.Equal(t, 0, user.ID)
	assert.Equal(t, "", user.FirstName)
	assert.Equal(t, "", user.LastName)
	assert.Equal(t, "", user.Email)
	assert.Equal(t, "", user.PhoneNumber)
}
