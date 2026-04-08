package dtos_test

import (
	"testing"

	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/entities"
	"github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/models/dtos"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserRequest_ToUser(t *testing.T) {
	tests := []struct {
		name     string
		request  dtos.CreateUserRequest
		expected entities.User
	}{
		{
			name: "full user data",
			request: dtos.CreateUserRequest{
				FirstName:   "John",
				LastName:    "Doe",
				Email:       "john@example.com",
				PhoneNumber: "1234567890",
			},
			expected: entities.User{
				FirstName:   "John",
				LastName:    "Doe",
				Email:       "john@example.com",
				PhoneNumber: "1234567890",
			},
		},
		{
			name: "with special characters in phone",
			request: dtos.CreateUserRequest{
				FirstName:   "Jane",
				LastName:    "Smith",
				Email:       "jane@example.com",
				PhoneNumber: "+1-555-123-4567",
			},
			expected: entities.User{
				FirstName:   "Jane",
				LastName:    "Smith",
				Email:       "jane@example.com",
				PhoneNumber: "+1-555-123-4567",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := tt.request.ToUser()
			assert.Equal(t, tt.expected.FirstName, user.FirstName)
			assert.Equal(t, tt.expected.LastName, user.LastName)
			assert.Equal(t, tt.expected.Email, user.Email)
			assert.Equal(t, tt.expected.PhoneNumber, user.PhoneNumber)
		})
	}
}

func TestUpdateUserRequest_ToUser(t *testing.T) {
	tests := []struct {
		name     string
		request  dtos.UpdateUserRequest
		expected entities.User
	}{
		{
			name: "full update data",
			request: dtos.UpdateUserRequest{
				FirstName:   "Updated",
				LastName:    "Name",
				Email:       "updated@example.com",
				PhoneNumber: "9998887777",
			},
			expected: entities.User{
				FirstName:   "Updated",
				LastName:    "Name",
				Email:       "updated@example.com",
				PhoneNumber: "9998887777",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := tt.request.ToUser()
			assert.Equal(t, tt.expected.FirstName, user.FirstName)
			assert.Equal(t, tt.expected.LastName, user.LastName)
			assert.Equal(t, tt.expected.Email, user.Email)
			assert.Equal(t, tt.expected.PhoneNumber, user.PhoneNumber)
		})
	}
}

func TestCreateUserResponse_FromUser(t *testing.T) {
	user := &entities.User{
		ID:        42,
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@example.com",
	}

	response := &dtos.CreateUserResponse{}
	result := response.FromUser(user)

	assert.Equal(t, "Test", result.FirstName)
	assert.Equal(t, "User", result.LastName)
	assert.Equal(t, "test@example.com", result.Email)
	assert.Equal(t, "User created successfully.", result.Message)
}

func TestUserResponse_MapUserResponse(t *testing.T) {
	user := &entities.User{
		ID:          1,
		FirstName:   "Map",
		LastName:    "Test",
		Email:       "map@example.com",
		PhoneNumber: "1112223333",
	}

	response := &dtos.UserResponse{}
	response.MapUserResponse(user)

	assert.Equal(t, "Map", response.FirstName)
	assert.Equal(t, "Test", response.LastName)
	assert.Equal(t, "map@example.com", response.Email)
	assert.Equal(t, "1112223333", response.PhoneNumber)
}

func TestGetAllUsersResponse_MapUsersResponse(t *testing.T) {
	t.Run("maps multiple users", func(t *testing.T) {
		users := []*entities.User{
			{ID: 1, FirstName: "User1", LastName: "L1", Email: "u1@example.com", PhoneNumber: "111"},
			{ID: 2, FirstName: "User2", LastName: "L2", Email: "u2@example.com", PhoneNumber: "222"},
			{ID: 3, FirstName: "User3", LastName: "L3", Email: "u3@example.com", PhoneNumber: "333"},
		}

		response := &dtos.GetAllUsersResponse{}
		response.MapUsersResponse(users)

		assert.Len(t, response.Users, 3)
		assert.Equal(t, "User1", response.Users[0].FirstName)
		assert.Equal(t, "User2", response.Users[1].FirstName)
		assert.Equal(t, "User3", response.Users[2].FirstName)
	})

	t.Run("maps empty list", func(t *testing.T) {
		users := []*entities.User{}

		response := &dtos.GetAllUsersResponse{}
		response.MapUsersResponse(users)

		assert.Len(t, response.Users, 0)
	})

	t.Run("maps single user", func(t *testing.T) {
		users := []*entities.User{
			{ID: 1, FirstName: "Single", LastName: "User", Email: "single@example.com", PhoneNumber: "555"},
		}

		response := &dtos.GetAllUsersResponse{}
		response.MapUsersResponse(users)

		assert.Len(t, response.Users, 1)
		assert.Equal(t, "Single", response.Users[0].FirstName)
	})
}

func TestUserResponse_Fields(t *testing.T) {
	resp := dtos.UserResponse{
		FirstName:   "Field",
		LastName:    "Test",
		Email:       "field@example.com",
		PhoneNumber: "444555666",
	}

	assert.Equal(t, "Field", resp.FirstName)
	assert.Equal(t, "Test", resp.LastName)
	assert.Equal(t, "field@example.com", resp.Email)
	assert.Equal(t, "444555666", resp.PhoneNumber)
}

func TestGetAllUsersResponse_Fields(t *testing.T) {
	resp := dtos.GetAllUsersResponse{
		Users: []*dtos.UserResponse{
			{FirstName: "A", LastName: "B", Email: "a@b.com", PhoneNumber: "123"},
		},
	}

	assert.Len(t, resp.Users, 1)
}

func TestCreateUserRequest_Fields(t *testing.T) {
	req := dtos.CreateUserRequest{
		FirstName:   "Create",
		LastName:    "Request",
		Email:       "create@example.com",
		PhoneNumber: "777888999",
	}

	assert.Equal(t, "Create", req.FirstName)
	assert.Equal(t, "Request", req.LastName)
	assert.Equal(t, "create@example.com", req.Email)
	assert.Equal(t, "777888999", req.PhoneNumber)
}

func TestUpdateUserRequest_Fields(t *testing.T) {
	req := dtos.UpdateUserRequest{
		FirstName:   "Update",
		LastName:    "Request",
		Email:       "update@example.com",
		PhoneNumber: "000111222",
	}

	assert.Equal(t, "Update", req.FirstName)
	assert.Equal(t, "Request", req.LastName)
	assert.Equal(t, "update@example.com", req.Email)
	assert.Equal(t, "000111222", req.PhoneNumber)
}

func TestCreateUserResponse_Fields(t *testing.T) {
	resp := dtos.CreateUserResponse{
		FirstName: "CR",
		LastName:  "ESP",
		Email:     "cr@esp.com",
		Message:   "Created",
	}

	assert.Equal(t, "CR", resp.FirstName)
	assert.Equal(t, "ESP", resp.LastName)
	assert.Equal(t, "cr@esp.com", resp.Email)
	assert.Equal(t, "Created", resp.Message)
}
