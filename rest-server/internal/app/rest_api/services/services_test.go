package services_test

import (
	"testing"

	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/entities"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/models/dtos"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserRequest_ToUser(t *testing.T) {
	dto := &dtos.CreateUserRequest{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
	}

	user := dto.ToUser()

	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "1234567890", user.PhoneNumber)
}

func TestUpdateUserRequest_ToUser(t *testing.T) {
	dto := &dtos.UpdateUserRequest{
		FirstName:   "Jane",
		LastName:    "Smith",
		Email:       "jane@example.com",
		PhoneNumber: "0987654321",
	}

	user := dto.ToUser()

	assert.Equal(t, "Jane", user.FirstName)
	assert.Equal(t, "Smith", user.LastName)
	assert.Equal(t, "jane@example.com", user.Email)
	assert.Equal(t, "0987654321", user.PhoneNumber)
}

func TestCreateUserResponse_FromUser(t *testing.T) {
	user := &entities.User{
		ID:          1,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
	}

	response := &dtos.CreateUserResponse{}
	result := response.FromUser(user)

	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.Equal(t, "john@example.com", result.Email)
	assert.Equal(t, "User created successfully.", result.Message)
}

func TestUserResponse_MapUserResponse(t *testing.T) {
	user := &entities.User{
		ID:          1,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
	}

	response := &dtos.UserResponse{}
	response.MapUserResponse(user)

	assert.Equal(t, "John", response.FirstName)
	assert.Equal(t, "Doe", response.LastName)
	assert.Equal(t, "john@example.com", response.Email)
	assert.Equal(t, "1234567890", response.PhoneNumber)
}

func TestGetAllUsersResponse_MapUsersResponse(t *testing.T) {
	users := []*entities.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", PhoneNumber: "1234567890"},
		{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "jane@example.com", PhoneNumber: "0987654321"},
	}

	response := &dtos.GetAllUsersResponse{}
	response.MapUsersResponse(users)

	assert.Len(t, response.Users, 2)
	assert.Equal(t, "John", response.Users[0].FirstName)
	assert.Equal(t, "Jane", response.Users[1].FirstName)
}

func TestGetAllUsersResponse_MapUsersResponse_Empty(t *testing.T) {
	users := []*entities.User{}

	response := &dtos.GetAllUsersResponse{}
	response.MapUsersResponse(users)

	assert.Len(t, response.Users, 0)
}

func TestCompletedPart_Struct(t *testing.T) {
	part := struct {
		PartNumber int
		ETag       string
	}{
		PartNumber: 1,
		ETag:       "\"abc123\"",
	}

	assert.Equal(t, 1, part.PartNumber)
	assert.Equal(t, "\"abc123\"", part.ETag)
}
